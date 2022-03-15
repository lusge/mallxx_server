package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *pb.Product) (*pb.ProductResponse, error) {
	isOk := l.svcCtx.ProductModel.Update(in)
	if !isOk {
		return nil, merrorx.NewCodeError(500, "更改失败")
	}
	return &pb.ProductResponse{
		Code:   200,
		Detail: "更新成功",
	}, nil
}
