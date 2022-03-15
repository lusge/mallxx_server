package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandLogic {
	return &UpdateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBrandLogic) UpdateBrand(in *pb.Brand) (*pb.ProductResponse, error) {
	if l.svcCtx.BrandModel.Update(in) == false {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}

	return &pb.ProductResponse{
		Code:   200,
		Detail: "新增成功",
	}, nil
}
