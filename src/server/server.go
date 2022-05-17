package main

import (
	"Server-Monitor/src/proto"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"time"
)

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36"

var (
	clientVerify = &ClientVerify{}
	GrpcPort     string
	HttpPort     string
	fileName     string
	infoData     []*InfoData
)

type services struct {
	proto.UnimplementedMonitorServiceServer
}

type ClientVerify struct {
	HttpPort int `json:"http_port"`
	GrpcPort int `json:"grpc_port"`
	Report   struct {
		ReportTime int64  `json:"report_time"`
		BotApi     string `json:"bot_api"`
		ChatId     string `json:"chat_id"`
	} `json:"report"`
	Servers []struct {
		Id    int    `json:"id"`
		Token string `json:"token"`
		Name  string `json:"name"`
	} `json:"servers"`
}

type Info struct {
	Status int         `json:"status"`
	Data   []*InfoData `json:"data"`
	Ts     int64       `json:"ts"`
}

type InfoData struct {
	ServerStatus string      `json:"server_status"`
	StaticData   StaticData  `json:"static_data"`
	DynamicData  DynamicData `json:"dynamic_data"`
	NetworkData  NetworkData `json:"network_data"`
}

type StaticData struct {
	HostName string `json:"host_name"`
	CpuCore  uint64 `json:"cpuCore"`
	CpuName  string `json:"cpuName"`
	Os       string `json:"os"`
	Platform string `json:"platform"`
	Kernel   string `json:"kernel"`
	Arch     string `json:"arch"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Isp      string `json:"isp"`
	Ts       int64  `json:"ts"`
}

type DynamicData struct {
	Processes uint64   `json:"processes"`
	Uptime    uint64   `json:"uptime"`
	MemTotal  uint64   `json:"memTotal"`
	MemUsed   uint64   `json:"memUsed"`
	SwapTotal uint64   `json:"swapTotal"`
	SwapUsed  uint64   `json:"swap_used"`
	DiskTotal uint64   `json:"diskTotal"`
	DiskUsed  uint64   `json:"diskUsed"`
	CpuUsage  float64  `json:"cpuUsage"`
	MemUsage  float64  `json:"memUsage"`
	SwapUsage float64  `json:"swap_usage"`
	DiskUsage float64  `json:"diskUsage"`
	Load      LoadInfo `json:"load"`
	DiskInfo  []*proto.DiskInfo
	NetSpeed  []*proto.NetSpeed
	Ts        int64 `json:"ts"`
}

type LoadInfo struct {
	Load1  float64 `json:"load_1"`
	Load5  float64 `json:"load_5"`
	Load15 float64 `json:"load_15"`
}

type NetworkData struct {
	CT NetDelay `json:"ct"`
	CU NetDelay `json:"cu"`
	CM NetDelay `json:"cm"`
	Ts int64    `json:"ts"`
}

type NetDelay struct {
	Average uint64  `json:"average"`
	Loss    float64 `json:"loss"`
}

func (s *services) StaticService(_ context.Context, info *proto.StaticInfo) (*proto.ResponseStatus, error) {
	status := verifyClient(info.Verify.Id, info.Verify.Token, info.Ts)

	if status == false {
		return &proto.ResponseStatus{Status: -1}, nil
	}

	staticInfo := StaticData{
		HostName: clientVerify.Servers[info.Verify.Id-1].Name,
		CpuCore:  info.CpuCore,
		CpuName:  info.CpuName,
		Os:       info.Os,
		Platform: info.Platform,
		Kernel:   info.Kernel,
		Arch:     info.Arch,
		Country:  info.Country,
		Province: info.Province,
		City:     info.City,
		Isp:      info.Isp,
		Ts:       info.Ts,
	}

	infoData[info.Verify.Id-1].StaticData = staticInfo
	log.Printf("Client connect id: %v.", info.Verify.Id)
	return &proto.ResponseStatus{Status: 0}, nil
}

func (s *services) DynamicService(_ context.Context, info *proto.DynamicInfo) (*proto.ResponseStatus, error) {
	var loadInfo LoadInfo
	var responseStatus int64 = 0
	status := verifyClient(info.Verify.Id, info.Verify.Token, info.Ts)

	if status == false {
		responseStatus = -1
		return &proto.ResponseStatus{Status: responseStatus}, nil
	}

	if info.Load == nil {
		loadInfo = LoadInfo{
			Load1:  0.00,
			Load5:  0.00,
			Load15: 0.00,
		}
	} else {
		loadInfo = LoadInfo{
			Load1:  info.Load.Load1,
			Load5:  info.Load.Load5,
			Load15: info.Load.Load15,
		}
	}

	dynamicInfo := DynamicData{
		Processes: info.Processes,
		Uptime:    info.Uptime,
		MemTotal:  info.MemTotal,
		MemUsed:   info.MemUsed,
		SwapTotal: info.SwapTotal,
		SwapUsed:  info.SwapUsed,
		DiskTotal: info.DiskTotal,
		DiskUsed:  info.DiskUsed,
		CpuUsage:  info.CpuUsage,
		MemUsage:  info.MemUsage,
		SwapUsage: info.SwapUsage,
		DiskUsage: info.DiskUsage,
		Load:      loadInfo,
		DiskInfo:  info.DiskInfo,
		NetSpeed:  info.NetSpeed,
		Ts:        info.Ts,
	}

	// 上线提醒
	if infoData[info.Verify.Id-1].ServerStatus == "OFFLINE_REPORTED" {
		// 掉线后重新登陆
		log.Printf("Client online again id: %v.", info.Verify.Id)
		onlineReport(info.Verify.Id)
	}

	if infoData[info.Verify.Id-1].StaticData.Ts == 0 {
		responseStatus = -2
	}

	infoData[info.Verify.Id-1].ServerStatus = "ONLINE"
	infoData[info.Verify.Id-1].DynamicData = dynamicInfo

	return &proto.ResponseStatus{Status: responseStatus}, nil
}

func (s *services) NetDelayService(_ context.Context, info *proto.NetDelayInfo) (*proto.ResponseStatus, error) {

	status := verifyClient(info.Verify.Id, info.Verify.Token, info.Ts)
	//fmt.Println(info)
	if status == false {
		return &proto.ResponseStatus{Status: -1}, nil
	}

	networkData := NetworkData{
		CT: NetDelay{Loss: info.CT.Loss, Average: info.CT.Average},
		CU: NetDelay{Loss: info.CU.Loss, Average: info.CU.Average},
		CM: NetDelay{Loss: info.CM.Loss, Average: info.CM.Average},
		Ts: info.Ts,
	}

	infoData[info.Verify.Id-1].NetworkData = networkData
	return &proto.ResponseStatus{Status: 0}, nil
}

func (s *services) ReportService(_ context.Context, info *proto.ClientReport) (*proto.ResponseStatus, error) {
	status := verifyClient(info.Verify.Id, info.Verify.Token, info.Ts)

	if status == false {
		return &proto.ResponseStatus{Status: -1}, nil
	}

	if clientVerify.Report.BotApi == "" || clientVerify.Report.ChatId == "" {
		return &proto.ResponseStatus{Status: -3}, nil
	}

	telegramReport(info.Msg)
	return &proto.ResponseStatus{Status: 0}, nil
}

// 验证token
func verifyClient(id uint64, verifyToken string, now int64) bool {
	token := clientVerify.Servers[id-1].Token
	localSign := verifyClientToken(token, now)
	if localSign != verifyToken {
		//log.Println(token, now, localSign)
		return false
	}
	return true
}

// 计算MD5
func verifyClientToken(s string, n int64) string {
	str := fmt.Sprintf("%v%v%v", s, n, "JPCXcgFVnkj7z3vZ")
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// 验证掉线
func checkClient() {
	for {
		task := time.NewTimer(1 * time.Second)
		//log.Println("正在检查主机状态")
		now := time.Now().Unix()
		for n, i := range infoData {
			switch i.ServerStatus {
			case "":
				i.ServerStatus = "OFFLINE"
			case "ONLINE":
				if now > i.DynamicData.Ts+10 {
					i.ServerStatus = "OFFLINE"
					log.Printf("Client offline id: %v.", n+1)
				}
			case "OFFLINE":
				if i.DynamicData.Ts == 0 {
					continue
				} else if now > i.DynamicData.Ts+clientVerify.Report.ReportTime {
					i.ServerStatus = "OFFLINE_REPORTED"
					offlineReport(n, i.DynamicData.Ts)
					infoData[n].StaticData = StaticData{
						HostName: clientVerify.Servers[n].Name,
					}
					infoData[n].DynamicData = DynamicData{}
					infoData[n].NetworkData = NetworkData{}

				}
			case "OFFLINE_REPORTED":
				continue
			}
		}
		<-task.C
	}
}

func offlineReport(n int, i int64) {
	t := time.Unix(i, 0).Format("2006/01/02 15:04:05")
	notify := fmt.Sprintf("【⚠️OFFLINE REPORT⚠️】\n\nID: %v\nName: %v\n"+
		"CPU usage: %v %%\nMemory usage: %v %%\nSwap usage: %v %%\n"+
		"DIsk usage: %v %%\nLast online time: %v",
		n+1, clientVerify.Servers[n].Name, infoData[n].DynamicData.CpuUsage,
		infoData[n].DynamicData.MemUsage, infoData[n].DynamicData.SwapUsage,
		infoData[n].DynamicData.DiskUsage, t)
	//log.Printf(notify)

	if clientVerify.Report.BotApi != "" && clientVerify.Report.ChatId != "" {
		go telegramReport(notify)
		log.Printf("Clinet offline reported id: %v.", n+1)
	}
}

func onlineReport(n uint64) {
	t := time.Unix(time.Now().Unix(), 0).Format("2006/01/02 15:04:05")
	notify := fmt.Sprintf("【✅ONLINE REPORT✅】\n\nID: %v\nName: %v\nOnline time: %v",
		n, clientVerify.Servers[n-1].Name, t)
	//log.Printf(notify)

	if clientVerify.Report.BotApi != "" && clientVerify.Report.ChatId != "" {
		go telegramReport(notify)
		log.Printf("Clinet online reported id: %v.", n)
	}
}

func telegramReport(s string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", clientVerify.Report.BotApi)

	data := map[string]string{
		"chat_id": clientVerify.Report.ChatId,
		"text":    s,
	}

	client := resty.New()
	_, err := client.R().
		SetHeader("user-agent", userAgent).
		SetFormData(data).
		Post(url)
	checkError(err)
}

// gin中间件
func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func httpServer() {
	r := gin.Default()
	// 使用 kratos 中间件
	r.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	r.LoadHTMLGlob("web/index.html")
	r.Static("/web", "./web")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/api/info", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		i := &Info{
			Status: 0,
			Data:   infoData,
			Ts:     time.Now().Unix(),
		}
		c.JSON(200, i)
	})

	httpSrv := http.NewServer(http.Address(HttpPort))
	httpSrv.HandlePrefix("/", r)

	app := kratos.New(
		kratos.Name("gin"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	flag.StringVar(&fileName, "c", "./server.json", "path to config")
	flag.Parse()
}

func main() {
	jsonFile, err := ioutil.ReadFile(fileName)
	checkError(err)
	err = json.Unmarshal(jsonFile, clientVerify)
	checkError(err)
	HttpPort = fmt.Sprintf(":%v", clientVerify.HttpPort)
	GrpcPort = fmt.Sprintf(":%v", clientVerify.GrpcPort)

	infoData = make([]*InfoData, len(clientVerify.Servers))
	//staticData = make([]*StaticData, len(clientVerify.Servers))
	//dynamicData = make([]*pb.DynamicInfo, len(clientVerify.Servers))
	//netDelayData = make([]*pb.NetDelayInfo, len(clientVerify.Servers))

	// 初始化结构体
	for i, v := range clientVerify.Servers {
		infoData[i] = &InfoData{
			ServerStatus: "OFFLINE",
			StaticData: StaticData{
				HostName: v.Name,
			},
		}
		//fmt.Println(infoData[i].StaticData.HostName)
	}

	// 启动服务
	go httpServer()
	go checkClient()
	s := &services{}
	grpcSrv := grpc.NewServer(
		grpc.Address(GrpcPort),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)
	proto.RegisterMonitorServiceServer(grpcSrv, s)

	app := kratos.New(
		kratos.Name("Server monitor"),
		kratos.Server(
			grpcSrv,
		),
	)
	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}

func checkError(err error) {
	if err != err {
		log.Fatal(err)
	}
}
