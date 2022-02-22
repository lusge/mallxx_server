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
	list := l.svcCtx.AdvertiseModel.FindAll(in.Pos, in.CategoryId)

	return &pb.AdResponse{
		Code:   200,
		Detail: "ok",
		Data:   list,
	}, nil
}
