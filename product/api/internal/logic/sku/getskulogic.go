package sku

import (
	"context"
	"fmt"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSkuLogic {
	return GetSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSkuLogic) GetSku(req types.SkuStockRequest) (*types.SkuStockListResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetSku(l.ctx, &productservices.SkuStockRequest{
		ProductId: req.ProductId,
		Keyword:   req.Keyword,
	})

	if err != nil {
		return nil, err
	}

	var data []*types.SkuStock

	copier.Copy(&data, resp.Data)

	fmt.Println(resp.Data)

	return &types.SkuStockListResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
