// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: notification_service.proto

package notification

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
	NotificationService_Get_FullMethodName                             = "/notification.NotificationService/Get"
	NotificationService_GetAll_FullMethodName                          = "/notification.NotificationService/GetAll"
	NotificationService_GetByReceiver_FullMethodName                   = "/notification.NotificationService/GetByReceiver"
	NotificationService_GetBySender_FullMethodName                     = "/notification.NotificationService/GetBySender"
	NotificationService_CreateNotification_FullMethodName              = "/notification.NotificationService/CreateNotification"
	NotificationService_EditNotification_FullMethodName                = "/notification.NotificationService/EditNotification"
	NotificationService_DeleteNotification_FullMethodName              = "/notification.NotificationService/DeleteNotification"
	NotificationService_ReadAllByUsername_FullMethodName               = "/notification.NotificationService/readAllByUsername"
	NotificationService_GetNotificationFilterByUsername_FullMethodName = "/notification.NotificationService/GetNotificationFilterByUsername"
	NotificationService_EditNotificationFilter_FullMethodName          = "/notification.NotificationService/EditNotificationFilter"
	NotificationService_CreateNotificationFilter_FullMethodName        = "/notification.NotificationService/CreateNotificationFilter"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetByReceiver(ctx context.Context, in *GetRequestReceiver, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetBySender(ctx context.Context, in *GetRequestSender, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
	EditNotification(ctx context.Context, in *EditNotificationRequest, opts ...grpc.CallOption) (*EditNotificationResponse, error)
	DeleteNotification(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ReadAllByUsername(ctx context.Context, in *ReadAllByUsernameRequest, opts ...grpc.CallOption) (*ReadAllByUsernameResponse, error)
	GetNotificationFilterByUsername(ctx context.Context, in *ReadAllByUsernameRequest, opts ...grpc.CallOption) (*NotificationFilterResponse, error)
	EditNotificationFilter(ctx context.Context, in *NotificationFilterRequest, opts ...grpc.CallOption) (*NotificationFilterResponse, error)
	CreateNotificationFilter(ctx context.Context, in *NotificationFilterRequest, opts ...grpc.CallOption) (*NotificationFilterResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, NotificationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetByReceiver(ctx context.Context, in *GetRequestReceiver, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetByReceiver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetBySender(ctx context.Context, in *GetRequestSender, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetBySender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	out := new(CreateNotificationResponse)
	err := c.cc.Invoke(ctx, NotificationService_CreateNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) EditNotification(ctx context.Context, in *EditNotificationRequest, opts ...grpc.CallOption) (*EditNotificationResponse, error) {
	out := new(EditNotificationResponse)
	err := c.cc.Invoke(ctx, NotificationService_EditNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) DeleteNotification(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, NotificationService_DeleteNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ReadAllByUsername(ctx context.Context, in *ReadAllByUsernameRequest, opts ...grpc.CallOption) (*ReadAllByUsernameResponse, error) {
	out := new(ReadAllByUsernameResponse)
	err := c.cc.Invoke(ctx, NotificationService_ReadAllByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetNotificationFilterByUsername(ctx context.Context, in *ReadAllByUsernameRequest, opts ...grpc.CallOption) (*NotificationFilterResponse, error) {
	out := new(NotificationFilterResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetNotificationFilterByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) EditNotificationFilter(ctx context.Context, in *NotificationFilterRequest, opts ...grpc.CallOption) (*NotificationFilterResponse, error) {
	out := new(NotificationFilterResponse)
	err := c.cc.Invoke(ctx, NotificationService_EditNotificationFilter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) CreateNotificationFilter(ctx context.Context, in *NotificationFilterRequest, opts ...grpc.CallOption) (*NotificationFilterResponse, error) {
	out := new(NotificationFilterResponse)
	err := c.cc.Invoke(ctx, NotificationService_CreateNotificationFilter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetByReceiver(context.Context, *GetRequestReceiver) (*GetAllResponse, error)
	GetBySender(context.Context, *GetRequestSender) (*GetAllResponse, error)
	CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error)
	EditNotification(context.Context, *EditNotificationRequest) (*EditNotificationResponse, error)
	DeleteNotification(context.Context, *DeleteRequest) (*DeleteResponse, error)
	ReadAllByUsername(context.Context, *ReadAllByUsernameRequest) (*ReadAllByUsernameResponse, error)
	GetNotificationFilterByUsername(context.Context, *ReadAllByUsernameRequest) (*NotificationFilterResponse, error)
	EditNotificationFilter(context.Context, *NotificationFilterRequest) (*NotificationFilterResponse, error)
	CreateNotificationFilter(context.Context, *NotificationFilterRequest) (*NotificationFilterResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNotificationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedNotificationServiceServer) GetByReceiver(context.Context, *GetRequestReceiver) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByReceiver not implemented")
}
func (UnimplementedNotificationServiceServer) GetBySender(context.Context, *GetRequestSender) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySender not implemented")
}
func (UnimplementedNotificationServiceServer) CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationServiceServer) EditNotification(context.Context, *EditNotificationRequest) (*EditNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditNotification not implemented")
}
func (UnimplementedNotificationServiceServer) DeleteNotification(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}
func (UnimplementedNotificationServiceServer) ReadAllByUsername(context.Context, *ReadAllByUsernameRequest) (*ReadAllByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllByUsername not implemented")
}
func (UnimplementedNotificationServiceServer) GetNotificationFilterByUsername(context.Context, *ReadAllByUsernameRequest) (*NotificationFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationFilterByUsername not implemented")
}
func (UnimplementedNotificationServiceServer) EditNotificationFilter(context.Context, *NotificationFilterRequest) (*NotificationFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditNotificationFilter not implemented")
}
func (UnimplementedNotificationServiceServer) CreateNotificationFilter(context.Context, *NotificationFilterRequest) (*NotificationFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotificationFilter not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetByReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestReceiver)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetByReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetByReceiver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetByReceiver(ctx, req.(*GetRequestReceiver))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetBySender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestSender)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetBySender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetBySender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetBySender(ctx, req.(*GetRequestSender))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_EditNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).EditNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_EditNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).EditNotification(ctx, req.(*EditNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_DeleteNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ReadAllByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ReadAllByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_ReadAllByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ReadAllByUsername(ctx, req.(*ReadAllByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetNotificationFilterByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotificationFilterByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetNotificationFilterByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotificationFilterByUsername(ctx, req.(*ReadAllByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_EditNotificationFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).EditNotificationFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_EditNotificationFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).EditNotificationFilter(ctx, req.(*NotificationFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_CreateNotificationFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).CreateNotificationFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_CreateNotificationFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).CreateNotificationFilter(ctx, req.(*NotificationFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NotificationService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _NotificationService_GetAll_Handler,
		},
		{
			MethodName: "GetByReceiver",
			Handler:    _NotificationService_GetByReceiver_Handler,
		},
		{
			MethodName: "GetBySender",
			Handler:    _NotificationService_GetBySender_Handler,
		},
		{
			MethodName: "CreateNotification",
			Handler:    _NotificationService_CreateNotification_Handler,
		},
		{
			MethodName: "EditNotification",
			Handler:    _NotificationService_EditNotification_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _NotificationService_DeleteNotification_Handler,
		},
		{
			MethodName: "readAllByUsername",
			Handler:    _NotificationService_ReadAllByUsername_Handler,
		},
		{
			MethodName: "GetNotificationFilterByUsername",
			Handler:    _NotificationService_GetNotificationFilterByUsername_Handler,
		},
		{
			MethodName: "EditNotificationFilter",
			Handler:    _NotificationService_EditNotificationFilter_Handler,
		},
		{
			MethodName: "CreateNotificationFilter",
			Handler:    _NotificationService_CreateNotificationFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification_service.proto",
}
