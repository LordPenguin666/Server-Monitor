// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: src/proto/SysInfo.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StaticInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuCore  uint64  `protobuf:"varint,1,opt,name=cpuCore,proto3" json:"cpuCore,omitempty"`
	CpuName  string  `protobuf:"bytes,2,opt,name=cpuName,proto3" json:"cpuName,omitempty"`
	Os       string  `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	Platform string  `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Kernel   string  `protobuf:"bytes,5,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Arch     string  `protobuf:"bytes,6,opt,name=arch,proto3" json:"arch,omitempty"`
	Country  string  `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	Province string  `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	City     string  `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	Isp      string  `protobuf:"bytes,10,opt,name=isp,proto3" json:"isp,omitempty"`
	Verify   *Verify `protobuf:"bytes,1000,opt,name=verify,proto3" json:"verify,omitempty"`
	Ts       int64   `protobuf:"varint,10001,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *StaticInfo) Reset() {
	*x = StaticInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticInfo) ProtoMessage() {}

func (x *StaticInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticInfo.ProtoReflect.Descriptor instead.
func (*StaticInfo) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{0}
}

func (x *StaticInfo) GetCpuCore() uint64 {
	if x != nil {
		return x.CpuCore
	}
	return 0
}

func (x *StaticInfo) GetCpuName() string {
	if x != nil {
		return x.CpuName
	}
	return ""
}

func (x *StaticInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *StaticInfo) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *StaticInfo) GetKernel() string {
	if x != nil {
		return x.Kernel
	}
	return ""
}

func (x *StaticInfo) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *StaticInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *StaticInfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *StaticInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *StaticInfo) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

func (x *StaticInfo) GetVerify() *Verify {
	if x != nil {
		return x.Verify
	}
	return nil
}

func (x *StaticInfo) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type DynamicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes uint64      `protobuf:"varint,1,opt,name=processes,proto3" json:"processes,omitempty"`
	Uptime    uint64      `protobuf:"varint,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	MemTotal  uint64      `protobuf:"varint,3,opt,name=memTotal,proto3" json:"memTotal,omitempty"`
	MemUsed   uint64      `protobuf:"varint,4,opt,name=memUsed,proto3" json:"memUsed,omitempty"`
	SwapTotal uint64      `protobuf:"varint,5,opt,name=swapTotal,proto3" json:"swapTotal,omitempty"`
	SwapUsed  uint64      `protobuf:"varint,6,opt,name=swapUsed,proto3" json:"swapUsed,omitempty"`
	DiskTotal uint64      `protobuf:"varint,7,opt,name=diskTotal,proto3" json:"diskTotal,omitempty"`
	DiskUsed  uint64      `protobuf:"varint,8,opt,name=diskUsed,proto3" json:"diskUsed,omitempty"`
	CpuUsage  float64     `protobuf:"fixed64,9,opt,name=cpuUsage,proto3" json:"cpuUsage,omitempty"`
	MemUsage  float64     `protobuf:"fixed64,10,opt,name=memUsage,proto3" json:"memUsage,omitempty"`
	SwapUsage float64     `protobuf:"fixed64,11,opt,name=swapUsage,proto3" json:"swapUsage,omitempty"`
	DiskUsage float64     `protobuf:"fixed64,12,opt,name=diskUsage,proto3" json:"diskUsage,omitempty"`
	Load      *Load       `protobuf:"bytes,100,opt,name=load,proto3" json:"load,omitempty"`
	DiskInfo  []*DiskInfo `protobuf:"bytes,101,rep,name=diskInfo,proto3" json:"diskInfo,omitempty"`
	NetSpeed  []*NetSpeed `protobuf:"bytes,102,rep,name=netSpeed,proto3" json:"netSpeed,omitempty"`
	Verify    *Verify     `protobuf:"bytes,1000,opt,name=verify,proto3" json:"verify,omitempty"`
	Ts        int64       `protobuf:"varint,10001,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *DynamicInfo) Reset() {
	*x = DynamicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicInfo) ProtoMessage() {}

func (x *DynamicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicInfo.ProtoReflect.Descriptor instead.
func (*DynamicInfo) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{1}
}

func (x *DynamicInfo) GetProcesses() uint64 {
	if x != nil {
		return x.Processes
	}
	return 0
}

func (x *DynamicInfo) GetUptime() uint64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *DynamicInfo) GetMemTotal() uint64 {
	if x != nil {
		return x.MemTotal
	}
	return 0
}

func (x *DynamicInfo) GetMemUsed() uint64 {
	if x != nil {
		return x.MemUsed
	}
	return 0
}

func (x *DynamicInfo) GetSwapTotal() uint64 {
	if x != nil {
		return x.SwapTotal
	}
	return 0
}

func (x *DynamicInfo) GetSwapUsed() uint64 {
	if x != nil {
		return x.SwapUsed
	}
	return 0
}

func (x *DynamicInfo) GetDiskTotal() uint64 {
	if x != nil {
		return x.DiskTotal
	}
	return 0
}

func (x *DynamicInfo) GetDiskUsed() uint64 {
	if x != nil {
		return x.DiskUsed
	}
	return 0
}

func (x *DynamicInfo) GetCpuUsage() float64 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *DynamicInfo) GetMemUsage() float64 {
	if x != nil {
		return x.MemUsage
	}
	return 0
}

func (x *DynamicInfo) GetSwapUsage() float64 {
	if x != nil {
		return x.SwapUsage
	}
	return 0
}

func (x *DynamicInfo) GetDiskUsage() float64 {
	if x != nil {
		return x.DiskUsage
	}
	return 0
}

func (x *DynamicInfo) GetLoad() *Load {
	if x != nil {
		return x.Load
	}
	return nil
}

func (x *DynamicInfo) GetDiskInfo() []*DiskInfo {
	if x != nil {
		return x.DiskInfo
	}
	return nil
}

func (x *DynamicInfo) GetNetSpeed() []*NetSpeed {
	if x != nil {
		return x.NetSpeed
	}
	return nil
}

func (x *DynamicInfo) GetVerify() *Verify {
	if x != nil {
		return x.Verify
	}
	return nil
}

func (x *DynamicInfo) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type NetDelayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CT     *NetDelay `protobuf:"bytes,1,opt,name=CT,proto3" json:"CT,omitempty"`
	CU     *NetDelay `protobuf:"bytes,2,opt,name=CU,proto3" json:"CU,omitempty"`
	CM     *NetDelay `protobuf:"bytes,3,opt,name=CM,proto3" json:"CM,omitempty"`
	Verify *Verify   `protobuf:"bytes,1000,opt,name=verify,proto3" json:"verify,omitempty"`
	Ts     int64     `protobuf:"varint,10001,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *NetDelayInfo) Reset() {
	*x = NetDelayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetDelayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetDelayInfo) ProtoMessage() {}

