// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notice.proto

package notice

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
	return fileDescriptor_642492014393dbdb, []int{0}
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
	return fileDescriptor_642492014393dbdb, []int{1}
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

type EmailReq struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	FromName             string   `protobuf:"bytes,2,opt,name=fromName,proto3" json:"fromName,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Body                 string   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailReq) Reset()         { *m = EmailReq{} }
func (m *EmailReq) String() string { return proto.CompactTextString(m) }
func (*EmailReq) ProtoMessage()    {}
func (*EmailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_642492014393dbdb, []int{2}
}

func (m *EmailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailReq.Unmarshal(m, b)
}
func (m *EmailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailReq.Marshal(b, m, deterministic)
}
func (m *EmailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailReq.Merge(m, src)
}
func (m *EmailReq) XXX_Size() int {
	return xxx_messageInfo_EmailReq.Size(m)
}
func (m *EmailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailReq.DiscardUnknown(m)
}

var xxx_messageInfo_EmailReq proto.InternalMessageInfo

func (m *EmailReq) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *EmailReq) GetFromName() string {
	if m != nil {
		return m.FromName
	}
	return ""
}

func (m *EmailReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EmailReq) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type EmailResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailResp) Reset()         { *m = EmailResp{} }
func (m *EmailResp) String() string { return proto.CompactTextString(m) }
func (*EmailResp) ProtoMessage()    {}
func (*EmailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_642492014393dbdb, []int{3}
}

func (m *EmailResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailResp.Unmarshal(m, b)
}
func (m *EmailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailResp.Marshal(b, m, deterministic)
}
func (m *EmailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailResp.Merge(m, src)
}
func (m *EmailResp) XXX_Size() int {
	return xxx_messageInfo_EmailResp.Size(m)
}
func (m *EmailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailResp.DiscardUnknown(m)
}

var xxx_messageInfo_EmailResp proto.InternalMessageInfo

func (m *EmailResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EmailResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DingTalkReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RobotSecret          string   `protobuf:"bytes,2,opt,name=robotSecret,proto3" json:"robotSecret,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	IsAtAll              bool     `protobuf:"varint,5,opt,name=isAtAll,proto3" json:"isAtAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DingTalkReq) Reset()         { *m = DingTalkReq{} }
func (m *DingTalkReq) String() string { return proto.CompactTextString(m) }
func (*DingTalkReq) ProtoMessage()    {}
func (*DingTalkReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_642492014393dbdb, []int{4}
}

func (m *DingTalkReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DingTalkReq.Unmarshal(m, b)
}
func (m *DingTalkReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DingTalkReq.Marshal(b, m, deterministic)
}
func (m *DingTalkReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DingTalkReq.Merge(m, src)
}
func (m *DingTalkReq) XXX_Size() int {
	return xxx_messageInfo_DingTalkReq.Size(m)
}
func (m *DingTalkReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DingTalkReq.DiscardUnknown(m)
}

var xxx_messageInfo_DingTalkReq proto.InternalMessageInfo

func (m *DingTalkReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DingTalkReq) GetRobotSecret() string {
	if m != nil {
		return m.RobotSecret
	}
	return ""
}

func (m *DingTalkReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DingTalkReq) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *DingTalkReq) GetIsAtAll() bool {
	if m != nil {
		return m.IsAtAll
	}
	return false
}

type DingTalkResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DingTalkResp) Reset()         { *m = DingTalkResp{} }
func (m *DingTalkResp) String() string { return proto.CompactTextString(m) }
func (*DingTalkResp) ProtoMessage()    {}
func (*DingTalkResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_642492014393dbdb, []int{5}
}

func (m *DingTalkResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DingTalkResp.Unmarshal(m, b)
}
func (m *DingTalkResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DingTalkResp.Marshal(b, m, deterministic)
}
func (m *DingTalkResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DingTalkResp.Merge(m, src)
}
func (m *DingTalkResp) XXX_Size() int {
	return xxx_messageInfo_DingTalkResp.Size(m)
}
func (m *DingTalkResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DingTalkResp.DiscardUnknown(m)
}

var xxx_messageInfo_DingTalkResp proto.InternalMessageInfo

