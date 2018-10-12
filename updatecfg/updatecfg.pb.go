// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatecfg.proto

package updatecfg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TraceType int32

const (
	TraceType_DUMP_DROP      TraceType = 0
	TraceType_DUMP_TRANSLATE TraceType = 1
	TraceType_DUMP_KNI       TraceType = 2
)

var TraceType_name = map[int32]string{
	0: "DUMP_DROP",
	1: "DUMP_TRANSLATE",
	2: "DUMP_KNI",
}
var TraceType_value = map[string]int32{
	"DUMP_DROP":      0,
	"DUMP_TRANSLATE": 1,
	"DUMP_KNI":       2,
}

func (x TraceType) String() string {
	return proto.EnumName(TraceType_name, int32(x))
}
func (TraceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{0}
}

type Protocol int32

const (
	Protocol_UNKNOWN   Protocol = 0
	Protocol_TCP       Protocol = 6
	Protocol_UDP       Protocol = 17
	Protocol_IPv6_Flag Protocol = 65536
	Protocol_TCP6      Protocol = 65542
	Protocol_UDP6      Protocol = 65553
)

var Protocol_name = map[int32]string{
	0:     "UNKNOWN",
	6:     "TCP",
	17:    "UDP",
	65536: "IPv6_Flag",
	65542: "TCP6",
	65553: "UDP6",
}
var Protocol_value = map[string]int32{
	"UNKNOWN":   0,
	"TCP":       6,
	"UDP":       17,
	"IPv6_Flag": 65536,
	"TCP6":      65542,
	"UDP6":      65553,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{1}
}

