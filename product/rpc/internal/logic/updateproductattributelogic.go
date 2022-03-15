package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeLogic {
	return &UpdateProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductAttributeLogic) UpdateProductAttribute(in *pb.ProductAttribute) (*pb.ProductResponse, error) {
	if l.svcCtx.ProductAttributeModel.Update(in) == false {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}

	return &pb.ProductResponse{
		Code:   200,
		Detail: "更新成功",
	}, nil
}
