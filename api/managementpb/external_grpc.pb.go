// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/external.proto

package managementpb

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

// ExternalClient is the client API for External service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalClient interface {
	// AddExternal adds external service and adds external exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds an "external exporter" agent to inventory, which is running on provided "runs_on_node_id".
	AddExternal(ctx context.Context, in *AddExternalRequest, opts ...grpc.CallOption) (*AddExternalResponse, error)
}

type externalClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalClient(cc grpc.ClientConnInterface) ExternalClient {
	return &externalClient{cc}
}

func (c *externalClient) AddExternal(ctx context.Context, in *AddExternalRequest, opts ...grpc.CallOption) (*AddExternalResponse, error) {
	out := new(AddExternalResponse)
	err := c.cc.Invoke(ctx, "/management.External/AddExternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalServer is the server API for External service.
// All implementations must embed UnimplementedExternalServer
// for forward compatibility
type ExternalServer interface {
	// AddExternal adds external service and adds external exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds an "external exporter" agent to inventory, which is running on provided "runs_on_node_id".
	AddExternal(context.Context, *AddExternalRequest) (*AddExternalResponse, error)
	mustEmbedUnimplementedExternalServer()
}

// UnimplementedExternalServer must be embedded to have forward compatible implementations.
type UnimplementedExternalServer struct{}

func (UnimplementedExternalServer) AddExternal(context.Context, *AddExternalRequest) (*AddExternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExternal not implemented")
}
func (UnimplementedExternalServer) mustEmbedUnimplementedExternalServer() {}

// UnsafeExternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalServer will
// result in compilation errors.
type UnsafeExternalServer interface {
	mustEmbedUnimplementedExternalServer()
}

func RegisterExternalServer(s grpc.ServiceRegistrar, srv ExternalServer) {
	s.RegisterService(&External_ServiceDesc, srv)
}

func _External_AddExternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServer).AddExternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.External/AddExternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServer).AddExternal(ctx, req.(*AddExternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// External_ServiceDesc is the grpc.ServiceDesc for External service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var External_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.External",
	HandlerType: (*ExternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddExternal",
			Handler:    _External_AddExternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/external.proto",
}
