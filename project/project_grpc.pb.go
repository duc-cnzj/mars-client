// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: project/project.proto

package project

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
	Project_List_FullMethodName          = "/project.Project/List"
	Project_Apply_FullMethodName         = "/project.Project/Apply"
	Project_ApplyDryRun_FullMethodName   = "/project.Project/ApplyDryRun"
	Project_Show_FullMethodName          = "/project.Project/Show"
	Project_Version_FullMethodName       = "/project.Project/Version"
	Project_Delete_FullMethodName        = "/project.Project/Delete"
	Project_AllContainers_FullMethodName = "/project.Project/AllContainers"
	Project_HostVariables_FullMethodName = "/project.Project/HostVariables"
)

// ProjectClient is the client API for Project service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectClient interface {
	// List 获取项目列表
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Apply grpc 创建/更新项目
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (Project_ApplyClient, error)
	// ApplyDryRun 创建/更新项目 '--dry-run' mode
	ApplyDryRun(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*DryRunApplyResponse, error)
	// Show 项目详情
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
	// Version 版本号, 如果不存在则返回 0
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Delete 删除项目
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// AllContainers 获取项目下的所有 pod
	AllContainers(ctx context.Context, in *AllContainersRequest, opts ...grpc.CallOption) (*AllContainersResponse, error)
	// HostVariables 获取 hosts 变量
	HostVariables(ctx context.Context, in *HostVariablesRequest, opts ...grpc.CallOption) (*HostVariablesResponse, error)
}

type projectClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectClient(cc grpc.ClientConnInterface) ProjectClient {
	return &projectClient{cc}
}

func (c *projectClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Project_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (Project_ApplyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Project_ServiceDesc.Streams[0], Project_Apply_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &projectApplyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Project_ApplyClient interface {
	Recv() (*ApplyResponse, error)
	grpc.ClientStream
}

type projectApplyClient struct {
	grpc.ClientStream
}

func (x *projectApplyClient) Recv() (*ApplyResponse, error) {
	m := new(ApplyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectClient) ApplyDryRun(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*DryRunApplyResponse, error) {
	out := new(DryRunApplyResponse)
	err := c.cc.Invoke(ctx, Project_ApplyDryRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, Project_Show_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, Project_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Project_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) AllContainers(ctx context.Context, in *AllContainersRequest, opts ...grpc.CallOption) (*AllContainersResponse, error) {
	out := new(AllContainersResponse)
	err := c.cc.Invoke(ctx, Project_AllContainers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) HostVariables(ctx context.Context, in *HostVariablesRequest, opts ...grpc.CallOption) (*HostVariablesResponse, error) {
	out := new(HostVariablesResponse)
	err := c.cc.Invoke(ctx, Project_HostVariables_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServer is the server API for Project service.
// All implementations must embed UnimplementedProjectServer
// for forward compatibility
type ProjectServer interface {
	// List 获取项目列表
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Apply grpc 创建/更新项目
	Apply(*ApplyRequest, Project_ApplyServer) error
	// ApplyDryRun 创建/更新项目 '--dry-run' mode
	ApplyDryRun(context.Context, *ApplyRequest) (*DryRunApplyResponse, error)
	// Show 项目详情
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
	// Version 版本号, 如果不存在则返回 0
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Delete 删除项目
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// AllContainers 获取项目下的所有 pod
	AllContainers(context.Context, *AllContainersRequest) (*AllContainersResponse, error)
	// HostVariables 获取 hosts 变量
	HostVariables(context.Context, *HostVariablesRequest) (*HostVariablesResponse, error)
	mustEmbedUnimplementedProjectServer()
}

// UnimplementedProjectServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServer struct {
}

func (UnimplementedProjectServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProjectServer) Apply(*ApplyRequest, Project_ApplyServer) error {
	return status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedProjectServer) ApplyDryRun(context.Context, *ApplyRequest) (*DryRunApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyDryRun not implemented")
}
func (UnimplementedProjectServer) Show(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedProjectServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedProjectServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProjectServer) AllContainers(context.Context, *AllContainersRequest) (*AllContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllContainers not implemented")
}
func (UnimplementedProjectServer) HostVariables(context.Context, *HostVariablesRequest) (*HostVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostVariables not implemented")
}
func (UnimplementedProjectServer) mustEmbedUnimplementedProjectServer() {}

// UnsafeProjectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServer will
// result in compilation errors.
type UnsafeProjectServer interface {
	mustEmbedUnimplementedProjectServer()
}

func RegisterProjectServer(s grpc.ServiceRegistrar, srv ProjectServer) {
	s.RegisterService(&Project_ServiceDesc, srv)
}

func _Project_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_Apply_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ApplyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServer).Apply(m, &projectApplyServer{stream})
}

type Project_ApplyServer interface {
	Send(*ApplyResponse) error
	grpc.ServerStream
}

type projectApplyServer struct {
	grpc.ServerStream
}

func (x *projectApplyServer) Send(m *ApplyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Project_ApplyDryRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).ApplyDryRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_ApplyDryRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).ApplyDryRun(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_AllContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).AllContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_AllContainers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).AllContainers(ctx, req.(*AllContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_HostVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).HostVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_HostVariables_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).HostVariables(ctx, req.(*HostVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Project_ServiceDesc is the grpc.ServiceDesc for Project service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Project_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.Project",
	HandlerType: (*ProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Project_List_Handler,
		},
		{
			MethodName: "ApplyDryRun",
			Handler:    _Project_ApplyDryRun_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Project_Show_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Project_Version_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Project_Delete_Handler,
		},
		{
			MethodName: "AllContainers",
			Handler:    _Project_AllContainers_Handler,
		},
		{
			MethodName: "HostVariables",
			Handler:    _Project_HostVariables_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Apply",
			Handler:       _Project_Apply_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "project/project.proto",
}
