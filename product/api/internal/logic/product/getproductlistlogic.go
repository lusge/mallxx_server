package product

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductListLogic {
	return GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListLogic) GetProductList(req types.ProductListRequest) (*types.ProductListRespone, error) {
	var reqList productservices.ProductListRequest
	copier.Copy(&reqList, req)
	resp, err := l.svcCtx.ProductRpc.GetProductList(l.ctx, &reqList)
	if err != nil {
		return nil, err
	}

	var data []*types.Product
	copier.Copy(&data, resp.Data)
	return &types.ProductListRespone{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
