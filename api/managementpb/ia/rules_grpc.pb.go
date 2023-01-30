// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/ia/rules.proto

package iav1beta1

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

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RulesClient interface {
	// ListAlertRules returns a list of all Alerting rules.
	ListAlertRules(ctx context.Context, in *ListAlertRulesRequest, opts ...grpc.CallOption) (*ListAlertRulesResponse, error)
	// CreateAlertRule creates Alerting rule.
	CreateAlertRule(ctx context.Context, in *CreateAlertRuleRequest, opts ...grpc.CallOption) (*CreateAlertRuleResponse, error)
	// UpdateAlertRule updates Alerting rule.
	UpdateAlertRule(ctx context.Context, in *UpdateAlertRuleRequest, opts ...grpc.CallOption) (*UpdateAlertRuleResponse, error)
	// ToggleAlertRule allows to switch between disabled and enabled states of an Alert Rule.
	ToggleAlertRule(ctx context.Context, in *ToggleAlertRuleRequest, opts ...grpc.CallOption) (*ToggleAlertRuleResponse, error)
	// DeleteAlertRule deletes Alerting rule.
	DeleteAlertRule(ctx context.Context, in *DeleteAlertRuleRequest, opts ...grpc.CallOption) (*DeleteAlertRuleResponse, error)
}

type rulesClient struct {
	cc grpc.ClientConnInterface
}

func NewRulesClient(cc grpc.ClientConnInterface) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) ListAlertRules(ctx context.Context, in *ListAlertRulesRequest, opts ...grpc.CallOption) (*ListAlertRulesResponse, error) {
	out := new(ListAlertRulesResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Rules/ListAlertRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) CreateAlertRule(ctx context.Context, in *CreateAlertRuleRequest, opts ...grpc.CallOption) (*CreateAlertRuleResponse, error) {
	out := new(CreateAlertRuleResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Rules/CreateAlertRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) UpdateAlertRule(ctx context.Context, in *UpdateAlertRuleRequest, opts ...grpc.CallOption) (*UpdateAlertRuleResponse, error) {
	out := new(UpdateAlertRuleResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Rules/UpdateAlertRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ToggleAlertRule(ctx context.Context, in *ToggleAlertRuleRequest, opts ...grpc.CallOption) (*ToggleAlertRuleResponse, error) {
	out := new(ToggleAlertRuleResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Rules/ToggleAlertRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) DeleteAlertRule(ctx context.Context, in *DeleteAlertRuleRequest, opts ...grpc.CallOption) (*DeleteAlertRuleResponse, error) {
	out := new(DeleteAlertRuleResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Rules/DeleteAlertRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesServer is the server API for Rules service.
// All implementations must embed UnimplementedRulesServer
// for forward compatibility
type RulesServer interface {
	// ListAlertRules returns a list of all Alerting rules.
	ListAlertRules(context.Context, *ListAlertRulesRequest) (*ListAlertRulesResponse, error)
	// CreateAlertRule creates Alerting rule.
	CreateAlertRule(context.Context, *CreateAlertRuleRequest) (*CreateAlertRuleResponse, error)
	// UpdateAlertRule updates Alerting rule.
	UpdateAlertRule(context.Context, *UpdateAlertRuleRequest) (*UpdateAlertRuleResponse, error)
	// ToggleAlertRule allows to switch between disabled and enabled states of an Alert Rule.
	ToggleAlertRule(context.Context, *ToggleAlertRuleRequest) (*ToggleAlertRuleResponse, error)
	// DeleteAlertRule deletes Alerting rule.
	DeleteAlertRule(context.Context, *DeleteAlertRuleRequest) (*DeleteAlertRuleResponse, error)
	mustEmbedUnimplementedRulesServer()
}

// UnimplementedRulesServer must be embedded to have forward compatible implementations.
type UnimplementedRulesServer struct {
}

func (UnimplementedRulesServer) ListAlertRules(context.Context, *ListAlertRulesRequest) (*ListAlertRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertRules not implemented")
}
func (UnimplementedRulesServer) CreateAlertRule(context.Context, *CreateAlertRuleRequest) (*CreateAlertRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertRule not implemented")
}
func (UnimplementedRulesServer) UpdateAlertRule(context.Context, *UpdateAlertRuleRequest) (*UpdateAlertRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertRule not implemented")
}
func (UnimplementedRulesServer) ToggleAlertRule(context.Context, *ToggleAlertRuleRequest) (*ToggleAlertRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleAlertRule not implemented")
}
func (UnimplementedRulesServer) DeleteAlertRule(context.Context, *DeleteAlertRuleRequest) (*DeleteAlertRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertRule not implemented")
}
func (UnimplementedRulesServer) mustEmbedUnimplementedRulesServer() {}

// UnsafeRulesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RulesServer will
// result in compilation errors.
type UnsafeRulesServer interface {
	mustEmbedUnimplementedRulesServer()
}

func RegisterRulesServer(s grpc.ServiceRegistrar, srv RulesServer) {
	s.RegisterService(&Rules_ServiceDesc, srv)
}

func _Rules_ListAlertRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ListAlertRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Rules/ListAlertRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ListAlertRules(ctx, req.(*ListAlertRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_CreateAlertRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).CreateAlertRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Rules/CreateAlertRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).CreateAlertRule(ctx, req.(*CreateAlertRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_UpdateAlertRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).UpdateAlertRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Rules/UpdateAlertRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).UpdateAlertRule(ctx, req.(*UpdateAlertRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ToggleAlertRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleAlertRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ToggleAlertRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Rules/ToggleAlertRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ToggleAlertRule(ctx, req.(*ToggleAlertRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_DeleteAlertRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).DeleteAlertRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Rules/DeleteAlertRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).DeleteAlertRule(ctx, req.(*DeleteAlertRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rules_ServiceDesc is the grpc.ServiceDesc for Rules service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rules_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ia.v1beta1.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlertRules",
			Handler:    _Rules_ListAlertRules_Handler,
		},
		{
			MethodName: "CreateAlertRule",
			Handler:    _Rules_CreateAlertRule_Handler,
		},
		{
			MethodName: "UpdateAlertRule",
			Handler:    _Rules_UpdateAlertRule_Handler,
		},
		{
			MethodName: "ToggleAlertRule",
			Handler:    _Rules_ToggleAlertRule_Handler,
		},
		{
			MethodName: "DeleteAlertRule",
			Handler:    _Rules_DeleteAlertRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/ia/rules.proto",
}
