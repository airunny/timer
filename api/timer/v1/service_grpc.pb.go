// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: timer/v1/service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "timer/api/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Service_Healthy_FullMethodName           = "/api.timer.v1.Service/Healthy"
	Service_AddApplication_FullMethodName    = "/api.timer.v1.Service/AddApplication"
	Service_GetApplication_FullMethodName    = "/api.timer.v1.Service/GetApplication"
	Service_DeleteApplication_FullMethodName = "/api.timer.v1.Service/DeleteApplication"
	Service_UpdateApplication_FullMethodName = "/api.timer.v1.Service/UpdateApplication"
	Service_ListApplication_FullMethodName   = "/api.timer.v1.Service/ListApplication"
	Service_AddTimer_FullMethodName          = "/api.timer.v1.Service/AddTimer"
	Service_GetTimer_FullMethodName          = "/api.timer.v1.Service/GetTimer"
	Service_DeleteTimer_FullMethodName       = "/api.timer.v1.Service/DeleteTimer"
	Service_ReplayTimer_FullMethodName       = "/api.timer.v1.Service/ReplayTimer"
	Service_ListTimer_FullMethodName         = "/api.timer.v1.Service/ListTimer"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Healthy(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*common.HealthyReply, error)
	// 应用
	AddApplication(ctx context.Context, in *AddApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationReply, error)
	UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationReply, error)
	ListApplication(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationReply, error)
	// 定时器
	AddTimer(ctx context.Context, in *AddTimerRequest, opts ...grpc.CallOption) (*AddTimerReply, error)
	GetTimer(ctx context.Context, in *GetTimerRequest, opts ...grpc.CallOption) (*Timer, error)
	DeleteTimer(ctx context.Context, in *DeleteTimerRequest, opts ...grpc.CallOption) (*DeleteTimerReply, error)
	ReplayTimer(ctx context.Context, in *ReplayTimerRequest, opts ...grpc.CallOption) (*ReplayTimerReply, error)
	ListTimer(ctx context.Context, in *ListTimerRequest, opts ...grpc.CallOption) (*ListTimerReply, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Healthy(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*common.HealthyReply, error) {
	out := new(common.HealthyReply)
	err := c.cc.Invoke(ctx, Service_Healthy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddApplication(ctx context.Context, in *AddApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, Service_AddApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, Service_GetApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationReply, error) {
	out := new(DeleteApplicationReply)
	err := c.cc.Invoke(ctx, Service_DeleteApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationReply, error) {
	out := new(UpdateApplicationReply)
	err := c.cc.Invoke(ctx, Service_UpdateApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListApplication(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationReply, error) {
	out := new(ListApplicationReply)
	err := c.cc.Invoke(ctx, Service_ListApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddTimer(ctx context.Context, in *AddTimerRequest, opts ...grpc.CallOption) (*AddTimerReply, error) {
	out := new(AddTimerReply)
	err := c.cc.Invoke(ctx, Service_AddTimer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTimer(ctx context.Context, in *GetTimerRequest, opts ...grpc.CallOption) (*Timer, error) {
	out := new(Timer)
	err := c.cc.Invoke(ctx, Service_GetTimer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteTimer(ctx context.Context, in *DeleteTimerRequest, opts ...grpc.CallOption) (*DeleteTimerReply, error) {
	out := new(DeleteTimerReply)
	err := c.cc.Invoke(ctx, Service_DeleteTimer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ReplayTimer(ctx context.Context, in *ReplayTimerRequest, opts ...grpc.CallOption) (*ReplayTimerReply, error) {
	out := new(ReplayTimerReply)
	err := c.cc.Invoke(ctx, Service_ReplayTimer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListTimer(ctx context.Context, in *ListTimerRequest, opts ...grpc.CallOption) (*ListTimerReply, error) {
	out := new(ListTimerReply)
	err := c.cc.Invoke(ctx, Service_ListTimer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Healthy(context.Context, *common.EmptyRequest) (*common.HealthyReply, error)
	// 应用
	AddApplication(context.Context, *AddApplicationRequest) (*Application, error)
	GetApplication(context.Context, *GetApplicationRequest) (*Application, error)
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*DeleteApplicationReply, error)
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*UpdateApplicationReply, error)
	ListApplication(context.Context, *ListApplicationRequest) (*ListApplicationReply, error)
	// 定时器
	AddTimer(context.Context, *AddTimerRequest) (*AddTimerReply, error)
	GetTimer(context.Context, *GetTimerRequest) (*Timer, error)
	DeleteTimer(context.Context, *DeleteTimerRequest) (*DeleteTimerReply, error)
	ReplayTimer(context.Context, *ReplayTimerRequest) (*ReplayTimerReply, error)
	ListTimer(context.Context, *ListTimerRequest) (*ListTimerReply, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Healthy(context.Context, *common.EmptyRequest) (*common.HealthyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthy not implemented")
}
func (UnimplementedServiceServer) AddApplication(context.Context, *AddApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddApplication not implemented")
}
func (UnimplementedServiceServer) GetApplication(context.Context, *GetApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplication not implemented")
}
func (UnimplementedServiceServer) DeleteApplication(context.Context, *DeleteApplicationRequest) (*DeleteApplicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplication not implemented")
}
func (UnimplementedServiceServer) UpdateApplication(context.Context, *UpdateApplicationRequest) (*UpdateApplicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplication not implemented")
}
func (UnimplementedServiceServer) ListApplication(context.Context, *ListApplicationRequest) (*ListApplicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplication not implemented")
}
func (UnimplementedServiceServer) AddTimer(context.Context, *AddTimerRequest) (*AddTimerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTimer not implemented")
}
func (UnimplementedServiceServer) GetTimer(context.Context, *GetTimerRequest) (*Timer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimer not implemented")
}
func (UnimplementedServiceServer) DeleteTimer(context.Context, *DeleteTimerRequest) (*DeleteTimerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimer not implemented")
}
func (UnimplementedServiceServer) ReplayTimer(context.Context, *ReplayTimerRequest) (*ReplayTimerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplayTimer not implemented")
}
func (UnimplementedServiceServer) ListTimer(context.Context, *ListTimerRequest) (*ListTimerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTimer not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Healthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Healthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Healthy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Healthy(ctx, req.(*common.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddApplication(ctx, req.(*AddApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetApplication(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteApplication(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateApplication(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListApplication(ctx, req.(*ListApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddTimer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddTimer(ctx, req.(*AddTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetTimer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTimer(ctx, req.(*GetTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteTimer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteTimer(ctx, req.(*DeleteTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ReplayTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplayTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ReplayTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ReplayTimer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ReplayTimer(ctx, req.(*ReplayTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListTimer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListTimer(ctx, req.(*ListTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.timer.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthy",
			Handler:    _Service_Healthy_Handler,
		},
		{
			MethodName: "AddApplication",
			Handler:    _Service_AddApplication_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _Service_GetApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _Service_DeleteApplication_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _Service_UpdateApplication_Handler,
		},
		{
			MethodName: "ListApplication",
			Handler:    _Service_ListApplication_Handler,
		},
		{
			MethodName: "AddTimer",
			Handler:    _Service_AddTimer_Handler,
		},
		{
			MethodName: "GetTimer",
			Handler:    _Service_GetTimer_Handler,
		},
		{
			MethodName: "DeleteTimer",
			Handler:    _Service_DeleteTimer_Handler,
		},
		{
			MethodName: "ReplayTimer",
			Handler:    _Service_ReplayTimer_Handler,
		},
		{
			MethodName: "ListTimer",
			Handler:    _Service_ListTimer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timer/v1/service.proto",
}