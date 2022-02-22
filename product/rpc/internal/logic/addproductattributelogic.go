package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeLogic {
	return &AddProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductAttributeLogic) AddProductAttribute(in *pb.ProductAttribute) (*pb.Response, error) {
	if l.svcCtx.ProductAttributeModel.Insert(in) == false {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}

	return &pb.Response{
		Code:   200,
		Detail: "新增成功",
	}, nil
}
