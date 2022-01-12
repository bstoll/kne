// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package controller

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

// TopologyManagerClient is the client API for TopologyManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopologyManagerClient interface {
	CreateTopology(ctx context.Context, in *CreateTopologyRequest, opts ...grpc.CallOption) (*CreateTopologyResponse, error)
	DeleteTopology(ctx context.Context, in *DeleteTopologyRequest, opts ...grpc.CallOption) (*DeleteTopologyResponse, error)
	ShowTopology(ctx context.Context, in *ShowTopologyRequest, opts ...grpc.CallOption) (*ShowTopologyResponse, error)
}

type topologyManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTopologyManagerClient(cc grpc.ClientConnInterface) TopologyManagerClient {
	return &topologyManagerClient{cc}
}

func (c *topologyManagerClient) CreateTopology(ctx context.Context, in *CreateTopologyRequest, opts ...grpc.CallOption) (*CreateTopologyResponse, error) {
	out := new(CreateTopologyResponse)
	err := c.cc.Invoke(ctx, "/controller.TopologyManager/CreateTopology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyManagerClient) DeleteTopology(ctx context.Context, in *DeleteTopologyRequest, opts ...grpc.CallOption) (*DeleteTopologyResponse, error) {
	out := new(DeleteTopologyResponse)
	err := c.cc.Invoke(ctx, "/controller.TopologyManager/DeleteTopology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyManagerClient) ShowTopology(ctx context.Context, in *ShowTopologyRequest, opts ...grpc.CallOption) (*ShowTopologyResponse, error) {
	out := new(ShowTopologyResponse)
	err := c.cc.Invoke(ctx, "/controller.TopologyManager/ShowTopology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopologyManagerServer is the server API for TopologyManager service.
// All implementations must embed UnimplementedTopologyManagerServer
// for forward compatibility
type TopologyManagerServer interface {
	CreateTopology(context.Context, *CreateTopologyRequest) (*CreateTopologyResponse, error)
	DeleteTopology(context.Context, *DeleteTopologyRequest) (*DeleteTopologyResponse, error)
	ShowTopology(context.Context, *ShowTopologyRequest) (*ShowTopologyResponse, error)
	mustEmbedUnimplementedTopologyManagerServer()
}

// UnimplementedTopologyManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTopologyManagerServer struct {
}

func (UnimplementedTopologyManagerServer) CreateTopology(context.Context, *CreateTopologyRequest) (*CreateTopologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopology not implemented")
}
func (UnimplementedTopologyManagerServer) DeleteTopology(context.Context, *DeleteTopologyRequest) (*DeleteTopologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopology not implemented")
}
func (UnimplementedTopologyManagerServer) ShowTopology(context.Context, *ShowTopologyRequest) (*ShowTopologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowTopology not implemented")
}
func (UnimplementedTopologyManagerServer) mustEmbedUnimplementedTopologyManagerServer() {}

// UnsafeTopologyManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopologyManagerServer will
// result in compilation errors.
type UnsafeTopologyManagerServer interface {
	mustEmbedUnimplementedTopologyManagerServer()
}

func RegisterTopologyManagerServer(s grpc.ServiceRegistrar, srv TopologyManagerServer) {
	s.RegisterService(&TopologyManager_ServiceDesc, srv)
}

func _TopologyManager_CreateTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyManagerServer).CreateTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.TopologyManager/CreateTopology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyManagerServer).CreateTopology(ctx, req.(*CreateTopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyManager_DeleteTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyManagerServer).DeleteTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.TopologyManager/DeleteTopology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyManagerServer).DeleteTopology(ctx, req.(*DeleteTopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyManager_ShowTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowTopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyManagerServer).ShowTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.TopologyManager/ShowTopology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyManagerServer).ShowTopology(ctx, req.(*ShowTopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TopologyManager_ServiceDesc is the grpc.ServiceDesc for TopologyManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopologyManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controller.TopologyManager",
	HandlerType: (*TopologyManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopology",
			Handler:    _TopologyManager_CreateTopology_Handler,
		},
		{
			MethodName: "DeleteTopology",
			Handler:    _TopologyManager_DeleteTopology_Handler,
		},
		{
			MethodName: "ShowTopology",
			Handler:    _TopologyManager_ShowTopology_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller.proto",
}