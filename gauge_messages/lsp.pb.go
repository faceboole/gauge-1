// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lsp.proto

package gauge_messages

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

// Empty is a blank response, to be used when there is no return expected.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Empty)(nil), "gauge.messages.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LspService service

type LspServiceClient interface {
	GetStepNames(ctx context.Context, in *StepNamesRequest, opts ...grpc.CallOption) (*StepNamesResponse, error)
	CacheFile(ctx context.Context, in *CacheFileRequest, opts ...grpc.CallOption) (*Empty, error)
	GetStepPositions(ctx context.Context, in *StepPositionsRequest, opts ...grpc.CallOption) (*StepPositionsResponse, error)
	GetImplementationFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ImplementationFileListResponse, error)
	ImplementStub(ctx context.Context, in *StubImplementationCodeRequest, opts ...grpc.CallOption) (*FileDiff, error)
}

type lspServiceClient struct {
	cc *grpc.ClientConn
}

func NewLspServiceClient(cc *grpc.ClientConn) LspServiceClient {
	return &lspServiceClient{cc}
}

func (c *lspServiceClient) GetStepNames(ctx context.Context, in *StepNamesRequest, opts ...grpc.CallOption) (*StepNamesResponse, error) {
	out := new(StepNamesResponse)
	err := grpc.Invoke(ctx, "/gauge.messages.lspService/GetStepNames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lspServiceClient) CacheFile(ctx context.Context, in *CacheFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/gauge.messages.lspService/CacheFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lspServiceClient) GetStepPositions(ctx context.Context, in *StepPositionsRequest, opts ...grpc.CallOption) (*StepPositionsResponse, error) {
	out := new(StepPositionsResponse)
	err := grpc.Invoke(ctx, "/gauge.messages.lspService/GetStepPositions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lspServiceClient) GetImplementationFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ImplementationFileListResponse, error) {
	out := new(ImplementationFileListResponse)
	err := grpc.Invoke(ctx, "/gauge.messages.lspService/GetImplementationFiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lspServiceClient) ImplementStub(ctx context.Context, in *StubImplementationCodeRequest, opts ...grpc.CallOption) (*FileDiff, error) {
	out := new(FileDiff)
	err := grpc.Invoke(ctx, "/gauge.messages.lspService/ImplementStub", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LspService service

type LspServiceServer interface {
	GetStepNames(context.Context, *StepNamesRequest) (*StepNamesResponse, error)
	CacheFile(context.Context, *CacheFileRequest) (*Empty, error)
	GetStepPositions(context.Context, *StepPositionsRequest) (*StepPositionsResponse, error)
	GetImplementationFiles(context.Context, *Empty) (*ImplementationFileListResponse, error)
	ImplementStub(context.Context, *StubImplementationCodeRequest) (*FileDiff, error)
}

func RegisterLspServiceServer(s *grpc.Server, srv LspServiceServer) {
	s.RegisterService(&_LspService_serviceDesc, srv)
}

func _LspService_GetStepNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetStepNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gauge.messages.lspService/GetStepNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetStepNames(ctx, req.(*StepNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_CacheFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).CacheFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gauge.messages.lspService/CacheFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).CacheFile(ctx, req.(*CacheFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_GetStepPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetStepPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gauge.messages.lspService/GetStepPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetStepPositions(ctx, req.(*StepPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_GetImplementationFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetImplementationFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gauge.messages.lspService/GetImplementationFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetImplementationFiles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_ImplementStub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubImplementationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).ImplementStub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gauge.messages.lspService/ImplementStub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).ImplementStub(ctx, req.(*StubImplementationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LspService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gauge.messages.lspService",
	HandlerType: (*LspServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStepNames",
			Handler:    _LspService_GetStepNames_Handler,
		},
		{
			MethodName: "CacheFile",
			Handler:    _LspService_CacheFile_Handler,
		},
		{
			MethodName: "GetStepPositions",
			Handler:    _LspService_GetStepPositions_Handler,
		},
		{
			MethodName: "GetImplementationFiles",
			Handler:    _LspService_GetImplementationFiles_Handler,
		},
		{
			MethodName: "ImplementStub",
			Handler:    _LspService_ImplementStub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lsp.proto",
}

func init() { proto.RegisterFile("lsp.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x41, 0x51, 0xe9, 0xa0, 0x41, 0x07, 0x14, 0xc9, 0x49, 0x45, 0x8f, 0xee, 0x41, 0xdf,
	0xc0, 0x56, 0x83, 0xa0, 0x22, 0xe6, 0xe2, 0xad, 0x6c, 0xea, 0x34, 0x2e, 0x64, 0xb3, 0x6b, 0x67,
	0x22, 0xf8, 0x4a, 0xbe, 0x86, 0x2f, 0x26, 0x89, 0xc9, 0x42, 0x93, 0x62, 0x6f, 0x21, 0xdf, 0xbf,
	0xff, 0xff, 0xb1, 0x0b, 0xa3, 0x82, 0xbd, 0xf2, 0x0b, 0x27, 0x0e, 0xa3, 0x5c, 0x57, 0x39, 0x29,
	0x4b, 0xcc, 0x3a, 0x27, 0x8e, 0xa3, 0xee, 0xeb, 0x8f, 0x9f, 0xed, 0xc0, 0xd6, 0xad, 0xf5, 0xf2,
	0x75, 0xf5, 0xb3, 0x09, 0x50, 0xb0, 0x4f, 0x69, 0xf1, 0x69, 0x66, 0x84, 0x29, 0xec, 0x26, 0x24,
	0xa9, 0x90, 0x7f, 0xd2, 0x96, 0x18, 0x4f, 0xd4, 0x72, 0x91, 0x0a, 0xe8, 0x85, 0x3e, 0x2a, 0x62,
	0x89, 0x4f, 0xff, 0x49, 0xb0, 0x77, 0x25, 0x13, 0x4e, 0x60, 0x34, 0xd6, 0xb3, 0x77, 0xba, 0x33,
	0x05, 0x0d, 0x1b, 0x03, 0xea, 0x1a, 0x0f, 0xfb, 0x89, 0xc6, 0x14, 0xa7, 0xb0, 0xdf, 0xaa, 0x3d,
	0x3b, 0x36, 0x62, 0x5c, 0xc9, 0x78, 0xbe, 0x6a, 0x3c, 0xe0, 0xae, 0xf0, 0x62, 0x4d, 0xaa, 0xd5,
	0x9c, 0xc2, 0x51, 0x42, 0x72, 0x6f, 0x7d, 0x41, 0x96, 0x4a, 0xd1, 0x35, 0xad, 0xbd, 0x18, 0x57,
	0x1b, 0xc5, 0xaa, 0xff, 0x7b, 0x78, 0xf6, 0xc1, 0xb0, 0x84, 0x81, 0x57, 0xd8, 0x0b, 0x89, 0x54,
	0xaa, 0x0c, 0x2f, 0x87, 0x62, 0x55, 0xb6, 0x5c, 0x32, 0x76, 0x6f, 0xe1, 0x62, 0x8e, 0xfb, 0xf1,
	0x7a, 0x61, 0x62, 0xe6, 0xf3, 0x9b, 0x83, 0xef, 0x8d, 0x28, 0x69, 0xd8, 0x63, 0xcb, 0xb2, 0xed,
	0xe6, 0xa1, 0xaf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xb4, 0x9f, 0xbf, 0x15, 0x02, 0x00,
	0x00,
}
