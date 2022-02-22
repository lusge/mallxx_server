package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBrandFactoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandFactoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandFactoryStatusLogic {
	return &UpdateBrandFactoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBrandFactoryStatusLogic) UpdateBrandFactoryStatus(in *pb.BrandStatusRequest) (*pb.Response, error) {
	if l.svcCtx.BrandModel.UpdateFactoryStatus(in.Id, in.Status) == false {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}

	return &pb.Response{
		Code:   200,
		Detail: "更新成功",
	}, nil
}
