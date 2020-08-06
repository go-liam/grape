// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package region

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SayHelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloReq) Reset()         { *m = SayHelloReq{} }
func (m *SayHelloReq) String() string { return proto.CompactTextString(m) }
func (*SayHelloReq) ProtoMessage()    {}
func (*SayHelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *SayHelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloReq.Unmarshal(m, b)
}
func (m *SayHelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloReq.Marshal(b, m, deterministic)
}
func (m *SayHelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloReq.Merge(m, src)
}
func (m *SayHelloReq) XXX_Size() int {
	return xxx_messageInfo_SayHelloReq.Size(m)
}
func (m *SayHelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloReq proto.InternalMessageInfo

func (m *SayHelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SayHelloResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloResp) Reset()         { *m = SayHelloResp{} }
func (m *SayHelloResp) String() string { return proto.CompactTextString(m) }
func (*SayHelloResp) ProtoMessage()    {}
func (*SayHelloResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *SayHelloResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloResp.Unmarshal(m, b)
}
func (m *SayHelloResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloResp.Marshal(b, m, deterministic)
}
func (m *SayHelloResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResp.Merge(m, src)
}
func (m *SayHelloResp) XXX_Size() int {
	return xxx_messageInfo_SayHelloResp.Size(m)
}
func (m *SayHelloResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResp.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResp proto.InternalMessageInfo

func (m *SayHelloResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SayHelloResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HealthReq struct {
	ClientName           string   `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthReq) Reset()         { *m = HealthReq{} }
func (m *HealthReq) String() string { return proto.CompactTextString(m) }
func (*HealthReq) ProtoMessage()    {}
func (*HealthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *HealthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthReq.Unmarshal(m, b)
}
func (m *HealthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthReq.Marshal(b, m, deterministic)
}
func (m *HealthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthReq.Merge(m, src)
}
func (m *HealthReq) XXX_Size() int {
	return xxx_messageInfo_HealthReq.Size(m)
}
func (m *HealthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthReq.DiscardUnknown(m)
}

var xxx_messageInfo_HealthReq proto.InternalMessageInfo

func (m *HealthReq) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type HealthResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	ServerName           string   `protobuf:"bytes,2,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResp) Reset()         { *m = HealthResp{} }
func (m *HealthResp) String() string { return proto.CompactTextString(m) }
func (*HealthResp) ProtoMessage()    {}
func (*HealthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{3}
}

func (m *HealthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResp.Unmarshal(m, b)
}
func (m *HealthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResp.Marshal(b, m, deterministic)
}
func (m *HealthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResp.Merge(m, src)
}
func (m *HealthResp) XXX_Size() int {
	return xxx_messageInfo_HealthResp.Size(m)
}
func (m *HealthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResp.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResp proto.InternalMessageInfo

func (m *HealthResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HealthResp) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *HealthResp) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *HealthResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloReq)(nil), "region.SayHelloReq")
	proto.RegisterType((*SayHelloResp)(nil), "region.SayHelloResp")
	proto.RegisterType((*HealthReq)(nil), "region.HealthReq")
	proto.RegisterType((*HealthResp)(nil), "region.HealthResp")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x55, 0x28, 0x81, 0xbe, 0xb2, 0x70, 0x30, 0x44, 0x1d, 0x2a, 0xc8, 0x84, 0x84, 0x14,
	0x24, 0x10, 0x1b, 0x23, 0x43, 0x27, 0x86, 0xf4, 0x17, 0x84, 0xf4, 0x54, 0x22, 0x25, 0x3e, 0x63,
	0x9b, 0x48, 0xfc, 0x7b, 0x14, 0x1b, 0x13, 0x03, 0x62, 0xbb, 0x7b, 0x79, 0xb9, 0xf7, 0xdd, 0x19,
	0xd8, 0x71, 0x2f, 0xa5, 0x36, 0xe2, 0x84, 0x32, 0xc3, 0xfb, 0x56, 0x54, 0x71, 0x85, 0xe5, 0xb6,
	0xfe, 0xd8, 0x70, 0xd7, 0x49, 0xc5, 0x6f, 0x44, 0x98, 0xab, 0xba, 0xe7, 0x7c, 0x76, 0x39, 0xbb,
	0x5e, 0x54, 0xbe, 0x2e, 0x1e, 0x71, 0x3a, 0x59, 0xac, 0x1e, 0x3d, 0x8d, 0xec, 0x82, 0xe7, 0xa8,
	0xf2, 0x35, 0xe5, 0x38, 0xee, 0xd9, 0xda, 0x7a, 0xcf, 0xf9, 0x81, 0xff, 0x35, 0xb6, 0xc5, 0x0d,
	0x16, 0x1b, 0xae, 0x3b, 0xf7, 0x3a, 0x8e, 0x5f, 0x03, 0x4d, 0xd7, 0xb2, 0x72, 0xcf, 0x53, 0x48,
	0xa2, 0x14, 0x0e, 0x88, 0xe6, 0x7f, 0x82, 0xd6, 0x80, 0x65, 0x33, 0xb0, 0xf1, 0x13, 0x42, 0x56,
	0xa2, 0x8c, 0x20, 0x03, 0x1b, 0xdb, 0x8a, 0xca, 0x0f, 0x03, 0xc8, 0x57, 0x9b, 0x22, 0xce, 0x7f,
	0x20, 0xde, 0xbd, 0x63, 0xf9, 0xc4, 0xbd, 0x6c, 0xd9, 0x0c, 0x6d, 0xc3, 0xf4, 0x80, 0x93, 0xb8,
	0x2f, 0x9d, 0x97, 0xe1, 0x4e, 0x65, 0x72, 0xa4, 0xd5, 0xc5, 0x5f, 0xd1, 0x6a, 0xba, 0x45, 0x16,
	0xd8, 0xe9, 0x2c, 0x7e, 0xff, 0x5e, 0x7c, 0x45, 0xbf, 0x25, 0xab, 0x5f, 0x32, 0xff, 0x12, 0xf7,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xeb, 0x58, 0xed, 0x97, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloResp, error)
	Health(ctx context.Context, in *HealthReq, opts ...grpc.CallOption) (*HealthResp, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloResp, error) {
	out := new(SayHelloResp)
	err := c.cc.Invoke(ctx, "/region.DemoService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) Health(ctx context.Context, in *HealthReq, opts ...grpc.CallOption) (*HealthResp, error) {
	out := new(HealthResp)
	err := c.cc.Invoke(ctx, "/region.DemoService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	SayHello(context.Context, *SayHelloReq) (*SayHelloResp, error)
	Health(context.Context, *HealthReq) (*HealthResp, error)
}

// UnimplementedDemoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (*UnimplementedDemoServiceServer) SayHello(ctx context.Context, req *SayHelloReq) (*SayHelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedDemoServiceServer) Health(ctx context.Context, req *HealthReq) (*HealthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.DemoService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).SayHello(ctx, req.(*SayHelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.DemoService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Health(ctx, req.(*HealthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "region.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _DemoService_SayHello_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _DemoService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
