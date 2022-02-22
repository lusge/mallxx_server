package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBrandShowStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandShowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandShowStatusLogic {
	return &UpdateBrandShowStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBrandShowStatusLogic) UpdateBrandShowStatus(in *pb.BrandStatusRequest) (*pb.Response, error) {
	if !l.svcCtx.BrandModel.UpdateShowStatus(in.Id, in.Status) {
		return nil, merrorx.NewCodeError(500, "更新成功")
	}

	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
