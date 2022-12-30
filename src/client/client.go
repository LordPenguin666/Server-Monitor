package main

import (
	"Server-Monitor/src/proto"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-ping/ping"
	"github.com/go-resty/resty/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"google.golang.org/grpc"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"sync"
	"time"
)

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (HTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36"

const (
	// SecondsPerMinute 定义每分钟的秒数
	SecondsPerMinute = 60

	// SecondsPerHour 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60

	// SecondsPerDay 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

var (
	replyTime int
	cpuExec, memExec, swapExec, diskExec,
	cpuReported, memReported, swapReported, diskReported bool
	cpuExecTime, memExecTime, swapExecTime, diskExecTime,
	cpuNormalTime, memNormalTime, swapNormalTime, diskNormalTime int64
	fileName, serverAddr, token                  string
	address                                      []string
	recentDevices, recentDisks                   []string
	recentRead, recentWrite, recentIn, recentOut []uint64
	startTask                                    sync.WaitGroup
	verifyInfo                                   = &VerifyInfo{}
	locationData                                 = &Location{}
	conn                                         *grpc.ClientConn
)

// Location 位置信息
type Location struct {
	Data struct {
		Country  string `json:"country"`
		Province string `json:"province"`
		City     string `json:"city"`
		Isp      string `json:"isp"`
	}
}

// VerifyInfo 登录信息
type VerifyInfo struct {
	Id      int    `json:"id"`
	Token   string `json:"token"`
	Server  string `json:"server"`
	Address struct {
		CT string `json:"CT"`
		CU string `json:"CU"`
		CM string `json:"CM"`
	} `json:"Address"`
	Report struct {
		CPU            int   `json:"CPU"`
		Mem            int   `json:"Mem"`
		Swap           int   `json:"Swap"`
		Disk           int   `json:"Disk"`
		CpuReportTime  int64 `json:"cpuReportTime"`
		MemReportTime  int64 `json:"memReportTime"`
		SwapReportTime int64 `json:"swapReportTime"`
		DiskReportTime int64 `json:"diskReportTime"`
	} `json:"report"`
}

// 测试ICMP
func netDelay() *proto.NetDelayInfo {
	c := [3]chan *proto.NetDelay{}
	for i := 0; i < 3; i++ {
		c[i] = make(chan *proto.NetDelay)
		go pong(address[i], c[i])
	}
	result1 := <-c[0]
	result2 := <-c[1]
	result3 := <-c[2]

	return &proto.NetDelayInfo{
		CT: &proto.NetDelay{Average: result1.Average, Loss: result1.Loss},
		CU: &proto.NetDelay{Average: result2.Average, Loss: result2.Loss},
		CM: &proto.NetDelay{Average: result3.Average, Loss: result3.Loss},
	}
}

// 执行ping
func pong(addr string, c chan *proto.NetDelay) {
	h, err := host.Info()
	pinger, err := ping.NewPinger(addr)
	if h.Platform != "arch" {
		pinger.SetPrivileged(true)
	}
	checkErr(err)
	pinger.Count = 100
	err = pinger.Run() // Blocks until finished.
	checkErr(err)
	s := &proto.NetDelay{
		Average: uint64(pinger.Statistics().AvgRtt.Milliseconds()),
		Loss:    pinger.Statistics().PacketLoss,
	}
	// get send/receive/duplicate/rtt stats
	//fmt.Println(addr)
	c <- s
}

// 获取各个硬盘的IO
func disks() []*proto.DiskInfo {
	var d []string           //硬盘列表（乱序）
	var dd []*proto.DiskInfo //需要上传的信息

	i, err := disk.IOCounters()
	checkErr(err)

	for diskName := range i {
		// 正则，分离出硬盘，去掉分区
		diskN, err := regexp.MatchString("[a-zA-Z]:|hd[a-zA-Z]\\b|sd[a-zA-Z]\\b|vd[a-zA-Z]\\b|nvme\\dn\\d\\b|mmcblk\\d\\b", diskName)
		checkErr(err)
		if diskN == false {
			continue
		}
		// 将硬盘名称添加到一个切片中
		d = append(d, diskName)

	}

	// 对硬盘名称进行排序
	sort.Strings(d)

	// 通过硬盘名称获取总读写

	switch {
	// 第一次运行
	case len(recentRead) == 0:
		// fmt.Println("First time to run.")
		// 初始化切片
		for _, a := range d {
			recentDisks = append(recentDisks, a)
			recentRead = append(recentRead, i[a].ReadBytes)
			recentWrite = append(recentWrite, i[a].WriteBytes)
		}
		dd = nil
	// 出现设备增加或者减少
	case len(d) != len(recentDisks):
		recentDisks = nil
		recentRead = nil
		recentWrite = nil
		dd = nil
	// 正常运行
	default:
		for n, a := range d {
			infos := &proto.DiskInfo{
				Name:  a,
				Read:  (i[a].ReadBytes - recentRead[n]) / 3,
				Write: (i[a].WriteBytes - recentWrite[n]) / 3,
			}
			dd = append(dd, infos)
			recentRead[n] = i[a].ReadBytes
			recentWrite[n] = i[a].WriteBytes
		}
	}
	return dd
}

// 获取网卡信息
func netInfo() []*proto.NetSpeed {
	var netDevices []*proto.NetSpeed

	n, err := net.IOCounters(true)
	checkErr(err)
	switch {
	case len(recentDevices) == 0:
		for _, v := range n {
			recentDevices = append(recentDevices, v.Name)
			recentIn = append(recentIn, v.BytesRecv)
			recentOut = append(recentOut, v.BytesSent)
		}
		netDevices = nil
	case len(n) != len(recentDevices):
		recentDevices = nil
		recentIn = nil
		recentOut = nil
		netDevices = nil
	default:
		for i, v := range n {
			netNIO := &proto.NetSpeed{
				Name: v.Name,
				In:   (v.BytesRecv - recentIn[i]) / 3,
				Out:  (v.BytesSent - recentOut[i]) / 3,
			}
			recentIn[i] = v.BytesRecv
			recentOut[i] = v.BytesSent
			netDevices = append(netDevices, netNIO)
		}
	}
	return netDevices
}

// 获取平均负载
func getLoad() *proto.Load {
	loadAvg, err := load.Avg()
	checkErr(err)
	l := &proto.Load{
		Load1:  loadAvg.Load1,
		Load5:  loadAvg.Load5,
		Load15: loadAvg.Load15,
	}
	return l
}

// 获取IP位置
func getLocation() {
	client := resty.New()
	_, err := client.R().
		EnableTrace().
		SetHeader("user-agent", userAgent).
		SetResult(locationData).
		Get("https://api.bilibili.com/x/web-interface/zone")
	checkErr(err)
}

// 获取静态数据
func acqStatic() {
	now := time.Now().Unix()

	// 获取CPU名称、核心数
	c, err := cpu.Info()
	checkErr(err)
	cpuName := c[0].ModelName
	cpuCores, err := cpu.Counts(true)
	checkErr(err)
	//fmt.Printf("%s (%v)\n", cpuName, cpuCores)

	// 获取主机信息
	h, err := host.Info()
	checkErr(err)
	osSys := h.OS
	platform := h.Platform
	kernel := h.KernelVersion
	arch := h.KernelArch
	//fmt.Printf("\nOS: %v %v\nKernel: %v\nArch: %v\n\n", os, platform, kernel, arch)

	//fmt.Printf("%v %v%v\nISP: %v\n", data.Country, data.Province, data.City, data.Isp)

	// 生成md5
	//log.Println(1, token, now)
	verify := &proto.Verify{
		Id:    uint64(verifyInfo.Id),
		Token: verifyToken(token, now),
	}

	// 发送至服务端
	grpcClient := proto.NewMonitorServiceClient(conn)
	reply, err := grpcClient.StaticService(context.Background(), &proto.StaticInfo{
		CpuCore:  uint64(cpuCores),
		CpuName:  cpuName,
		Os:       osSys,
		Platform: platform,
		Kernel:   kernel,
		Arch:     arch,
		Country:  locationData.Data.Country,
		Province: locationData.Data.Province,
		City:     locationData.Data.City,
		Isp:      locationData.Data.Isp,
		Verify:   verify,
		Ts:       now,
	})

	if err != nil {
		log.Println(err)
	}

	if reply != nil {
		switch reply.Status {
		case 0:
			log.Println("Connected to server.")
		case -1:
			log.Fatal("Token access denied.")
		}
	}
}

// 获取动态数据
func acqDynamic() {
	var swapUsage float64

	for {
		task := time.NewTimer(3 * time.Second)
		now := time.Now().Unix()

		// 获取CPU占用
		c, err := cpu.Percent(time.Second, false)
		checkErr(err)
		cpuUsage, err := strconv.ParseFloat(fmt.Sprintf("%.0f", c[0]), 64)
		checkErr(err)
		//fmt.Printf("\nCPU\nUsage: "+"v%%", cpuUsage)

		// 主机信息
		h, err := host.Info()
		checkErr(err)
		proc := h.Procs
		uptime := h.Uptime

		//fmt.Printf("\nProcess: %v\nUptime: %v\n\n", proc, uptime)

		m, err := mem.VirtualMemory()
		checkErr(err)
		// 内存信息
		memTotal := m.Total
		memUsed := m.Used
		memUsage, err := strconv.ParseFloat(fmt.Sprintf("%.0f", float64(memUsed)/float64(memTotal)*100), 64)
		checkErr(err)
		//fmt.Printf("Memory\nTotal: %v\nUsed: %v\nPresent: %.2f %% \n\n", memTotal, memUsed, memPercent)

		// Swap
		s, err := mem.SwapMemory()
		checkErr(err)

		swapTotal := s.Total
		swapUsed := s.Used
		if swapTotal == 0 {
			swapUsage = 0
		} else {
			swapUsage, err = strconv.ParseFloat(fmt.Sprintf("%.0f", float64(swapUsed)/float64(swapTotal)*100), 64)
			checkErr(err)
		}

		//fmt.Printf("Swap\nTotal: %v\nUsed: %v\nPresent: %.2f %% \n\n", swapTotal, swapUsed, swapPercent)

		// 硬盘使用率(根分区)
		u, err := disk.Usage("/")
		checkErr(err)
		diskUsage, err := strconv.ParseFloat(fmt.Sprintf("%.0f", u.UsedPercent), 64)
		checkErr(err)

		//警报提示
		go check(cpuUsage, memUsage, swapUsage, diskUsage, now)
		//fmt.Printf("Disk\nTotal: %v\nUsed: %v\nUsedPercent: %.2f %%\n\n", u.Total, u.Used, u.UsedPercent)

		// 发送至服务端
		grpcClient := proto.NewMonitorServiceClient(conn)

		// 生成md5
		//log.Println(2, token, now)
		verify := &proto.Verify{
			Id:    uint64(verifyInfo.Id),
			Token: verifyToken(token, now),
		}

		reply, err := grpcClient.DynamicService(context.Background(), &proto.DynamicInfo{
			Processes: proc,
			Uptime:    uptime,
			MemTotal:  memTotal,
			MemUsed:   memUsed,
			SwapTotal: swapTotal,
			SwapUsed:  swapUsed,
			DiskTotal: u.Total,
			DiskUsed:  u.Used,
			CpuUsage:  cpuUsage,
			MemUsage:  memUsage,
			SwapUsage: swapUsage,
			DiskUsage: diskUsage,
			DiskInfo:  disks(),
			NetSpeed:  netInfo(),
			Load:      getLoad(),
			Verify:    verify,
			Ts:        now,
		})

		if err != nil {
			log.Println(err)
		}

		if reply != nil {
			switch reply.Status {
			case -1:
				log.Fatal("Token access denied.")
			case -2:
				log.Println("Try to reconnect to the server.")
				go acqStatic()
			}
		} else {
			task.Reset(30 * time.Second)
		}

		<-task.C
	}
}

// 获取网络延迟
func acqNetDelay() {
	for {
		now := time.Now().Unix()
		netDelayInfo := netDelay()

		grpcClient := proto.NewMonitorServiceClient(conn)

		// 生成md5
		//log.Println(3, token, now)
		verify := &proto.Verify{
			Id:    uint64(verifyInfo.Id),
			Token: verifyToken(token, now),
		}

		reply, err := grpcClient.NetDelayService(context.Background(), &proto.NetDelayInfo{
			CT:     netDelayInfo.CT,
			CU:     netDelayInfo.CU,
			CM:     netDelayInfo.CM,
			Verify: verify,
			Ts:     now,
		})

		if err != nil {
			log.Println(err)
		}

		if reply != nil {
			if reply.Status == -1 {
				log.Println(reply.Status)
				log.Fatal("Token access denied.")
			}
		}
	}
}

func reportToServer(msg string) {
	grpcClient := proto.NewMonitorServiceClient(conn)

	now := time.Now().Unix()
	verify := &proto.Verify{
		Id:    uint64(verifyInfo.Id),
		Token: verifyToken(token, now),
	}

	reply, err := grpcClient.ReportService(context.Background(), &proto.ClientReport{
		Msg:    msg,
		Verify: verify,
		Ts:     now,
	})

	if err != nil {
		log.Println(err)
	}

	if reply != nil {
		switch reply.Status {
		case 0:
			replyTime = 0
		case -1:
			log.Println(reply.Status)
			log.Fatal("Token access denied.")
		case -3:
			if replyTime == 0 {
				log.Println("Please enter your bot config.")
				replyTime++
			}
		}
	}
}

func check(cpu, mem, swap, disk float64, ts int64) {
	//所有状态默认均为 false
	if verifyInfo.Report.CPU != 0 {
		if int(cpu) >= verifyInfo.Report.CPU {
			// 状态从false到true
			if cpuExec == false {
				cpuExec = true
				cpuExecTime = ts
			}
			if ts >= cpuExecTime+verifyInfo.Report.CpuReportTime && cpuReported == false {
				msg := errReportMessage("CPU", verifyInfo.Report.CPU, cpu,
					verifyInfo.Report.CpuReportTime, cpuExecTime)
				go reportToServer(msg)
				cpuReported = true
			}
		} else {
			if cpuExec == true {
				cpuExec = false
				cpuNormalTime = ts
			}
			if ts >= cpuNormalTime+30 && cpuReported == true {
				msg := normalReportMessage("CPU", verifyInfo.Report.CPU, cpu,
					ts-cpuExecTime, cpuNormalTime)
				go reportToServer(msg)
				cpuReported = false
				cpuNormalTime, cpuExecTime = 0, 0
			}
		}
	}

	if verifyInfo.Report.Mem != 0 {
		if int(mem) >= verifyInfo.Report.Mem {
			if memExec == false {
				memExec = true
				memExecTime = ts
			}
			if ts >= memExecTime+verifyInfo.Report.MemReportTime && memReported == false {
				msg := errReportMessage("MEM", verifyInfo.Report.Mem, mem,
					verifyInfo.Report.MemReportTime, memExecTime)
				go reportToServer(msg)
				memReported = true
			}
		} else {
			if memExec == true {
				memExec = false
				memNormalTime = ts
			}
			if ts >= memNormalTime+30 && memReported == true {
				msg := normalReportMessage("MEM", verifyInfo.Report.Mem, mem,
					ts-memExecTime, memNormalTime)
				go reportToServer(msg)
				memReported = false
				memNormalTime, memExecTime = 0, 0
			}
		}
	}

	if verifyInfo.Report.Swap != 0 {
		if int(swap) >= verifyInfo.Report.Swap {
			if swapExec == false {
				swapExec = true
				swapExecTime = ts
			}
			if ts >= swapExecTime+verifyInfo.Report.SwapReportTime && swapReported == false {
				msg := errReportMessage("Swap", verifyInfo.Report.Swap, swap,
					verifyInfo.Report.SwapReportTime, swapExecTime)
				go reportToServer(msg)
				swapReported = true
			}
		} else {
			if swapExec == true {
				swapExec = false
				swapNormalTime = ts
			}
			if ts >= swapNormalTime+30 && swapReported == true {
				msg := normalReportMessage("Swap", verifyInfo.Report.Swap, swap,
					ts-swapExecTime, swapNormalTime)
				go reportToServer(msg)
				swapReported = false
				swapNormalTime, swapExecTime = 0, 0
			}
		}
	}

	if verifyInfo.Report.Disk != 0 {
		if int(disk) >= verifyInfo.Report.Disk {
			if diskExec == false {
				diskExec = true
				diskExecTime = ts
			}
			if ts >= diskExecTime+verifyInfo.Report.DiskReportTime && diskReported == false {
				msg := errReportMessage("Disk", verifyInfo.Report.Disk, disk,
					verifyInfo.Report.DiskReportTime, diskExecTime)
				go reportToServer(msg)
				diskReported = true
			}
		} else {
			if diskExec == true {
				diskExec = false
				diskNormalTime = ts
			}
			if ts >= diskNormalTime+30 && diskReported == true {
				msg := normalReportMessage("Disk", verifyInfo.Report.Disk, disk,
					ts-diskExecTime, diskNormalTime)
				go reportToServer(msg)
				diskReported = false
				diskNormalTime, diskExecTime = 0, 0
			}
		}
	}
}

func errReportMessage(model string, execPercent int, nowPercent float64, reportTime, execTime int64) string {
	t := time.Unix(execTime, 0).Format("2006/01/02 15:04:05")
	str := fmt.Sprintf("【%v 异常警告】\n\n"+
		"警告阈值: %v%%\n"+
		"当前使用率: %v%%\n"+
		"异常持续时间: %v\n"+
		"异常出现时间:%v\n",
		model, execPercent, nowPercent, formatSecond(reportTime), t)
	return str
}

func normalReportMessage(model string, execPercent int, nowPercent float64, reportTime, normalTime int64) string {
	t := time.Unix(normalTime, 0).Format("2006/01/02 15:04:05")
	str := fmt.Sprintf("【%v 恢复提示】\n\n"+
		"警告阈值: %v%%\n"+
		"当前使用率: %v%%\n"+
		"异常持续时间: %v\n"+
		"异常恢复时间:%v\n",
		model, execPercent, nowPercent, formatSecond(reportTime), t)
	return str
}

func formatSecond(seconds int64) string {
	var d, h, m, s int64

	switch {
	case seconds > SecondsPerDay:
		d = seconds / SecondsPerDay
		h = seconds % SecondsPerDay / SecondsPerHour
		m = seconds % SecondsPerDay % SecondsPerHour / SecondsPerMinute
		s = seconds % 60
		return fmt.Sprintf("%v天%v小时%v分%v秒", d, h, m, s)
	case seconds > SecondsPerHour:
		h = seconds / SecondsPerHour
		m = seconds % SecondsPerHour / SecondsPerMinute
		s = seconds % 60
		return fmt.Sprintf("%v小时%v分%v秒", h, m, s)
	case seconds > SecondsPerMinute:
		m = seconds / SecondsPerMinute
		s = seconds % 60
		return fmt.Sprintf("%v分%v秒", m, s)
	default:
		s = seconds
		return fmt.Sprintf("%v秒", s)
	}
}

// 获取登录信息
func startClient() {
	acqStatic()
	startTask.Add(2)
	go acqDynamic()
	go acqNetDelay()
	startTask.Wait()
}

// 计算MD5
func verifyToken(s string, n int64) string {
	str := fmt.Sprintf("%v%v%v", s, n, "JPCXcgFVnkj7z3vZ")
	m := md5.New()
	m.Write([]byte(str))
	sign := hex.EncodeToString(m.Sum(nil))
	//log.Println("sign:", s, n, sign)
	return sign
}

func init() {
	flag.StringVar(&fileName, "c", "./client.json", "log in user")
	flag.Parse()
}

func main() {
	getLocation()
	jsonFile, err := os.ReadFile(fileName)
	checkErr(err)
	err = json.Unmarshal(jsonFile, verifyInfo)
	checkErr(err)

	serverAddr = verifyInfo.Server
	token = verifyInfo.Token

	address = append(address, verifyInfo.Address.CT)
	address = append(address, verifyInfo.Address.CU)
	address = append(address, verifyInfo.Address.CM)

	//tlsConfig := &tls.Config{
	//	InsecureSkipVerify: true,
	//}
	//
	//c := grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))
	//
	//conn, err = grpc.Dial(serverAddr, c)

	conn, err = transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint(serverAddr),
		transgrpc.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	checkErr(err)

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	// 运行客户端，获取系统信息
	startClient()
}

func checkErr(err error) {
	if err != err {
		log.Fatal(err)
	}
}
