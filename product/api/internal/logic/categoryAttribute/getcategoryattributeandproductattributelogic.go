package categoryAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryAttributeAndProductAttributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryAttributeAndProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryAttributeAndProductAttributeLogic {
	return GetCategoryAttributeAndProductAttributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryAttributeAndProductAttributeLogic) GetCategoryAttributeAndProductAttribute() (*types.CategoryAttrAndProductAttrListResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetCategoryAttributeAndProductAttribute(l.ctx, &productservices.ProductEmptyRequest{})

	if err != nil {
		return nil, err
	}

	var data []*types.CategoryAttrAndProductAttr
	copier.Copy(&data, resp.Data)

	return &types.CategoryAttrAndProductAttrListResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
