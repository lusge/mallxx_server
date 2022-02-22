package productAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductAttributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddProductAttributeLogic {
	return AddProductAttributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductAttributeLogic) AddProductAttribute(req types.ProductAttribute) (*types.Response, error) {
	var attr productservices.ProductAttribute
	copier.Copy(&attr, req)
	resp, err := l.svcCtx.ProductRpc.AddProductAttribute(l.ctx, &attr)

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
