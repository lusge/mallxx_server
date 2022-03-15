package logic

import (
	"context"

	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemverLevelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemverLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemverLevelListLogic {
	return &GetMemverLevelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemverLevelListLogic) GetMemverLevelList(in *pb.MemberEmptyRequest) (*pb.MemverLeveListResponse, error) {
	list := l.svcCtx.MemberLevelModel.FindAll()

	return &pb.MemverLeveListResponse{
		Code:   200,
		Detail: "ok",
		Data:   list,
	}, nil
}
