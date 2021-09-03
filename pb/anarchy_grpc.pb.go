// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AnarchyClient is the client API for Anarchy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnarchyClient interface {
	// Elements
	GetElem(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyElement, error)
	GetCombination(ctx context.Context, in *AnarchyCombination, opts ...grpc.CallOption) (*AnarchyCombinationResult, error)
	// Savefile
	GetInv(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyInventory, error)
	AddFound(ctx context.Context, in *AnarchyUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type anarchyClient struct {
	cc grpc.ClientConnInterface
}

func NewAnarchyClient(cc grpc.ClientConnInterface) AnarchyClient {
	return &anarchyClient{cc}
}

func (c *anarchyClient) GetElem(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyElement, error) {
	out := new(AnarchyElement)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/GetElem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) GetCombination(ctx context.Context, in *AnarchyCombination, opts ...grpc.CallOption) (*AnarchyCombinationResult, error) {
	out := new(AnarchyCombinationResult)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/GetCombination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) GetInv(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyInventory, error) {
	out := new(AnarchyInventory)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/GetInv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) AddFound(ctx context.Context, in *AnarchyUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/AddFound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnarchyServer is the server API for Anarchy service.
// All implementations must embed UnimplementedAnarchyServer
// for forward compatibility
type AnarchyServer interface {
	// Elements
	GetElem(context.Context, *wrapperspb.StringValue) (*AnarchyElement, error)
	GetCombination(context.Context, *AnarchyCombination) (*AnarchyCombinationResult, error)
	// Savefile
	GetInv(context.Context, *wrapperspb.StringValue) (*AnarchyInventory, error)
	AddFound(context.Context, *AnarchyUserRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAnarchyServer()
}

// UnimplementedAnarchyServer must be embedded to have forward compatible implementations.
type UnimplementedAnarchyServer struct {
}

func (UnimplementedAnarchyServer) GetElem(context.Context, *wrapperspb.StringValue) (*AnarchyElement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetElem not implemented")
}
func (UnimplementedAnarchyServer) GetCombination(context.Context, *AnarchyCombination) (*AnarchyCombinationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCombination not implemented")
}
func (UnimplementedAnarchyServer) GetInv(context.Context, *wrapperspb.StringValue) (*AnarchyInventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInv not implemented")
}
func (UnimplementedAnarchyServer) AddFound(context.Context, *AnarchyUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFound not implemented")
}
func (UnimplementedAnarchyServer) mustEmbedUnimplementedAnarchyServer() {}

// UnsafeAnarchyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnarchyServer will
// result in compilation errors.
type UnsafeAnarchyServer interface {
	mustEmbedUnimplementedAnarchyServer()
}

func RegisterAnarchyServer(s grpc.ServiceRegistrar, srv AnarchyServer) {
	s.RegisterService(&Anarchy_ServiceDesc, srv)
}

func _Anarchy_GetElem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).GetElem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/GetElem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnarchyServer).GetElem(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_GetCombination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnarchyCombination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).GetCombination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/GetCombination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnarchyServer).GetCombination(ctx, req.(*AnarchyCombination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_GetInv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).GetInv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/GetInv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnarchyServer).GetInv(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_AddFound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnarchyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).AddFound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/AddFound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnarchyServer).AddFound(ctx, req.(*AnarchyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Anarchy_ServiceDesc is the grpc.ServiceDesc for Anarchy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Anarchy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anarchy.Anarchy",
	HandlerType: (*AnarchyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetElem",
			Handler:    _Anarchy_GetElem_Handler,
		},
		{
			MethodName: "GetCombination",
			Handler:    _Anarchy_GetCombination_Handler,
		},
		{
			MethodName: "GetInv",
			Handler:    _Anarchy_GetInv_Handler,
		},
		{
			MethodName: "AddFound",
			Handler:    _Anarchy_AddFound_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/anarchy.proto",
}
