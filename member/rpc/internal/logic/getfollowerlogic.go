package logic

import (
	"context"

	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerLogic {
	return &GetFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerLogic) GetFollower(in *pb.MemberRequest) (*pb.FollowerResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.FollowerResponse{}, nil
}