type DumpControlRequest struct {
	EnableTrace          bool      `protobuf:"varint,1,opt,name=enable_trace,json=enableTrace,proto3" json:"enable_trace,omitempty"`
	TraceType            TraceType `protobuf:"varint,2,opt,name=trace_type,json=traceType,proto3,enum=updatecfg.TraceType" json:"trace_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DumpControlRequest) Reset()         { *m = DumpControlRequest{} }
func (m *DumpControlRequest) String() string { return proto.CompactTextString(m) }
func (*DumpControlRequest) ProtoMessage()    {}
func (*DumpControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{0}
}
func (m *DumpControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpControlRequest.Unmarshal(m, b)
}
func (m *DumpControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpControlRequest.Marshal(b, m, deterministic)
}
func (dst *DumpControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpControlRequest.Merge(dst, src)
}
func (m *DumpControlRequest) XXX_Size() int {
	return xxx_messageInfo_DumpControlRequest.Size(m)
}
func (m *DumpControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpControlRequest proto.InternalMessageInfo

func (m *DumpControlRequest) GetEnableTrace() bool {
	if m != nil {
		return m.EnableTrace
	}
	return false
}

func (m *DumpControlRequest) GetTraceType() TraceType {
	if m != nil {
		return m.TraceType
	}
	return TraceType_DUMP_DROP
}

type IPAddress struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPAddress) Reset()         { *m = IPAddress{} }
func (m *IPAddress) String() string { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()    {}
func (*IPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{1}
}
func (m *IPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddress.Unmarshal(m, b)
}
func (m *IPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddress.Marshal(b, m, deterministic)
}
func (dst *IPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddress.Merge(dst, src)
}
func (m *IPAddress) XXX_Size() int {
	return xxx_messageInfo_IPAddress.Size(m)
}
func (m *IPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddress proto.InternalMessageInfo

func (m *IPAddress) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type Subnet struct {
	Address              *IPAddress `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MaskBitsNumber       uint32     `protobuf:"varint,2,opt,name=mask_bits_number,json=maskBitsNumber,proto3" json:"mask_bits_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Subnet) Reset()         { *m = Subnet{} }
func (m *Subnet) String() string { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()    {}
func (*Subnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{2}
}
func (m *Subnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subnet.Unmarshal(m, b)
}
func (m *Subnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subnet.Marshal(b, m, deterministic)
}
func (dst *Subnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subnet.Merge(dst, src)
}
func (m *Subnet) XXX_Size() int {
	return xxx_messageInfo_Subnet.Size(m)
}
func (m *Subnet) XXX_DiscardUnknown() {
	xxx_messageInfo_Subnet.DiscardUnknown(m)
}

var xxx_messageInfo_Subnet proto.InternalMessageInfo

func (m *Subnet) GetAddress() *IPAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Subnet) GetMaskBitsNumber() uint32 {
	if m != nil {
		return m.MaskBitsNumber
	}
	return 0
}

type InterfaceAddressChangeRequest struct {
	InterfaceId          uint32   `protobuf:"varint,1,opt,name=interface_id,json=interfaceId,proto3" json:"interface_id,omitempty"`
	PortSubnet           *Subnet  `protobuf:"bytes,2,opt,name=port_subnet,json=portSubnet,proto3" json:"port_subnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceAddressChangeRequest) Reset()         { *m = InterfaceAddressChangeRequest{} }
func (m *InterfaceAddressChangeRequest) String() string { return proto.CompactTextString(m) }
func (*InterfaceAddressChangeRequest) ProtoMessage()    {}
func (*InterfaceAddressChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{3}
}
func (m *InterfaceAddressChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceAddressChangeRequest.Unmarshal(m, b)
}
func (m *InterfaceAddressChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceAddressChangeRequest.Marshal(b, m, deterministic)
}
func (dst *InterfaceAddressChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceAddressChangeRequest.Merge(dst, src)
}
func (m *InterfaceAddressChangeRequest) XXX_Size() int {
	return xxx_messageInfo_InterfaceAddressChangeRequest.Size(m)
}
func (m *InterfaceAddressChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceAddressChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceAddressChangeRequest proto.InternalMessageInfo

func (m *InterfaceAddressChangeRequest) GetInterfaceId() uint32 {
	if m != nil {
		return m.InterfaceId
	}
	return 0
}

func (m *InterfaceAddressChangeRequest) GetPortSubnet() *Subnet {
	if m != nil {
		return m.PortSubnet
	}
	return nil
}

type ForwardedPort struct {
	SourcePortNumber     uint32     `protobuf:"varint,1,opt,name=source_port_number,json=sourcePortNumber,proto3" json:"source_port_number,omitempty"`
	TargetAddress        *IPAddress `protobuf:"bytes,2,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	TargetPortNumber     uint32     `protobuf:"varint,3,opt,name=target_port_number,json=targetPortNumber,proto3" json:"target_port_number,omitempty"`
	Protocol             Protocol   `protobuf:"varint,4,opt,name=protocol,proto3,enum=updatecfg.Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ForwardedPort) Reset()         { *m = ForwardedPort{} }
func (m *ForwardedPort) String() string { return proto.CompactTextString(m) }
func (*ForwardedPort) ProtoMessage()    {}
func (*ForwardedPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{4}
}
func (m *ForwardedPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardedPort.Unmarshal(m, b)
}
func (m *ForwardedPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardedPort.Marshal(b, m, deterministic)
}
func (dst *ForwardedPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardedPort.Merge(dst, src)
}
func (m *ForwardedPort) XXX_Size() int {
	return xxx_messageInfo_ForwardedPort.Size(m)
}
func (m *ForwardedPort) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardedPort.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardedPort proto.InternalMessageInfo

func (m *ForwardedPort) GetSourcePortNumber() uint32 {
	if m != nil {
		return m.SourcePortNumber
	}
	return 0
}

func (m *ForwardedPort) GetTargetAddress() *IPAddress {
	if m != nil {
		return m.TargetAddress
	}
	return nil
}

func (m *ForwardedPort) GetTargetPortNumber() uint32 {
	if m != nil {
		return m.TargetPortNumber
	}
	return 0
}

func (m *ForwardedPort) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_UNKNOWN
}

type PortForwardingChangeRequest struct {
	EnableForwarding     bool           `protobuf:"varint,1,opt,name=enable_forwarding,json=enableForwarding,proto3" json:"enable_forwarding,omitempty"`
	InterfaceId          uint32         `protobuf:"varint,2,opt,name=interface_id,json=interfaceId,proto3" json:"interface_id,omitempty"`
	Port                 *ForwardedPort `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PortForwardingChangeRequest) Reset()         { *m = PortForwardingChangeRequest{} }
func (m *PortForwardingChangeRequest) String() string { return proto.CompactTextString(m) }
func (*PortForwardingChangeRequest) ProtoMessage()    {}
func (*PortForwardingChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{5}
}
func (m *PortForwardingChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortForwardingChangeRequest.Unmarshal(m, b)
}
func (m *PortForwardingChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortForwardingChangeRequest.Marshal(b, m, deterministic)
}
func (dst *PortForwardingChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortForwardingChangeRequest.Merge(dst, src)
}
func (m *PortForwardingChangeRequest) XXX_Size() int {
	return xxx_messageInfo_PortForwardingChangeRequest.Size(m)
}
func (m *PortForwardingChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PortForwardingChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PortForwardingChangeRequest proto.InternalMessageInfo

func (m *PortForwardingChangeRequest) GetEnableForwarding() bool {
	if m != nil {
		return m.EnableForwarding
	}
	return false
}

func (m *PortForwardingChangeRequest) GetInterfaceId() uint32 {
	if m != nil {
		return m.InterfaceId
	}
	return 0
}

func (m *PortForwardingChangeRequest) GetPort() *ForwardedPort {
	if m != nil {
		return m.Port
	}
	return nil
}

type Reply struct {
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_updatecfg_fa7d68b5067fd899, []int{6}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*DumpControlRequest)(nil), "updatecfg.DumpControlRequest")
	proto.RegisterType((*IPAddress)(nil), "updatecfg.IPAddress")
	proto.RegisterType((*Subnet)(nil), "updatecfg.Subnet")
	proto.RegisterType((*InterfaceAddressChangeRequest)(nil), "updatecfg.InterfaceAddressChangeRequest")
	proto.RegisterType((*ForwardedPort)(nil), "updatecfg.ForwardedPort")
	proto.RegisterType((*PortForwardingChangeRequest)(nil), "updatecfg.PortForwardingChangeRequest")
	proto.RegisterType((*Reply)(nil), "updatecfg.Reply")
	proto.RegisterEnum("updatecfg.TraceType", TraceType_name, TraceType_value)
	proto.RegisterEnum("updatecfg.Protocol", Protocol_name, Protocol_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpdaterClient is the client API for Updater service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdaterClient interface {
	ControlDump(ctx context.Context, in *DumpControlRequest, opts ...grpc.CallOption) (*Reply, error)
	ChangeInterfaceAddress(ctx context.Context, in *InterfaceAddressChangeRequest, opts ...grpc.CallOption) (*Reply, error)
	ChangePortForwarding(ctx context.Context, in *PortForwardingChangeRequest, opts ...grpc.CallOption) (*Reply, error)
}

type updaterClient struct {
	cc *grpc.ClientConn
}

func NewUpdaterClient(cc *grpc.ClientConn) UpdaterClient {
	return &updaterClient{cc}
}

func (c *updaterClient) ControlDump(ctx context.Context, in *DumpControlRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/updatecfg.Updater/ControlDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updaterClient) ChangeInterfaceAddress(ctx context.Context, in *InterfaceAddressChangeRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/updatecfg.Updater/ChangeInterfaceAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updaterClient) ChangePortForwarding(ctx context.Context, in *PortForwardingChangeRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/updatecfg.Updater/ChangePortForwarding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdaterServer is the server API for Updater service.
type UpdaterServer interface {
	ControlDump(context.Context, *DumpControlRequest) (*Reply, error)
	ChangeInterfaceAddress(context.Context, *InterfaceAddressChangeRequest) (*Reply, error)
	ChangePortForwarding(context.Context, *PortForwardingChangeRequest) (*Reply, error)
}

func RegisterUpdaterServer(s *grpc.Server, srv UpdaterServer) {
	s.RegisterService(&_Updater_serviceDesc, srv)
}

func _Updater_ControlDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdaterServer).ControlDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updatecfg.Updater/ControlDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdaterServer).ControlDump(ctx, req.(*DumpControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Updater_ChangeInterfaceAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterfaceAddressChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdaterServer).ChangeInterfaceAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updatecfg.Updater/ChangeInterfaceAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdaterServer).ChangeInterfaceAddress(ctx, req.(*InterfaceAddressChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Updater_ChangePortForwarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortForwardingChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdaterServer).ChangePortForwarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updatecfg.Updater/ChangePortForwarding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdaterServer).ChangePortForwarding(ctx, req.(*PortForwardingChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Updater_serviceDesc = grpc.ServiceDesc{
	ServiceName: "updatecfg.Updater",
	HandlerType: (*UpdaterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ControlDump",
			Handler:    _Updater_ControlDump_Handler,
		},
		{
			MethodName: "ChangeInterfaceAddress",
			Handler:    _Updater_ChangeInterfaceAddress_Handler,
		},
		{
			MethodName: "ChangePortForwarding",
			Handler:    _Updater_ChangePortForwarding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "updatecfg.proto",
}

func init() { proto.RegisterFile("updatecfg.proto", fileDescriptor_updatecfg_fa7d68b5067fd899) }

var fileDescriptor_updatecfg_fa7d68b5067fd899 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xad, 0xd3, 0x7c, 0x4d, 0x32, 0x4e, 0xd2, 0xed, 0x7c, 0x15, 0x0a, 0xa0, 0x4a, 0xc5, 0x12,
	0x28, 0x2a, 0x55, 0x91, 0x52, 0xa9, 0x17, 0x38, 0xd0, 0x26, 0x54, 0x8a, 0x0a, 0xae, 0xb5, 0x4d,
	0xe0, 0x68, 0xd9, 0xf1, 0xc6, 0x44, 0x38, 0xb6, 0x59, 0xaf, 0x8b, 0x7a, 0xeb, 0x89, 0x3b, 0x67,
	0x7e, 0x19, 0xbf, 0x84, 0x2b, 0xf2, 0xae, 0xed, 0x3a, 0xa4, 0xf4, 0x36, 0x9e, 0x79, 0x3b, 0x6f,
	0xe6, 0xf9, 0x0d, 0x6c, 0xa7, 0xb1, 0xe7, 0x08, 0x36, 0x9b, 0xfb, 0x47, 0x31, 0x8f, 0x44, 0x84,
	0xad, 0x32, 0x61, 0x04, 0x80, 0xa3, 0x74, 0x19, 0x0f, 0xa3, 0x50, 0xf0, 0x28, 0xa0, 0xec, 0x6b,
	0xca, 0x12, 0x81, 0xcf, 0xa0, 0xcd, 0x42, 0xc7, 0x0d, 0x98, 0x2d, 0xb8, 0x33, 0x63, 0x3d, 0x6d,
	0x5f, 0xeb, 0x37, 0xa9, 0xae, 0x72, 0x93, 0x2c, 0x85, 0xc7, 0x00, 0xb2, 0x66, 0x8b, 0x9b, 0x98,
	0xf5, 0x6a, 0xfb, 0x5a, 0xbf, 0x3b, 0xd8, 0x3d, 0xba, 0x63, 0x92, 0xa8, 0xc9, 0x4d, 0xcc, 0x68,
	0x4b, 0x14, 0xa1, 0xf1, 0x1c, 0x5a, 0x63, 0xeb, 0xd4, 0xf3, 0x38, 0x4b, 0x12, 0xec, 0x41, 0xc3,
	0x51, 0xa1, 0xec, 0xdf, 0xa6, 0xc5, 0xa7, 0xe1, 0xc2, 0xd6, 0x55, 0xea, 0x86, 0x4c, 0xe0, 0xd1,
	0x2a, 0x46, 0x5f, 0xa1, 0x28, 0x5b, 0x95, 0x2f, 0xb1, 0x0f, 0x64, 0xe9, 0x24, 0x5f, 0x6c, 0x77,
	0x21, 0x12, 0x3b, 0x4c, 0x97, 0x2e, 0xe3, 0x72, 0xb6, 0x0e, 0xed, 0x66, 0xf9, 0xb3, 0x85, 0x48,
	0x4c, 0x99, 0x35, 0xae, 0x61, 0x6f, 0x1c, 0x0a, 0xc6, 0xe7, 0xce, 0x8c, 0xe5, 0x6d, 0x86, 0x9f,
	0x9d, 0xd0, 0x67, 0x15, 0x0d, 0x16, 0x05, 0xc0, 0x5e, 0x78, 0x92, 0xbf, 0x43, 0xf5, 0x32, 0x37,
	0xf6, 0x70, 0x00, 0x7a, 0x1c, 0x71, 0x61, 0x27, 0x72, 0x58, 0x49, 0xa4, 0x0f, 0x76, 0x2a, 0x13,
	0xaa, 0x2d, 0x28, 0x64, 0x28, 0x15, 0x1b, 0xbf, 0x34, 0xe8, 0x9c, 0x47, 0xfc, 0x9b, 0xc3, 0x3d,
	0xe6, 0x59, 0x11, 0x17, 0x78, 0x08, 0x98, 0x44, 0x29, 0x9f, 0x31, 0x5b, 0x36, 0xcb, 0xa7, 0x56,
	0x74, 0x44, 0x55, 0x32, 0x9c, 0x9a, 0x1b, 0x5f, 0x43, 0x57, 0x38, 0xdc, 0x67, 0xc2, 0x2e, 0x84,
	0xa9, 0x3d, 0x20, 0x4c, 0x47, 0x61, 0x0b, 0xc9, 0x0f, 0x01, 0xf3, 0xc7, 0x55, 0xaa, 0x4d, 0x45,
	0xa5, 0x2a, 0x15, 0xaa, 0x57, 0xd0, 0x94, 0x7e, 0x99, 0x45, 0x41, 0xaf, 0x2e, 0x7f, 0xf0, 0xff,
	0x15, 0x12, 0x2b, 0x2f, 0xd1, 0x12, 0x64, 0xfc, 0xd4, 0xe0, 0x69, 0xf6, 0x3e, 0xdf, 0x6f, 0x11,
	0xfa, 0xab, 0x92, 0xbe, 0x84, 0x9d, 0xdc, 0x56, 0xf3, 0x12, 0x91, 0x7b, 0x8b, 0xa8, 0xc2, 0xdd,
	0xcb, 0x35, 0xfd, 0x6b, 0xeb, 0xfa, 0x1f, 0x42, 0x3d, 0xdb, 0x43, 0x2e, 0xa0, 0x0f, 0x7a, 0x95,
	0xe1, 0x56, 0x14, 0xa6, 0x12, 0x65, 0x3c, 0x86, 0xff, 0x28, 0x8b, 0x83, 0x1b, 0x24, 0xb0, 0xb9,
	0x4c, 0x7c, 0xd9, 0xb0, 0x45, 0xb3, 0xf0, 0xe0, 0x0d, 0xb4, 0x4a, 0xbf, 0x62, 0x07, 0x5a, 0xa3,
	0xe9, 0x07, 0xcb, 0x1e, 0xd1, 0x4b, 0x8b, 0x6c, 0x20, 0x42, 0x57, 0x7e, 0x4e, 0xe8, 0xa9, 0x79,
	0xf5, 0xfe, 0x74, 0xf2, 0x8e, 0x68, 0xd8, 0x86, 0xa6, 0xcc, 0x5d, 0x98, 0x63, 0x52, 0x3b, 0xa0,
	0xd0, 0x2c, 0xc4, 0x40, 0x1d, 0x1a, 0x53, 0xf3, 0xc2, 0xbc, 0xfc, 0x64, 0x92, 0x0d, 0x6c, 0xc0,
	0xe6, 0x64, 0x68, 0x91, 0xad, 0x2c, 0x98, 0x8e, 0x2c, 0xb2, 0x83, 0xdb, 0xd9, 0x01, 0x5c, 0x9f,
	0xd8, 0xe7, 0x81, 0xe3, 0x93, 0xdb, 0xdb, 0x3a, 0x02, 0xd4, 0x27, 0x43, 0xeb, 0x84, 0x7c, 0x57,
	0xf1, 0x74, 0x64, 0x9d, 0x90, 0x1f, 0xb7, 0xf5, 0xc1, 0x6f, 0x0d, 0x1a, 0x53, 0xb9, 0x0e, 0xc7,
	0xb7, 0xa0, 0xe7, 0xf7, 0x99, 0x9d, 0x2a, 0xee, 0x55, 0xf6, 0x5c, 0xbf, 0xdd, 0x27, 0xa4, 0x52,
	0x96, 0xfb, 0x1a, 0x1b, 0xf8, 0x11, 0x1e, 0xa9, 0x3f, 0xf1, 0xb7, 0xe5, 0xb1, 0x5f, 0xb5, 0xcd,
	0x43, 0xf7, 0x70, 0x6f, 0x5f, 0x0a, 0xbb, 0x0a, 0xb4, 0xfa, 0xd7, 0xf1, 0x45, 0xd5, 0x27, 0xff,
	0x36, 0xc4, 0x7d, 0x3d, 0xcf, 0xc8, 0x59, 0x5b, 0x2d, 0x6e, 0x3a, 0x62, 0x38, 0xf7, 0x2d, 0xcd,
	0xdd, 0x92, 0x06, 0x3b, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xe4, 0x11, 0xa6, 0xc8, 0x04,
	0x00, 0x00,
}