func (x *NetDelayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetDelayInfo.ProtoReflect.Descriptor instead.
func (*NetDelayInfo) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{2}
}

func (x *NetDelayInfo) GetCT() *NetDelay {
	if x != nil {
		return x.CT
	}
	return nil
}

func (x *NetDelayInfo) GetCU() *NetDelay {
	if x != nil {
		return x.CU
	}
	return nil
}

func (x *NetDelayInfo) GetCM() *NetDelay {
	if x != nil {
		return x.CM
	}
	return nil
}

func (x *NetDelayInfo) GetVerify() *Verify {
	if x != nil {
		return x.Verify
	}
	return nil
}

func (x *NetDelayInfo) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type ClientReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string  `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Verify *Verify `protobuf:"bytes,1000,opt,name=verify,proto3" json:"verify,omitempty"`
	Ts     int64   `protobuf:"varint,10001,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *ClientReport) Reset() {
	*x = ClientReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReport) ProtoMessage() {}

func (x *ClientReport) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReport.ProtoReflect.Descriptor instead.
func (*ClientReport) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{3}
}

func (x *ClientReport) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ClientReport) GetVerify() *Verify {
	if x != nil {
		return x.Verify
	}
	return nil
}

func (x *ClientReport) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type ResponseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseStatus) Reset() {
	*x = ResponseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatus) ProtoMessage() {}

func (x *ResponseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatus.ProtoReflect.Descriptor instead.
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseStatus) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DiskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Read  uint64 `protobuf:"varint,2,opt,name=read,proto3" json:"read,omitempty"`
	Write uint64 `protobuf:"varint,3,opt,name=write,proto3" json:"write,omitempty"`
}

func (x *DiskInfo) Reset() {
	*x = DiskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskInfo) ProtoMessage() {}

func (x *DiskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskInfo.ProtoReflect.Descriptor instead.
func (*DiskInfo) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{5}
}

