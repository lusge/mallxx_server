package logic

import (
	"context"

	"mallxx_server/advertise/rpc/internal/svc"
	"mallxx_server/advertise/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiGetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApiGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiGetListLogic {
	return &ApiGetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApiGetListLogic) ApiGetList(in *pb.AdRequest) (*pb.AdResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.AdResponse{}, nil
}
