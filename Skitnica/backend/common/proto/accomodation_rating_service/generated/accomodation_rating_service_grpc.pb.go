// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0--rc2
// source: accomodation_rating_service.proto

package accomodation_rating

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
	AccomodationRatingService_Get_FullMethodName                      = "/accomodation_rating.AccomodationRatingService/Get"
	AccomodationRatingService_GetAll_FullMethodName                   = "/accomodation_rating.AccomodationRatingService/GetAll"
	AccomodationRatingService_GetAllByAccomodationId_FullMethodName   = "/accomodation_rating.AccomodationRatingService/GetAllByAccomodationId"
	AccomodationRatingService_CreateAccomodationRating_FullMethodName = "/accomodation_rating.AccomodationRatingService/CreateAccomodationRating"
	AccomodationRatingService_EditAccomodationRating_FullMethodName   = "/accomodation_rating.AccomodationRatingService/EditAccomodationRating"
	AccomodationRatingService_DeleteAccomodationRating_FullMethodName = "/accomodation_rating.AccomodationRatingService/DeleteAccomodationRating"
	AccomodationRatingService_GetByAccomodationAndUser_FullMethodName = "/accomodation_rating.AccomodationRatingService/GetByAccomodationAndUser"
	AccomodationRatingService_GetAllRecommended_FullMethodName        = "/accomodation_rating.AccomodationRatingService/GetAllRecommended"
)