func (x *DiskInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiskInfo) GetRead() uint64 {
	if x != nil {
		return x.Read
	}
	return 0
}

func (x *DiskInfo) GetWrite() uint64 {
	if x != nil {
		return x.Write
	}
	return 0
}

type NetSpeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	In   uint64 `protobuf:"varint,2,opt,name=In,proto3" json:"In,omitempty"`
	Out  uint64 `protobuf:"varint,3,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *NetSpeed) Reset() {
	*x = NetSpeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetSpeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetSpeed) ProtoMessage() {}

func (x *NetSpeed) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetSpeed.ProtoReflect.Descriptor instead.
func (*NetSpeed) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{6}
}

func (x *NetSpeed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetSpeed) GetIn() uint64 {
	if x != nil {
		return x.In
	}
	return 0
}

func (x *NetSpeed) GetOut() uint64 {
	if x != nil {
		return x.Out
	}
	return 0
}

type Load struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Load1  float64 `protobuf:"fixed64,1,opt,name=load1,proto3" json:"load1,omitempty"`
	Load5  float64 `protobuf:"fixed64,2,opt,name=load5,proto3" json:"load5,omitempty"`
	Load15 float64 `protobuf:"fixed64,3,opt,name=load15,proto3" json:"load15,omitempty"`
}

func (x *Load) Reset() {
	*x = Load{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Load) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Load) ProtoMessage() {}

func (x *Load) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Load.ProtoReflect.Descriptor instead.
func (*Load) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{7}
}

func (x *Load) GetLoad1() float64 {
	if x != nil {
		return x.Load1
	}
	return 0
}

func (x *Load) GetLoad5() float64 {
	if x != nil {
		return x.Load5
	}
	return 0
}

func (x *Load) GetLoad15() float64 {
	if x != nil {
		return x.Load15
	}
	return 0
}

type NetDelay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average uint64  `protobuf:"varint,1,opt,name=average,proto3" json:"average,omitempty"`
	Loss    float64 `protobuf:"fixed64,2,opt,name=loss,proto3" json:"loss,omitempty"`
}

func (x *NetDelay) Reset() {
	*x = NetDelay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetDelay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetDelay) ProtoMessage() {}

func (x *NetDelay) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetDelay.ProtoReflect.Descriptor instead.
func (*NetDelay) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{8}
}

func (x *NetDelay) GetAverage() uint64 {
	if x != nil {
		return x.Average
	}
	return 0
}

func (x *NetDelay) GetLoss() float64 {
	if x != nil {
		return x.Loss
	}
	return 0
}

type Verify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Verify) Reset() {
	*x = Verify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_SysInfo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verify) ProtoMessage() {}

func (x *Verify) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_SysInfo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verify.ProtoReflect.Descriptor instead.
func (*Verify) Descriptor() ([]byte, []int) {
	return file_src_proto_SysInfo_proto_rawDescGZIP(), []int{9}
}

func (x *Verify) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Verify) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_src_proto_SysInfo_proto protoreflect.FileDescriptor

var file_src_proto_SysInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x79, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x70, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x70,
	0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x70, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x79, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x06, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x12, 0x0f, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x91, 0x4e, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x74, 0x73, 0x22, 0x9d, 0x04, 0x0a, 0x0b, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x77, 0x61, 0x70, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x77, 0x61, 0x70, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x77, 0x61, 0x70, 0x55, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x73, 0x77, 0x61, 0x70, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x73, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x64,
	0x69, 0x73, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b,
	0x55, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b,
	0x55, 0x73, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x77, 0x61, 0x70, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x73, 0x77, 0x61, 0x70, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x73, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64,
	0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x08, 0x6e, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x66, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52,
	0x08, 0x6e, 0x65, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x79, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x06, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x12, 0x0f, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x91, 0x4e, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x74, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x02, 0x43, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x52, 0x02, 0x43, 0x54, 0x12, 0x21, 0x0a, 0x02, 0x43, 0x55, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4e,
	0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x02, 0x43, 0x55, 0x12, 0x21, 0x0a, 0x02, 0x43,
	0x4d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x02, 0x43, 0x4d, 0x12, 0x28,
	0x0a, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x52, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0f, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x91,
	0x4e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x22, 0x5b, 0x0a, 0x0c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x06, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x06, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0f, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x91, 0x4e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x48, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x72, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x40, 0x0a, 0x08, 0x4e, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x22, 0x4a, 0x0a, 0x04,
	0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x61, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x35,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35, 0x22, 0x38, 0x0a, 0x08, 0x4e, 0x65, 0x74, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f,
	0x73, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x32, 0x94, 0x02, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x73, 0x79,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x73,
	0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x17, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x16, 0x5a, 0x14, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2d, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_SysInfo_proto_rawDescOnce sync.Once
	file_src_proto_SysInfo_proto_rawDescData = file_src_proto_SysInfo_proto_rawDesc
)

func file_src_proto_SysInfo_proto_rawDescGZIP() []byte {
	file_src_proto_SysInfo_proto_rawDescOnce.Do(func() {
		file_src_proto_SysInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_SysInfo_proto_rawDescData)
	})
	return file_src_proto_SysInfo_proto_rawDescData
}

var file_src_proto_SysInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_src_proto_SysInfo_proto_goTypes = []interface{}{
	(*StaticInfo)(nil),     // 0: sysInfo.StaticInfo
	(*DynamicInfo)(nil),    // 1: sysInfo.DynamicInfo
	(*NetDelayInfo)(nil),   // 2: sysInfo.NetDelayInfo
	(*ClientReport)(nil),   // 3: sysInfo.ClientReport
	(*ResponseStatus)(nil), // 4: sysInfo.ResponseStatus
	(*DiskInfo)(nil),       // 5: sysInfo.DiskInfo
	(*NetSpeed)(nil),       // 6: sysInfo.NetSpeed
	(*Load)(nil),           // 7: sysInfo.Load
	(*NetDelay)(nil),       // 8: sysInfo.NetDelay
	(*Verify)(nil),         // 9: sysInfo.Verify
}
var file_src_proto_SysInfo_proto_depIdxs = []int32{
	9,  // 0: sysInfo.StaticInfo.verify:type_name -> sysInfo.Verify
	7,  // 1: sysInfo.DynamicInfo.load:type_name -> sysInfo.Load
	5,  // 2: sysInfo.DynamicInfo.diskInfo:type_name -> sysInfo.DiskInfo
	6,  // 3: sysInfo.DynamicInfo.netSpeed:type_name -> sysInfo.NetSpeed
	9,  // 4: sysInfo.DynamicInfo.verify:type_name -> sysInfo.Verify
	8,  // 5: sysInfo.NetDelayInfo.CT:type_name -> sysInfo.NetDelay
	8,  // 6: sysInfo.NetDelayInfo.CU:type_name -> sysInfo.NetDelay
	8,  // 7: sysInfo.NetDelayInfo.CM:type_name -> sysInfo.NetDelay
	9,  // 8: sysInfo.NetDelayInfo.verify:type_name -> sysInfo.Verify
	9,  // 9: sysInfo.ClientReport.verify:type_name -> sysInfo.Verify
	0,  // 10: sysInfo.MonitorService.StaticService:input_type -> sysInfo.StaticInfo
	1,  // 11: sysInfo.MonitorService.DynamicService:input_type -> sysInfo.DynamicInfo
	2,  // 12: sysInfo.MonitorService.NetDelayService:input_type -> sysInfo.NetDelayInfo
	3,  // 13: sysInfo.MonitorService.reportService:input_type -> sysInfo.ClientReport
	4,  // 14: sysInfo.MonitorService.StaticService:output_type -> sysInfo.ResponseStatus
	4,  // 15: sysInfo.MonitorService.DynamicService:output_type -> sysInfo.ResponseStatus
	4,  // 16: sysInfo.MonitorService.NetDelayService:output_type -> sysInfo.ResponseStatus
	4,  // 17: sysInfo.MonitorService.reportService:output_type -> sysInfo.ResponseStatus
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_src_proto_SysInfo_proto_init() }
func file_src_proto_SysInfo_proto_init() {
	if File_src_proto_SysInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_SysInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetDelayInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetSpeed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Load); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetDelay); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_SysInfo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verify); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_proto_SysInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_SysInfo_proto_goTypes,
		DependencyIndexes: file_src_proto_SysInfo_proto_depIdxs,
		MessageInfos:      file_src_proto_SysInfo_proto_msgTypes,
	}.Build()
	File_src_proto_SysInfo_proto = out.File
	file_src_proto_SysInfo_proto_rawDesc = nil
	file_src_proto_SysInfo_proto_goTypes = nil
	file_src_proto_SysInfo_proto_depIdxs = nil
}
