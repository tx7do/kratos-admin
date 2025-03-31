// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_file.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-admin/api/gen/go/file/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileService_ListFile_FullMethodName   = "/admin.service.v1.FileService/ListFile"
	FileService_GetFile_FullMethodName    = "/admin.service.v1.FileService/GetFile"
	FileService_CreateFile_FullMethodName = "/admin.service.v1.FileService/CreateFile"
	FileService_UpdateFile_FullMethodName = "/admin.service.v1.FileService/UpdateFile"
	FileService_DeleteFile_FullMethodName = "/admin.service.v1.FileService/DeleteFile"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 文件管理服务
type FileServiceClient interface {
	// 查询文件列表
	ListFile(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListFileResponse, error)
	// 查询文件详情
	GetFile(ctx context.Context, in *v11.GetFileRequest, opts ...grpc.CallOption) (*v11.File, error)
	// 创建文件
	CreateFile(ctx context.Context, in *v11.CreateFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新文件
	UpdateFile(ctx context.Context, in *v11.UpdateFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除文件
	DeleteFile(ctx context.Context, in *v11.DeleteFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) ListFile(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.ListFileResponse)
	err := c.cc.Invoke(ctx, FileService_ListFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFile(ctx context.Context, in *v11.GetFileRequest, opts ...grpc.CallOption) (*v11.File, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.File)
	err := c.cc.Invoke(ctx, FileService_GetFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CreateFile(ctx context.Context, in *v11.CreateFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FileService_CreateFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateFile(ctx context.Context, in *v11.UpdateFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FileService_UpdateFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteFile(ctx context.Context, in *v11.DeleteFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FileService_DeleteFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
//
// 文件管理服务
type FileServiceServer interface {
	// 查询文件列表
	ListFile(context.Context, *v1.PagingRequest) (*v11.ListFileResponse, error)
	// 查询文件详情
	GetFile(context.Context, *v11.GetFileRequest) (*v11.File, error)
	// 创建文件
	CreateFile(context.Context, *v11.CreateFileRequest) (*emptypb.Empty, error)
	// 更新文件
	UpdateFile(context.Context, *v11.UpdateFileRequest) (*emptypb.Empty, error)
	// 删除文件
	DeleteFile(context.Context, *v11.DeleteFileRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) ListFile(context.Context, *v1.PagingRequest) (*v11.ListFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}
func (UnimplementedFileServiceServer) GetFile(context.Context, *v11.GetFileRequest) (*v11.File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileServiceServer) CreateFile(context.Context, *v11.CreateFileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedFileServiceServer) UpdateFile(context.Context, *v11.UpdateFileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFileServiceServer) DeleteFile(context.Context, *v11.DeleteFileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_ListFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ListFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_ListFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ListFile(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFile(ctx, req.(*v11.GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_CreateFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateFile(ctx, req.(*v11.CreateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_UpdateFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateFile(ctx, req.(*v11.UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteFile(ctx, req.(*v11.DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFile",
			Handler:    _FileService_ListFile_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _FileService_GetFile_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _FileService_CreateFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _FileService_UpdateFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileService_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_file.proto",
}
