package logic

import (
	"context"

	"mallxx_server/advertise/rpc/internal/svc"
	"mallxx_server/advertise/rpc/pb"
	"mallxx_server/common/merrorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *pb.Advertise) (*pb.Response, error) {
	if !l.svcCtx.AdvertiseModel.Update(in) {
		return nil, merrorx.NewCodeError(500, "delete failed")
	}
	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
