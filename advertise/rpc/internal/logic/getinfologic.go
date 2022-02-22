package logic

import (
	"context"

	"mallxx_server/advertise/rpc/internal/svc"
	"mallxx_server/advertise/rpc/pb"
	"mallxx_server/common/merrorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *pb.AdOneRequest) (*pb.AdInfoResponse, error) {
	info := l.svcCtx.AdvertiseModel.FindOne(in.Id)
	if info == nil {
		return nil, merrorx.NewCodeError(500, "not found")
	}
	return &pb.AdInfoResponse{
		Code:   200,
		Data:   info,
		Detail: "ok",
	}, nil
}
