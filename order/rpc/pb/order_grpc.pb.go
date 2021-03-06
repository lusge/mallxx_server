// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// OrderserviceClient is the client API for Orderservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderserviceClient interface {
	OrderFrim(ctx context.Context, in *OrderFirmRequest, opts ...grpc.CallOption) (*OrderFirmResponse, error)
	OrderGenerate(ctx context.Context, in *GenerateOrderRequest, opts ...grpc.CallOption) (*GenerateOrderResponse, error)
	OrderMemberList(ctx context.Context, in *OrderMemberRequest, opts ...grpc.CallOption) (*OrderListResponse, error)
	OrderInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderInfoRsponse, error)
	OrderCancel(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	OrderDelete(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderserviceClient(cc grpc.ClientConnInterface) OrderserviceClient {
	return &orderserviceClient{cc}
}

func (c *orderserviceClient) OrderFrim(ctx context.Context, in *OrderFirmRequest, opts ...grpc.CallOption) (*OrderFirmResponse, error) {
	out := new(OrderFirmResponse)
	err := c.cc.Invoke(ctx, "/pb.Orderservice/OrderFrim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderserviceClient) OrderGenerate(ctx context.Context, in *GenerateOrderRequest, opts ...grpc.CallOption) (*GenerateOrderResponse, error) {
	out := new(GenerateOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.Orderservice/OrderGenerate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderserviceClient) OrderMemberList(ctx context.Context, in *OrderMemberRequest, opts ...grpc.CallOption) (*OrderListResponse, error) {
	out := new(OrderListResponse)
	err := c.cc.Invoke(ctx, "/pb.Orderservice/OrderMemberList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderserviceClient) OrderInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderInfoRsponse, error) {
	out := new(OrderInfoRsponse)
	err := c.cc.Invoke(ctx, "/pb.Orderservice/OrderInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderserviceClient) OrderCancel(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/pb.Orderservice/OrderCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderserviceClient) OrderDelete(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/pb.Orderservice/OrderDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderserviceServer is the server API for Orderservice service.
// All implementations must embed UnimplementedOrderserviceServer
// for forward compatibility
type OrderserviceServer interface {
	OrderFrim(context.Context, *OrderFirmRequest) (*OrderFirmResponse, error)
	OrderGenerate(context.Context, *GenerateOrderRequest) (*GenerateOrderResponse, error)
	OrderMemberList(context.Context, *OrderMemberRequest) (*OrderListResponse, error)
	OrderInfo(context.Context, *OrderInfoRequest) (*OrderInfoRsponse, error)
	OrderCancel(context.Context, *OrderInfoRequest) (*OrderResponse, error)
	OrderDelete(context.Context, *OrderInfoRequest) (*OrderResponse, error)
	mustEmbedUnimplementedOrderserviceServer()
}

// UnimplementedOrderserviceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderserviceServer struct {
}

func (UnimplementedOrderserviceServer) OrderFrim(context.Context, *OrderFirmRequest) (*OrderFirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderFrim not implemented")
}
func (UnimplementedOrderserviceServer) OrderGenerate(context.Context, *GenerateOrderRequest) (*GenerateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderGenerate not implemented")
}
func (UnimplementedOrderserviceServer) OrderMemberList(context.Context, *OrderMemberRequest) (*OrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderMemberList not implemented")
}
func (UnimplementedOrderserviceServer) OrderInfo(context.Context, *OrderInfoRequest) (*OrderInfoRsponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderInfo not implemented")
}
func (UnimplementedOrderserviceServer) OrderCancel(context.Context, *OrderInfoRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCancel not implemented")
}
func (UnimplementedOrderserviceServer) OrderDelete(context.Context, *OrderInfoRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDelete not implemented")
}
func (UnimplementedOrderserviceServer) mustEmbedUnimplementedOrderserviceServer() {}

// UnsafeOrderserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderserviceServer will
// result in compilation errors.
type UnsafeOrderserviceServer interface {
	mustEmbedUnimplementedOrderserviceServer()
}

func RegisterOrderserviceServer(s grpc.ServiceRegistrar, srv OrderserviceServer) {
	s.RegisterService(&Orderservice_ServiceDesc, srv)
}

func _Orderservice_OrderFrim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderserviceServer).OrderFrim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Orderservice/OrderFrim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderserviceServer).OrderFrim(ctx, req.(*OrderFirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderservice_OrderGenerate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderserviceServer).OrderGenerate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Orderservice/OrderGenerate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderserviceServer).OrderGenerate(ctx, req.(*GenerateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderservice_OrderMemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderserviceServer).OrderMemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Orderservice/OrderMemberList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderserviceServer).OrderMemberList(ctx, req.(*OrderMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderservice_OrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderserviceServer).OrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Orderservice/OrderInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderserviceServer).OrderInfo(ctx, req.(*OrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderservice_OrderCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderserviceServer).OrderCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Orderservice/OrderCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderserviceServer).OrderCancel(ctx, req.(*OrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderservice_OrderDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderserviceServer).OrderDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Orderservice/OrderDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderserviceServer).OrderDelete(ctx, req.(*OrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Orderservice_ServiceDesc is the grpc.ServiceDesc for Orderservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orderservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Orderservice",
	HandlerType: (*OrderserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderFrim",
			Handler:    _Orderservice_OrderFrim_Handler,
		},
		{
			MethodName: "OrderGenerate",
			Handler:    _Orderservice_OrderGenerate_Handler,
		},
		{
			MethodName: "OrderMemberList",
			Handler:    _Orderservice_OrderMemberList_Handler,
		},
		{
			MethodName: "OrderInfo",
			Handler:    _Orderservice_OrderInfo_Handler,
		},
		{
			MethodName: "OrderCancel",
			Handler:    _Orderservice_OrderCancel_Handler,
		},
		{
			MethodName: "OrderDelete",
			Handler:    _Orderservice_OrderDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/rpc/order.proto",
}
