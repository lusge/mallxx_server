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

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberServiceClient interface {
	GetMemverLevelList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MemverLeveListResponse, error)
	Login(ctx context.Context, in *MemberLoginRequest, opts ...grpc.CallOption) (*MemberResponse, error)
	Register(ctx context.Context, in *Member, opts ...grpc.CallOption) (*MemberResponse, error)
	GetMemberInfo(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*MemberResponse, error)
	GetMemberList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MemberListResponse, error)
	GetAddress(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*ReceiveAddressListResponse, error)
	AddAddress(ctx context.Context, in *ReceiveAddress, opts ...grpc.CallOption) (*ReceiveAddressResponse, error)
	DelAddress(ctx context.Context, in *ReceiveAddressRequest, opts ...grpc.CallOption) (*Response, error)
	GetFollower(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*FollowerResponse, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) GetMemverLevelList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MemverLeveListResponse, error) {
	out := new(MemverLeveListResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/GetMemverLevelList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Login(ctx context.Context, in *MemberLoginRequest, opts ...grpc.CallOption) (*MemberResponse, error) {
	out := new(MemberResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Register(ctx context.Context, in *Member, opts ...grpc.CallOption) (*MemberResponse, error) {
	out := new(MemberResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetMemberInfo(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*MemberResponse, error) {
	out := new(MemberResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/GetMemberInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetMemberList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MemberListResponse, error) {
	out := new(MemberListResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/GetMemberList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetAddress(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*ReceiveAddressListResponse, error) {
	out := new(ReceiveAddressListResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) AddAddress(ctx context.Context, in *ReceiveAddress, opts ...grpc.CallOption) (*ReceiveAddressResponse, error) {
	out := new(ReceiveAddressResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/AddAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) DelAddress(ctx context.Context, in *ReceiveAddressRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.MemberService/DelAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetFollower(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*FollowerResponse, error) {
	out := new(FollowerResponse)
	err := c.cc.Invoke(ctx, "/pb.MemberService/GetFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
// All implementations must embed UnimplementedMemberServiceServer
// for forward compatibility
type MemberServiceServer interface {
	GetMemverLevelList(context.Context, *EmptyRequest) (*MemverLeveListResponse, error)
	Login(context.Context, *MemberLoginRequest) (*MemberResponse, error)
	Register(context.Context, *Member) (*MemberResponse, error)
	GetMemberInfo(context.Context, *MemberRequest) (*MemberResponse, error)
	GetMemberList(context.Context, *EmptyRequest) (*MemberListResponse, error)
	GetAddress(context.Context, *MemberRequest) (*ReceiveAddressListResponse, error)
	AddAddress(context.Context, *ReceiveAddress) (*ReceiveAddressResponse, error)
	DelAddress(context.Context, *ReceiveAddressRequest) (*Response, error)
	GetFollower(context.Context, *MemberRequest) (*FollowerResponse, error)
	mustEmbedUnimplementedMemberServiceServer()
}

// UnimplementedMemberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServiceServer struct {
}

func (UnimplementedMemberServiceServer) GetMemverLevelList(context.Context, *EmptyRequest) (*MemverLeveListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemverLevelList not implemented")
}
func (UnimplementedMemberServiceServer) Login(context.Context, *MemberLoginRequest) (*MemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMemberServiceServer) Register(context.Context, *Member) (*MemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedMemberServiceServer) GetMemberInfo(context.Context, *MemberRequest) (*MemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberInfo not implemented")
}
func (UnimplementedMemberServiceServer) GetMemberList(context.Context, *EmptyRequest) (*MemberListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberList not implemented")
}
func (UnimplementedMemberServiceServer) GetAddress(context.Context, *MemberRequest) (*ReceiveAddressListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedMemberServiceServer) AddAddress(context.Context, *ReceiveAddress) (*ReceiveAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddress not implemented")
}
func (UnimplementedMemberServiceServer) DelAddress(context.Context, *ReceiveAddressRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAddress not implemented")
}
func (UnimplementedMemberServiceServer) GetFollower(context.Context, *MemberRequest) (*FollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollower not implemented")
}
func (UnimplementedMemberServiceServer) mustEmbedUnimplementedMemberServiceServer() {}

// UnsafeMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServiceServer will
// result in compilation errors.
type UnsafeMemberServiceServer interface {
	mustEmbedUnimplementedMemberServiceServer()
}

func RegisterMemberServiceServer(s grpc.ServiceRegistrar, srv MemberServiceServer) {
	s.RegisterService(&MemberService_ServiceDesc, srv)
}

func _MemberService_GetMemverLevelList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetMemverLevelList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/GetMemverLevelList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetMemverLevelList(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Login(ctx, req.(*MemberLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Register(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetMemberInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetMemberInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/GetMemberInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetMemberInfo(ctx, req.(*MemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetMemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetMemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/GetMemberList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetMemberList(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetAddress(ctx, req.(*MemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_AddAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).AddAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/AddAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).AddAddress(ctx, req.(*ReceiveAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_DelAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).DelAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/DelAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).DelAddress(ctx, req.(*ReceiveAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MemberService/GetFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetFollower(ctx, req.(*MemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberService_ServiceDesc is the grpc.ServiceDesc for MemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMemverLevelList",
			Handler:    _MemberService_GetMemverLevelList_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _MemberService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _MemberService_Register_Handler,
		},
		{
			MethodName: "GetMemberInfo",
			Handler:    _MemberService_GetMemberInfo_Handler,
		},
		{
			MethodName: "GetMemberList",
			Handler:    _MemberService_GetMemberList_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _MemberService_GetAddress_Handler,
		},
		{
			MethodName: "AddAddress",
			Handler:    _MemberService_AddAddress_Handler,
		},
		{
			MethodName: "DelAddress",
			Handler:    _MemberService_DelAddress_Handler,
		},
		{
			MethodName: "GetFollower",
			Handler:    _MemberService_GetFollower_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member.proto",
}