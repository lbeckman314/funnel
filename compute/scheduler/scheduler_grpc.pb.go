// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: scheduler.proto

package scheduler

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

// SchedulerServiceClient is the client API for SchedulerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerServiceClient interface {
	PutNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*PutNodeResponse, error)
	DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error)
}

type schedulerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerServiceClient(cc grpc.ClientConnInterface) SchedulerServiceClient {
	return &schedulerServiceClient{cc}
}

func (c *schedulerServiceClient) PutNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*PutNodeResponse, error) {
	out := new(PutNodeResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerService/PutNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerService/DeleteNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerService/ListNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerService/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServiceServer is the server API for SchedulerService service.
// All implementations must embed UnimplementedSchedulerServiceServer
// for forward compatibility
type SchedulerServiceServer interface {
	PutNode(context.Context, *Node) (*PutNodeResponse, error)
	DeleteNode(context.Context, *Node) (*DeleteNodeResponse, error)
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	GetNode(context.Context, *GetNodeRequest) (*Node, error)
	mustEmbedUnimplementedSchedulerServiceServer()
}

// UnimplementedSchedulerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerServiceServer struct {
}

func (UnimplementedSchedulerServiceServer) PutNode(context.Context, *Node) (*PutNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutNode not implemented")
}
func (UnimplementedSchedulerServiceServer) DeleteNode(context.Context, *Node) (*DeleteNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedSchedulerServiceServer) ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedSchedulerServiceServer) GetNode(context.Context, *GetNodeRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedSchedulerServiceServer) mustEmbedUnimplementedSchedulerServiceServer() {}

// UnsafeSchedulerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServiceServer will
// result in compilation errors.
type UnsafeSchedulerServiceServer interface {
	mustEmbedUnimplementedSchedulerServiceServer()
}

func RegisterSchedulerServiceServer(s grpc.ServiceRegistrar, srv SchedulerServiceServer) {
	s.RegisterService(&SchedulerService_ServiceDesc, srv)
}

func _SchedulerService_PutNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).PutNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/PutNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).PutNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).DeleteNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchedulerService_ServiceDesc is the grpc.ServiceDesc for SchedulerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchedulerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.SchedulerService",
	HandlerType: (*SchedulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutNode",
			Handler:    _SchedulerService_PutNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _SchedulerService_DeleteNode_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _SchedulerService_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _SchedulerService_GetNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}
