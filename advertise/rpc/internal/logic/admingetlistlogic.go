package logic

import (
	"context"

	"mallxx_server/advertise/rpc/internal/svc"
	"mallxx_server/advertise/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetListLogic {
	return &AdminGetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminGetListLogic) AdminGetList(in *pb.AdRequest) (*pb.AdResponse, error) {
	start, pageSize := utils.Pager(in.PageNum, in.PageSize)

	list := l.svcCtx.AdvertiseModel.Find(in, start, pageSize)
	total := l.svcCtx.AdvertiseModel.Count(in)

	return &pb.AdResponse{
		Total:  total,
		Code:   200,
		Detail: "ok",
		Data:   list,
	}, nil
}
