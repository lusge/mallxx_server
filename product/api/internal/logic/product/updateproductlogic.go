package product

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateProductLogic {
	return UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req types.Product) (*types.Response, error) {
	var product productservices.Product
	err := copier.Copy(&product, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, "参数错误")
	}

	resp, err := l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &product)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
