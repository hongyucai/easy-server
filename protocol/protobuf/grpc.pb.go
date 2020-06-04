// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_test.proto

package grpc

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

// 定义一个请求类型
type RequestData struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestData) Reset()         { *m = RequestData{} }
func (m *RequestData) String() string { return proto.CompactTextString(m) }
func (*RequestData) ProtoMessage()    {}
func (*RequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdef1f95da3079a4, []int{0}
}

func (m *RequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestData.Unmarshal(m, b)
}
func (m *RequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestData.Marshal(b, m, deterministic)
}
func (m *RequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestData.Merge(m, src)
}
func (m *RequestData) XXX_Size() int {
	return xxx_messageInfo_RequestData.Size(m)
}
func (m *RequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestData.DiscardUnknown(m)
}

var xxx_messageInfo_RequestData proto.InternalMessageInfo

func (m *RequestData) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *RequestData) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

// 定义一个返回类型
type ResponseData struct {
	C                    int32    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseData) Reset()         { *m = ResponseData{} }
func (m *ResponseData) String() string { return proto.CompactTextString(m) }
func (*ResponseData) ProtoMessage()    {}
func (*ResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdef1f95da3079a4, []int{1}
}

func (m *ResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseData.Unmarshal(m, b)
}
func (m *ResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseData.Marshal(b, m, deterministic)
}
func (m *ResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseData.Merge(m, src)
}
func (m *ResponseData) XXX_Size() int {
	return xxx_messageInfo_ResponseData.Size(m)
}
func (m *ResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseData proto.InternalMessageInfo

func (m *ResponseData) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestData)(nil), "app.grpc.grpc_test.RequestData")
	proto.RegisterType((*ResponseData)(nil), "app.grpc.grpc_test.ResponseData")
}

func init() {
	proto.RegisterFile("grpc_test.proto", fileDescriptor_fdef1f95da3079a4)
}

var fileDescriptor_fdef1f95da3079a4 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2f, 0x2a, 0x48,
	0x8e, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0x2c, 0x28,
	0xd0, 0x03, 0x09, 0xea, 0xc1, 0x65, 0x94, 0x34, 0xb9, 0xb8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x78, 0xb8, 0x18, 0x13, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x83, 0x18, 0xc1, 0xbc, 0x24, 0x09, 0x26, 0x08, 0x2f, 0x49, 0x49, 0x86, 0x8b, 0x27, 0x28, 0xb5,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xa6, 0x36, 0x19, 0xa6, 0x36, 0xd9, 0x68, 0x09, 0x23, 0x17,
	0x97, 0x73, 0x62, 0x4e, 0x72, 0x69, 0x4e, 0x62, 0x49, 0x7e, 0x91, 0x90, 0x17, 0x17, 0xb3, 0x63,
	0x4a, 0x8a, 0x90, 0xbc, 0x1e, 0xa6, 0x9d, 0x7a, 0x48, 0x16, 0x4a, 0x29, 0x60, 0x57, 0x80, 0xb0,
	0x46, 0x89, 0x01, 0x64, 0x56, 0x70, 0x69, 0x12, 0x55, 0xcc, 0x4a, 0x62, 0x03, 0x07, 0x85, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x89, 0xeb, 0x35, 0x1d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error)
	Sub(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error) {
	out := new(ResponseData)
	err := c.cc.Invoke(ctx, "/app.grpc.grpc_test.Calculator/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Sub(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error) {
	out := new(ResponseData)
	err := c.cc.Invoke(ctx, "/app.grpc.grpc_test.Calculator/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Add(context.Context, *RequestData) (*ResponseData, error)
	Sub(context.Context, *RequestData) (*ResponseData, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Add(ctx context.Context, req *RequestData) (*ResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculatorServer) Sub(ctx context.Context, req *RequestData) (*ResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.grpc.grpc_test.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*RequestData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.grpc.grpc_test.Calculator/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sub(ctx, req.(*RequestData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.grpc.grpc_test.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Calculator_Sub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_test.proto",
}