// AccomodationRatingServiceClient is the client API for AccomodationRatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccomodationRatingServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetAllByAccomodationId(ctx context.Context, in *GetAllByAccomodationIdRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreateAccomodationRating(ctx context.Context, in *CreateAccomodationRatingRequest, opts ...grpc.CallOption) (*CreateAccomodationRatingResponse, error)
	EditAccomodationRating(ctx context.Context, in *EditAccomodationRatingRequest, opts ...grpc.CallOption) (*EditAccomodationRatingResponse, error)
	DeleteAccomodationRating(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetByAccomodationAndUser(ctx context.Context, in *GetByAccomodationAndUserRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetAllRecommended(ctx context.Context, in *GetAllRecommendedRequest, opts ...grpc.CallOption) (*GetAllRecommendedResponse, error)
}

type accomodationRatingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccomodationRatingServiceClient(cc grpc.ClientConnInterface) AccomodationRatingServiceClient {
	return &accomodationRatingServiceClient{cc}
}

func (c *accomodationRatingServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationRatingServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationRatingServiceClient) GetAllByAccomodationId(ctx context.Context, in *GetAllByAccomodationIdRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_GetAllByAccomodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationRatingServiceClient) CreateAccomodationRating(ctx context.Context, in *CreateAccomodationRatingRequest, opts ...grpc.CallOption) (*CreateAccomodationRatingResponse, error) {
	out := new(CreateAccomodationRatingResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_CreateAccomodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationRatingServiceClient) EditAccomodationRating(ctx context.Context, in *EditAccomodationRatingRequest, opts ...grpc.CallOption) (*EditAccomodationRatingResponse, error) {
	out := new(EditAccomodationRatingResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_EditAccomodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationRatingServiceClient) DeleteAccomodationRating(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_DeleteAccomodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationRatingServiceClient) GetByAccomodationAndUser(ctx context.Context, in *GetByAccomodationAndUserRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_GetByAccomodationAndUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accomodationRatingServiceClient) GetAllRecommended(ctx context.Context, in *GetAllRecommendedRequest, opts ...grpc.CallOption) (*GetAllRecommendedResponse, error) {
	out := new(GetAllRecommendedResponse)
	err := c.cc.Invoke(ctx, AccomodationRatingService_GetAllRecommended_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccomodationRatingServiceServer is the server API for AccomodationRatingService service.
// All implementations must embed UnimplementedAccomodationRatingServiceServer
// for forward compatibility
type AccomodationRatingServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetAllByAccomodationId(context.Context, *GetAllByAccomodationIdRequest) (*GetAllResponse, error)
	CreateAccomodationRating(context.Context, *CreateAccomodationRatingRequest) (*CreateAccomodationRatingResponse, error)
	EditAccomodationRating(context.Context, *EditAccomodationRatingRequest) (*EditAccomodationRatingResponse, error)
	DeleteAccomodationRating(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetByAccomodationAndUser(context.Context, *GetByAccomodationAndUserRequest) (*GetAllResponse, error)
	GetAllRecommended(context.Context, *GetAllRecommendedRequest) (*GetAllRecommendedResponse, error)
	mustEmbedUnimplementedAccomodationRatingServiceServer()
}

// UnimplementedAccomodationRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccomodationRatingServiceServer struct {
}

func (UnimplementedAccomodationRatingServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) GetAllByAccomodationId(context.Context, *GetAllByAccomodationIdRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByAccomodationId not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) CreateAccomodationRating(context.Context, *CreateAccomodationRatingRequest) (*CreateAccomodationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccomodationRating not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) EditAccomodationRating(context.Context, *EditAccomodationRatingRequest) (*EditAccomodationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAccomodationRating not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) DeleteAccomodationRating(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccomodationRating not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) GetByAccomodationAndUser(context.Context, *GetByAccomodationAndUserRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccomodationAndUser not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) GetAllRecommended(context.Context, *GetAllRecommendedRequest) (*GetAllRecommendedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRecommended not implemented")
}
func (UnimplementedAccomodationRatingServiceServer) mustEmbedUnimplementedAccomodationRatingServiceServer() {
}

// UnsafeAccomodationRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccomodationRatingServiceServer will
// result in compilation errors.
type UnsafeAccomodationRatingServiceServer interface {
	mustEmbedUnimplementedAccomodationRatingServiceServer()
}

func RegisterAccomodationRatingServiceServer(s grpc.ServiceRegistrar, srv AccomodationRatingServiceServer) {
	s.RegisterService(&AccomodationRatingService_ServiceDesc, srv)
}

func _AccomodationRatingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationRatingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationRatingService_GetAllByAccomodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByAccomodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).GetAllByAccomodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_GetAllByAccomodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).GetAllByAccomodationId(ctx, req.(*GetAllByAccomodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationRatingService_CreateAccomodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccomodationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).CreateAccomodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_CreateAccomodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).CreateAccomodationRating(ctx, req.(*CreateAccomodationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationRatingService_EditAccomodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAccomodationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).EditAccomodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_EditAccomodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).EditAccomodationRating(ctx, req.(*EditAccomodationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationRatingService_DeleteAccomodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).DeleteAccomodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_DeleteAccomodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).DeleteAccomodationRating(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationRatingService_GetByAccomodationAndUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccomodationAndUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).GetByAccomodationAndUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_GetByAccomodationAndUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).GetByAccomodationAndUser(ctx, req.(*GetByAccomodationAndUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccomodationRatingService_GetAllRecommended_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRecommendedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationRatingServiceServer).GetAllRecommended(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationRatingService_GetAllRecommended_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationRatingServiceServer).GetAllRecommended(ctx, req.(*GetAllRecommendedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccomodationRatingService_ServiceDesc is the grpc.ServiceDesc for AccomodationRatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccomodationRatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accomodation_rating.AccomodationRatingService",
	HandlerType: (*AccomodationRatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccomodationRatingService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AccomodationRatingService_GetAll_Handler,
		},
		{
			MethodName: "GetAllByAccomodationId",
			Handler:    _AccomodationRatingService_GetAllByAccomodationId_Handler,
		},
		{
			MethodName: "CreateAccomodationRating",
			Handler:    _AccomodationRatingService_CreateAccomodationRating_Handler,
		},
		{
			MethodName: "EditAccomodationRating",
			Handler:    _AccomodationRatingService_EditAccomodationRating_Handler,
		},
		{
			MethodName: "DeleteAccomodationRating",
			Handler:    _AccomodationRatingService_DeleteAccomodationRating_Handler,
		},
		{
			MethodName: "GetByAccomodationAndUser",
			Handler:    _AccomodationRatingService_GetByAccomodationAndUser_Handler,
		},
		{
			MethodName: "GetAllRecommended",
			Handler:    _AccomodationRatingService_GetAllRecommended_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accomodation_rating_service.proto",
}
