// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0--rc2
// source: hostMark_service.proto

package hostMark

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

const (
	HostMarkService_Get_FullMethodName            = "/hostMark.HostMarkService/Get"
	HostMarkService_GetAll_FullMethodName         = "/hostMark.HostMarkService/GetAll"
	HostMarkService_CreateHostMark_FullMethodName = "/hostMark.HostMarkService/CreateHostMark"
	HostMarkService_EditHostMark_FullMethodName   = "/hostMark.HostMarkService/EditHostMark"
	HostMarkService_DeleteHostMark_FullMethodName = "/hostMark.HostMarkService/DeleteHostMark"
)

// HostMarkServiceClient is the client API for HostMarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostMarkServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreateHostMark(ctx context.Context, in *CreateHostMarkRequest, opts ...grpc.CallOption) (*CreateHostMarkResponse, error)
	EditHostMark(ctx context.Context, in *EditHostMarkRequest, opts ...grpc.CallOption) (*EditHostMarkResponse, error)
	DeleteHostMark(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type hostMarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostMarkServiceClient(cc grpc.ClientConnInterface) HostMarkServiceClient {
	return &hostMarkServiceClient{cc}
}

func (c *hostMarkServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, HostMarkService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostMarkServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, HostMarkService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostMarkServiceClient) CreateHostMark(ctx context.Context, in *CreateHostMarkRequest, opts ...grpc.CallOption) (*CreateHostMarkResponse, error) {
	out := new(CreateHostMarkResponse)
	err := c.cc.Invoke(ctx, HostMarkService_CreateHostMark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostMarkServiceClient) EditHostMark(ctx context.Context, in *EditHostMarkRequest, opts ...grpc.CallOption) (*EditHostMarkResponse, error) {
	out := new(EditHostMarkResponse)
	err := c.cc.Invoke(ctx, HostMarkService_EditHostMark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostMarkServiceClient) DeleteHostMark(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, HostMarkService_DeleteHostMark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostMarkServiceServer is the server API for HostMarkService service.
// All implementations must embed UnimplementedHostMarkServiceServer
// for forward compatibility
type HostMarkServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	CreateHostMark(context.Context, *CreateHostMarkRequest) (*CreateHostMarkResponse, error)
	EditHostMark(context.Context, *EditHostMarkRequest) (*EditHostMarkResponse, error)
	DeleteHostMark(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedHostMarkServiceServer()
}

// UnimplementedHostMarkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHostMarkServiceServer struct {
}

func (UnimplementedHostMarkServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHostMarkServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedHostMarkServiceServer) CreateHostMark(context.Context, *CreateHostMarkRequest) (*CreateHostMarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHostMark not implemented")
}
func (UnimplementedHostMarkServiceServer) EditHostMark(context.Context, *EditHostMarkRequest) (*EditHostMarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditHostMark not implemented")
}
func (UnimplementedHostMarkServiceServer) DeleteHostMark(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHostMark not implemented")
}
func (UnimplementedHostMarkServiceServer) mustEmbedUnimplementedHostMarkServiceServer() {}

// UnsafeHostMarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostMarkServiceServer will
// result in compilation errors.
type UnsafeHostMarkServiceServer interface {
	mustEmbedUnimplementedHostMarkServiceServer()
}

func RegisterHostMarkServiceServer(s grpc.ServiceRegistrar, srv HostMarkServiceServer) {
	s.RegisterService(&HostMarkService_ServiceDesc, srv)
}

func _HostMarkService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostMarkServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostMarkService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostMarkServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostMarkService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostMarkServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostMarkService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostMarkServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostMarkService_CreateHostMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHostMarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostMarkServiceServer).CreateHostMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostMarkService_CreateHostMark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostMarkServiceServer).CreateHostMark(ctx, req.(*CreateHostMarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostMarkService_EditHostMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditHostMarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostMarkServiceServer).EditHostMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostMarkService_EditHostMark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostMarkServiceServer).EditHostMark(ctx, req.(*EditHostMarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostMarkService_DeleteHostMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostMarkServiceServer).DeleteHostMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostMarkService_DeleteHostMark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostMarkServiceServer).DeleteHostMark(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostMarkService_ServiceDesc is the grpc.ServiceDesc for HostMarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostMarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hostMark.HostMarkService",
	HandlerType: (*HostMarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _HostMarkService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _HostMarkService_GetAll_Handler,
		},
		{
			MethodName: "CreateHostMark",
			Handler:    _HostMarkService_CreateHostMark_Handler,
		},
		{
			MethodName: "EditHostMark",
			Handler:    _HostMarkService_EditHostMark_Handler,
		},
		{
			MethodName: "DeleteHostMark",
			Handler:    _HostMarkService_DeleteHostMark_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostMark_service.proto",
}
