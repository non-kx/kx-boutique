// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: api/shop/v1/carts.proto

package v1

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
	Carts_GetCart_FullMethodName   = "/api.shop.v1.Carts/GetCart"
	Carts_AddToCart_FullMethodName = "/api.shop.v1.Carts/AddToCart"
)

// CartsClient is the client API for Carts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartsClient interface {
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartReply, error)
	AddToCart(ctx context.Context, in *AddToCartReqeust, opts ...grpc.CallOption) (*AddToCartReply, error)
}

type cartsClient struct {
	cc grpc.ClientConnInterface
}

func NewCartsClient(cc grpc.ClientConnInterface) CartsClient {
	return &cartsClient{cc}
}

func (c *cartsClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartReply, error) {
	out := new(GetCartReply)
	err := c.cc.Invoke(ctx, Carts_GetCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) AddToCart(ctx context.Context, in *AddToCartReqeust, opts ...grpc.CallOption) (*AddToCartReply, error) {
	out := new(AddToCartReply)
	err := c.cc.Invoke(ctx, Carts_AddToCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartsServer is the server API for Carts service.
// All implementations must embed UnimplementedCartsServer
// for forward compatibility
type CartsServer interface {
	GetCart(context.Context, *GetCartRequest) (*GetCartReply, error)
	AddToCart(context.Context, *AddToCartReqeust) (*AddToCartReply, error)
	mustEmbedUnimplementedCartsServer()
}

// UnimplementedCartsServer must be embedded to have forward compatible implementations.
type UnimplementedCartsServer struct {
}

func (UnimplementedCartsServer) GetCart(context.Context, *GetCartRequest) (*GetCartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartsServer) AddToCart(context.Context, *AddToCartReqeust) (*AddToCartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedCartsServer) mustEmbedUnimplementedCartsServer() {}

// UnsafeCartsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartsServer will
// result in compilation errors.
type UnsafeCartsServer interface {
	mustEmbedUnimplementedCartsServer()
}

func RegisterCartsServer(s grpc.ServiceRegistrar, srv CartsServer) {
	s.RegisterService(&Carts_ServiceDesc, srv)
}

func _Carts_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carts_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carts_AddToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).AddToCart(ctx, req.(*AddToCartReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

// Carts_ServiceDesc is the grpc.ServiceDesc for Carts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Carts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.shop.v1.Carts",
	HandlerType: (*CartsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCart",
			Handler:    _Carts_GetCart_Handler,
		},
		{
			MethodName: "AddToCart",
			Handler:    _Carts_AddToCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/shop/v1/carts.proto",
}
