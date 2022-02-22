package logic

import (
	"context"

	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberInfoLogic {
	return &GetMemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberInfoLogic) GetMemberInfo(in *pb.MemberRequest) (*pb.MemberResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.MemberResponse{}, nil
}
