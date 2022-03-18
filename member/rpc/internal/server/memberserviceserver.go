// Code generated by goctl. DO NOT EDIT!
// Source: member.proto

package server

import (
	"context"

	"mallxx_server/member/rpc/internal/logic"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"
)

type MemberServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMemberServiceServer
}

func NewMemberServiceServer(svcCtx *svc.ServiceContext) *MemberServiceServer {
	return &MemberServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MemberServiceServer) GetMemverLevelList(ctx context.Context, in *pb.MemberEmptyRequest) (*pb.MemverLeveListResponse, error) {
	l := logic.NewGetMemverLevelListLogic(ctx, s.svcCtx)
	return l.GetMemverLevelList(in)
}

func (s *MemberServiceServer) GetMemberInfo(ctx context.Context, in *pb.MemberRequest) (*pb.MemberInfoResponse, error) {
	l := logic.NewGetMemberInfoLogic(ctx, s.svcCtx)
	return l.GetMemberInfo(in)
}

func (s *MemberServiceServer) GetMemberList(ctx context.Context, in *pb.MemberEmptyRequest) (*pb.MemberListResponse, error) {
	l := logic.NewGetMemberListLogic(ctx, s.svcCtx)
	return l.GetMemberList(in)
}

func (s *MemberServiceServer) GetAddress(ctx context.Context, in *pb.MemberRequest) (*pb.ReceiveAddressListResponse, error) {
	l := logic.NewGetAddressLogic(ctx, s.svcCtx)
	return l.GetAddress(in)
}

func (s *MemberServiceServer) AddAddress(ctx context.Context, in *pb.ReceiveAddress) (*pb.ReceiveAddressResponse, error) {
	l := logic.NewAddAddressLogic(ctx, s.svcCtx)
	return l.AddAddress(in)
}

func (s *MemberServiceServer) DelAddress(ctx context.Context, in *pb.ReceiveAddressRequest) (*pb.MemberResponse, error) {
	l := logic.NewDelAddressLogic(ctx, s.svcCtx)
	return l.DelAddress(in)
}

func (s *MemberServiceServer) SetDefaultAddress(ctx context.Context, in *pb.ReceiveAddressRequest) (*pb.MemberResponse, error) {
	l := logic.NewSetDefaultAddressLogic(ctx, s.svcCtx)
	return l.SetDefaultAddress(in)
}

func (s *MemberServiceServer) UpdateAddress(ctx context.Context, in *pb.ReceiveAddress) (*pb.MemberResponse, error) {
	l := logic.NewUpdateAddressLogic(ctx, s.svcCtx)
	return l.UpdateAddress(in)
}

func (s *MemberServiceServer) GetAddressInfo(ctx context.Context, in *pb.ReceiveAddressRequest) (*pb.ReceiveAddressResponse, error) {
	l := logic.NewGetAddressInfoLogic(ctx, s.svcCtx)
	return l.GetAddressInfo(in)
}

func (s *MemberServiceServer) GetFollower(ctx context.Context, in *pb.MemberRequest) (*pb.FollowerResponse, error) {
	l := logic.NewGetFollowerLogic(ctx, s.svcCtx)
	return l.GetFollower(in)
}
