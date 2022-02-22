package logic

import (
	"context"

	"mallxx_server/advertise/rpc/internal/svc"
	"mallxx_server/advertise/rpc/pb"
	"mallxx_server/common/merrorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.Advertise) (*pb.Response, error) {
	if !l.svcCtx.AdvertiseModel.Insert(in) {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}
	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
