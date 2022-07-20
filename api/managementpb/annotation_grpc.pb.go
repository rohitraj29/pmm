// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/annotation.proto

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

// AnnotationClient is the client API for Annotation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnnotationClient interface {
	// AddAnnotation adds annotation.
	AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error)
}

type annotationClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnotationClient(cc grpc.ClientConnInterface) AnnotationClient {
	return &annotationClient{cc}
}

func (c *annotationClient) AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error) {
	out := new(AddAnnotationResponse)
	err := c.cc.Invoke(ctx, "/management.Annotation/AddAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnotationServer is the server API for Annotation service.
// All implementations must embed UnimplementedAnnotationServer
// for forward compatibility
type AnnotationServer interface {
	// AddAnnotation adds annotation.
	AddAnnotation(context.Context, *AddAnnotationRequest) (*AddAnnotationResponse, error)
	mustEmbedUnimplementedAnnotationServer()
}

// UnimplementedAnnotationServer must be embedded to have forward compatible implementations.
type UnimplementedAnnotationServer struct {
}

func (UnimplementedAnnotationServer) AddAnnotation(context.Context, *AddAnnotationRequest) (*AddAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnnotation not implemented")
}
func (UnimplementedAnnotationServer) mustEmbedUnimplementedAnnotationServer() {}

// UnsafeAnnotationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnnotationServer will
// result in compilation errors.
type UnsafeAnnotationServer interface {
	mustEmbedUnimplementedAnnotationServer()
}

func RegisterAnnotationServer(s grpc.ServiceRegistrar, srv AnnotationServer) {
	s.RegisterService(&Annotation_ServiceDesc, srv)
}

func _Annotation_AddAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnotationServer).AddAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Annotation/AddAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnotationServer).AddAnnotation(ctx, req.(*AddAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Annotation_ServiceDesc is the grpc.ServiceDesc for Annotation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Annotation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.Annotation",
	HandlerType: (*AnnotationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAnnotation",
			Handler:    _Annotation_AddAnnotation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/annotation.proto",
}
