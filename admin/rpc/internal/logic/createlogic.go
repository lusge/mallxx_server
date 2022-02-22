package logic

import (
	"context"

	"mallxx_server/admin/rpc/internal/svc"
	"mallxx_server/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.AdminInfo) (*pb.AdminInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.AdminInfoResponse{}, nil
}
