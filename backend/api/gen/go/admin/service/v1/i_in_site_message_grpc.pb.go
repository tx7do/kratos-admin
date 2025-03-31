// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_in_site_message.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-admin/api/gen/go/system/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InSiteMessageService_ListInSiteMessage_FullMethodName   = "/admin.service.v1.InSiteMessageService/ListInSiteMessage"
	InSiteMessageService_GetInSiteMessage_FullMethodName    = "/admin.service.v1.InSiteMessageService/GetInSiteMessage"
	InSiteMessageService_CreateInSiteMessage_FullMethodName = "/admin.service.v1.InSiteMessageService/CreateInSiteMessage"
	InSiteMessageService_UpdateInSiteMessage_FullMethodName = "/admin.service.v1.InSiteMessageService/UpdateInSiteMessage"
	InSiteMessageService_DeleteInSiteMessage_FullMethodName = "/admin.service.v1.InSiteMessageService/DeleteInSiteMessage"
)

// InSiteMessageServiceClient is the client API for InSiteMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 站内信管理服务
type InSiteMessageServiceClient interface {
	// 查询站内信列表
	ListInSiteMessage(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListInSiteMessageResponse, error)
	// 查询站内信详情
	GetInSiteMessage(ctx context.Context, in *v11.GetInSiteMessageRequest, opts ...grpc.CallOption) (*v11.InSiteMessage, error)
	// 创建站内信
	CreateInSiteMessage(ctx context.Context, in *v11.CreateInSiteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新站内信
	UpdateInSiteMessage(ctx context.Context, in *v11.UpdateInSiteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除站内信
	DeleteInSiteMessage(ctx context.Context, in *v11.DeleteInSiteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type inSiteMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInSiteMessageServiceClient(cc grpc.ClientConnInterface) InSiteMessageServiceClient {
	return &inSiteMessageServiceClient{cc}
}

func (c *inSiteMessageServiceClient) ListInSiteMessage(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListInSiteMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.ListInSiteMessageResponse)
	err := c.cc.Invoke(ctx, InSiteMessageService_ListInSiteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inSiteMessageServiceClient) GetInSiteMessage(ctx context.Context, in *v11.GetInSiteMessageRequest, opts ...grpc.CallOption) (*v11.InSiteMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.InSiteMessage)
	err := c.cc.Invoke(ctx, InSiteMessageService_GetInSiteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inSiteMessageServiceClient) CreateInSiteMessage(ctx context.Context, in *v11.CreateInSiteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, InSiteMessageService_CreateInSiteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inSiteMessageServiceClient) UpdateInSiteMessage(ctx context.Context, in *v11.UpdateInSiteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, InSiteMessageService_UpdateInSiteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inSiteMessageServiceClient) DeleteInSiteMessage(ctx context.Context, in *v11.DeleteInSiteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, InSiteMessageService_DeleteInSiteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InSiteMessageServiceServer is the server API for InSiteMessageService service.
// All implementations must embed UnimplementedInSiteMessageServiceServer
// for forward compatibility.
//
// 站内信管理服务
type InSiteMessageServiceServer interface {
	// 查询站内信列表
	ListInSiteMessage(context.Context, *v1.PagingRequest) (*v11.ListInSiteMessageResponse, error)
	// 查询站内信详情
	GetInSiteMessage(context.Context, *v11.GetInSiteMessageRequest) (*v11.InSiteMessage, error)
	// 创建站内信
	CreateInSiteMessage(context.Context, *v11.CreateInSiteMessageRequest) (*emptypb.Empty, error)
	// 更新站内信
	UpdateInSiteMessage(context.Context, *v11.UpdateInSiteMessageRequest) (*emptypb.Empty, error)
	// 删除站内信
	DeleteInSiteMessage(context.Context, *v11.DeleteInSiteMessageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedInSiteMessageServiceServer()
}

// UnimplementedInSiteMessageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInSiteMessageServiceServer struct{}

func (UnimplementedInSiteMessageServiceServer) ListInSiteMessage(context.Context, *v1.PagingRequest) (*v11.ListInSiteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInSiteMessage not implemented")
}
func (UnimplementedInSiteMessageServiceServer) GetInSiteMessage(context.Context, *v11.GetInSiteMessageRequest) (*v11.InSiteMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInSiteMessage not implemented")
}
func (UnimplementedInSiteMessageServiceServer) CreateInSiteMessage(context.Context, *v11.CreateInSiteMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInSiteMessage not implemented")
}
func (UnimplementedInSiteMessageServiceServer) UpdateInSiteMessage(context.Context, *v11.UpdateInSiteMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInSiteMessage not implemented")
}
func (UnimplementedInSiteMessageServiceServer) DeleteInSiteMessage(context.Context, *v11.DeleteInSiteMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInSiteMessage not implemented")
}
func (UnimplementedInSiteMessageServiceServer) mustEmbedUnimplementedInSiteMessageServiceServer() {}
func (UnimplementedInSiteMessageServiceServer) testEmbeddedByValue()                              {}

// UnsafeInSiteMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InSiteMessageServiceServer will
// result in compilation errors.
type UnsafeInSiteMessageServiceServer interface {
	mustEmbedUnimplementedInSiteMessageServiceServer()
}

func RegisterInSiteMessageServiceServer(s grpc.ServiceRegistrar, srv InSiteMessageServiceServer) {
	// If the following call pancis, it indicates UnimplementedInSiteMessageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InSiteMessageService_ServiceDesc, srv)
}

func _InSiteMessageService_ListInSiteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InSiteMessageServiceServer).ListInSiteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InSiteMessageService_ListInSiteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InSiteMessageServiceServer).ListInSiteMessage(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InSiteMessageService_GetInSiteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetInSiteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InSiteMessageServiceServer).GetInSiteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InSiteMessageService_GetInSiteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InSiteMessageServiceServer).GetInSiteMessage(ctx, req.(*v11.GetInSiteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InSiteMessageService_CreateInSiteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreateInSiteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InSiteMessageServiceServer).CreateInSiteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InSiteMessageService_CreateInSiteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InSiteMessageServiceServer).CreateInSiteMessage(ctx, req.(*v11.CreateInSiteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InSiteMessageService_UpdateInSiteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateInSiteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InSiteMessageServiceServer).UpdateInSiteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InSiteMessageService_UpdateInSiteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InSiteMessageServiceServer).UpdateInSiteMessage(ctx, req.(*v11.UpdateInSiteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InSiteMessageService_DeleteInSiteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteInSiteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InSiteMessageServiceServer).DeleteInSiteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InSiteMessageService_DeleteInSiteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InSiteMessageServiceServer).DeleteInSiteMessage(ctx, req.(*v11.DeleteInSiteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InSiteMessageService_ServiceDesc is the grpc.ServiceDesc for InSiteMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InSiteMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.InSiteMessageService",
	HandlerType: (*InSiteMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInSiteMessage",
			Handler:    _InSiteMessageService_ListInSiteMessage_Handler,
		},
		{
			MethodName: "GetInSiteMessage",
			Handler:    _InSiteMessageService_GetInSiteMessage_Handler,
		},
		{
			MethodName: "CreateInSiteMessage",
			Handler:    _InSiteMessageService_CreateInSiteMessage_Handler,
		},
		{
			MethodName: "UpdateInSiteMessage",
			Handler:    _InSiteMessageService_UpdateInSiteMessage_Handler,
		},
		{
			MethodName: "DeleteInSiteMessage",
			Handler:    _InSiteMessageService_DeleteInSiteMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_in_site_message.proto",
}
