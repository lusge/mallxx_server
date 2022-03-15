package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeLogic {
	return &DeleteProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductAttributeLogic) DeleteProductAttribute(in *pb.ProductAttribute) (*pb.ProductResponse, error) {
	if l.svcCtx.ProductAttributeModel.DeleteById(in.Id) == false {
		return nil, merrorx.NewCodeError(500, "删除失败")
	}
	return &pb.ProductResponse{
		Code:   200,
		Detail: "删除成功",
	}, nil
}
