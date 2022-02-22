// Code generated by goctl. DO NOT EDIT!
// Source: member.proto

package memberservice

import (
	"context"

	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmptyRequest               = pb.EmptyRequest
	FollowerResponse           = pb.FollowerResponse
	ListRequest                = pb.ListRequest
	Member                     = pb.Member
	MemberData                 = pb.MemberData
	MemberLevel                = pb.MemberLevel
	MemberListResponse         = pb.MemberListResponse
	MemberLoginRequest         = pb.MemberLoginRequest
	MemberRequest              = pb.MemberRequest
	MemberResponse             = pb.MemberResponse
	MemverLeveListResponse     = pb.MemverLeveListResponse
	MemverLevelInfoRequest     = pb.MemverLevelInfoRequest
	ReceiveAddress             = pb.ReceiveAddress
	ReceiveAddressListResponse = pb.ReceiveAddressListResponse
	ReceiveAddressRequest      = pb.ReceiveAddressRequest
	ReceiveAddressResponse     = pb.ReceiveAddressResponse
	Response                   = pb.Response

	MemberService interface {
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

	defaultMemberService struct {
		cli zrpc.Client
	}
)

func NewMemberService(cli zrpc.Client) MemberService {
	return &defaultMemberService{
		cli: cli,
	}
}

func (m *defaultMemberService) GetMemverLevelList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MemverLeveListResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.GetMemverLevelList(ctx, in, opts...)
}

func (m *defaultMemberService) Login(ctx context.Context, in *MemberLoginRequest, opts ...grpc.CallOption) (*MemberResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultMemberService) Register(ctx context.Context, in *Member, opts ...grpc.CallOption) (*MemberResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultMemberService) GetMemberInfo(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*MemberResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.GetMemberInfo(ctx, in, opts...)
}

func (m *defaultMemberService) GetMemberList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*MemberListResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.GetMemberList(ctx, in, opts...)
}

func (m *defaultMemberService) GetAddress(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*ReceiveAddressListResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.GetAddress(ctx, in, opts...)
}

func (m *defaultMemberService) AddAddress(ctx context.Context, in *ReceiveAddress, opts ...grpc.CallOption) (*ReceiveAddressResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.AddAddress(ctx, in, opts...)
}

func (m *defaultMemberService) DelAddress(ctx context.Context, in *ReceiveAddressRequest, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.DelAddress(ctx, in, opts...)
}

func (m *defaultMemberService) GetFollower(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*FollowerResponse, error) {
	client := pb.NewMemberServiceClient(m.cli.Conn())
	return client.GetFollower(ctx, in, opts...)
}