func (m *DingTalkResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DingTalkResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthReq)(nil), "notice.HealthReq")
	proto.RegisterType((*HealthResp)(nil), "notice.HealthResp")
	proto.RegisterType((*EmailReq)(nil), "notice.EmailReq")
	proto.RegisterType((*EmailResp)(nil), "notice.EmailResp")
	proto.RegisterType((*DingTalkReq)(nil), "notice.DingTalkReq")
	proto.RegisterType((*DingTalkResp)(nil), "notice.DingTalkResp")
}

func init() { proto.RegisterFile("notice.proto", fileDescriptor_642492014393dbdb) }

var fileDescriptor_642492014393dbdb = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x95, 0x5d, 0x4c, 0x61, 0xa0, 0x15, 0x6c, 0x39, 0x58, 0x3e, 0x54, 0xc8, 0xa7, 0x4a, 0xad,
	0xa8, 0xd4, 0x9e, 0x2a, 0xb5, 0x07, 0xa4, 0x56, 0xea, 0xa5, 0x1c, 0x4c, 0xff, 0x80, 0x6d, 0xa6,
	0x74, 0x83, 0xed, 0x75, 0x76, 0x27, 0x04, 0x7e, 0x42, 0x7e, 0x4c, 0xfe, 0x63, 0xb4, 0x5f, 0xe0,
	0x10, 0xe5, 0x90, 0xdb, 0xbc, 0x37, 0xe3, 0x99, 0xf7, 0xd6, 0x0f, 0xc6, 0x8d, 0x20, 0x5e, 0xe2,
	0xa2, 0x95, 0x82, 0x04, 0xeb, 0x5b, 0x94, 0x7e, 0x84, 0xe1, 0x6f, 0xcc, 0x2b, 0xfa, 0x9f, 0xe1,
	0x35, 0x7b, 0x0f, 0x50, 0x56, 0x1c, 0x1b, 0x5a, 0xe5, 0x35, 0xc6, 0xc1, 0x3c, 0xf8, 0x30, 0xcc,
	0x3a, 0x4c, 0x4a, 0x00, 0x7e, 0x58, 0xb5, 0x8c, 0x41, 0xaf, 0x14, 0x1b, 0x3b, 0x17, 0x65, 0xa6,
	0xd6, 0x1b, 0x14, 0xca, 0x3d, 0x4a, 0xb3, 0x21, 0xb4, 0x1b, 0xce, 0x0c, 0x8b, 0xe1, 0xf5, 0x1e,
	0xa5, 0xe2, 0xa2, 0x89, 0x5f, 0x99, 0xa6, 0x87, 0xba, 0x53, 0xa3, 0x52, 0xf9, 0x16, 0xe3, 0x9e,
	0xed, 0x38, 0x98, 0x1e, 0x60, 0xf0, 0xab, 0xce, 0x79, 0xa5, 0x15, 0xbe, 0x85, 0x90, 0x84, 0x53,
	0x16, 0x92, 0x60, 0x09, 0x0c, 0xfe, 0x49, 0x51, 0x77, 0xae, 0x9d, 0xb0, 0xde, 0xa8, 0x6e, 0x8a,
	0x2b, 0x2c, 0xc9, 0xdf, 0x72, 0x50, 0x2b, 0xa7, 0x63, 0xeb, 0x0f, 0x99, 0x5a, 0x73, 0x85, 0xd8,
	0x1c, 0xe3, 0xc8, 0x72, 0xba, 0x4e, 0xbf, 0xc1, 0xd0, 0x5d, 0x7e, 0xc6, 0x6e, 0x47, 0x74, 0xf8,
	0x58, 0xf4, 0x5d, 0x00, 0xa3, 0x9f, 0xbc, 0xd9, 0xfe, 0xcd, 0xab, 0x9d, 0x16, 0x3e, 0x83, 0x88,
	0xc4, 0x0e, 0x1b, 0xa7, 0xdd, 0x02, 0x36, 0x87, 0x91, 0x14, 0x85, 0xa0, 0x35, 0x96, 0x12, 0xc9,
	0xed, 0xe8, 0x52, 0xe6, 0x3b, 0x4e, 0x15, 0x3a, 0x0b, 0x16, 0x18, 0x03, 0x78, 0xa0, 0x93, 0x01,
	0x3c, 0x90, 0xd6, 0xc2, 0xd5, 0x92, 0x96, 0x55, 0x65, 0x3c, 0x0c, 0x32, 0x0f, 0xd3, 0xef, 0x30,
	0x3e, 0x4b, 0x79, 0xa9, 0x93, 0x2f, 0xf7, 0x01, 0xbc, 0x59, 0x99, 0xb0, 0xac, 0x51, 0xee, 0x79,
	0x89, 0xec, 0x33, 0xf4, 0x6d, 0x0c, 0xd8, 0x74, 0xe1, 0x42, 0x75, 0xca, 0x50, 0xc2, 0x2e, 0x29,
	0xd5, 0xb2, 0x4f, 0x10, 0x99, 0x77, 0x64, 0x13, 0xdf, 0xf4, 0x3f, 0x34, 0x99, 0x5e, 0x30, 0xaa,
	0x65, 0x3f, 0x60, 0xe2, 0xe5, 0xfe, 0xc9, 0xe5, 0x6e, 0x23, 0x6e, 0x1b, 0xf6, 0xce, 0x8f, 0x75,
	0xde, 0x34, 0x99, 0x3d, 0x25, 0x55, 0x5b, 0xf4, 0x4d, 0xc0, 0xbf, 0x3e, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xad, 0x80, 0x14, 0x0a, 0xf0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoticeServiceClient is the client API for NoticeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoticeServiceClient interface {
	Health(ctx context.Context, in *HealthReq, opts ...grpc.CallOption) (*HealthResp, error)
	Email(ctx context.Context, in *EmailReq, opts ...grpc.CallOption) (*EmailResp, error)
	DingTalkMarkdown(ctx context.Context, in *DingTalkReq, opts ...grpc.CallOption) (*DingTalkResp, error)
}

type noticeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoticeServiceClient(cc *grpc.ClientConn) NoticeServiceClient {
	return &noticeServiceClient{cc}
}

func (c *noticeServiceClient) Health(ctx context.Context, in *HealthReq, opts ...grpc.CallOption) (*HealthResp, error) {
	out := new(HealthResp)
	err := c.cc.Invoke(ctx, "/notice.NoticeService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeServiceClient) Email(ctx context.Context, in *EmailReq, opts ...grpc.CallOption) (*EmailResp, error) {
	out := new(EmailResp)
	err := c.cc.Invoke(ctx, "/notice.NoticeService/Email", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeServiceClient) DingTalkMarkdown(ctx context.Context, in *DingTalkReq, opts ...grpc.CallOption) (*DingTalkResp, error) {
	out := new(DingTalkResp)
	err := c.cc.Invoke(ctx, "/notice.NoticeService/DingTalkMarkdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoticeServiceServer is the server API for NoticeService service.
type NoticeServiceServer interface {
	Health(context.Context, *HealthReq) (*HealthResp, error)
	Email(context.Context, *EmailReq) (*EmailResp, error)
	DingTalkMarkdown(context.Context, *DingTalkReq) (*DingTalkResp, error)
}

// UnimplementedNoticeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNoticeServiceServer struct {
}

func (*UnimplementedNoticeServiceServer) Health(ctx context.Context, req *HealthReq) (*HealthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedNoticeServiceServer) Email(ctx context.Context, req *EmailReq) (*EmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Email not implemented")
}
func (*UnimplementedNoticeServiceServer) DingTalkMarkdown(ctx context.Context, req *DingTalkReq) (*DingTalkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DingTalkMarkdown not implemented")
}

func RegisterNoticeServiceServer(s *grpc.Server, srv NoticeServiceServer) {
	s.RegisterService(&_NoticeService_serviceDesc, srv)
}

func _NoticeService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.NoticeService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).Health(ctx, req.(*HealthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoticeService_Email_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).Email(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.NoticeService/Email",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).Email(ctx, req.(*EmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoticeService_DingTalkMarkdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DingTalkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).DingTalkMarkdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.NoticeService/DingTalkMarkdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).DingTalkMarkdown(ctx, req.(*DingTalkReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoticeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notice.NoticeService",
	HandlerType: (*NoticeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _NoticeService_Health_Handler,
		},
		{
			MethodName: "Email",
			Handler:    _NoticeService_Email_Handler,
		},
		{
			MethodName: "DingTalkMarkdown",
			Handler:    _NoticeService_DingTalkMarkdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notice.proto",
}