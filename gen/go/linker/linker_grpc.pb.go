// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.0
// source: linker/linker.proto

package linker_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LinkerClient is the client API for Linker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkerClient interface {
	Post(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Pick(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type linkerClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkerClient(cc grpc.ClientConnInterface) LinkerClient {
	return &linkerClient{cc}
}

func (c *linkerClient) Post(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/linker.Linker/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkerClient) Pick(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/linker.Linker/Pick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkerClient) List(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/linker.Linker/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkerServer is the server API for Linker service.
// All implementations must embed UnimplementedLinkerServer
// for forward compatibility
type LinkerServer interface {
	Post(context.Context, *Request) (*Response, error)
	Pick(context.Context, *Request) (*Response, error)
	List(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedLinkerServer()
}

// UnimplementedLinkerServer must be embedded to have forward compatible implementations.
type UnimplementedLinkerServer struct {
}

func (UnimplementedLinkerServer) Post(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedLinkerServer) Pick(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pick not implemented")
}
func (UnimplementedLinkerServer) List(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLinkerServer) mustEmbedUnimplementedLinkerServer() {}

// UnsafeLinkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkerServer will
// result in compilation errors.
type UnsafeLinkerServer interface {
	mustEmbedUnimplementedLinkerServer()
}

func RegisterLinkerServer(s grpc.ServiceRegistrar, srv LinkerServer) {
	s.RegisterService(&Linker_ServiceDesc, srv)
}

func _Linker_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkerServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linker.Linker/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkerServer).Post(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Linker_Pick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkerServer).Pick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linker.Linker/Pick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkerServer).Pick(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Linker_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linker.Linker/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkerServer).List(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Linker_ServiceDesc is the grpc.ServiceDesc for Linker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Linker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linker.Linker",
	HandlerType: (*LinkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Linker_Post_Handler,
		},
		{
			MethodName: "Pick",
			Handler:    _Linker_Pick_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Linker_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "linker/linker.proto",
}
