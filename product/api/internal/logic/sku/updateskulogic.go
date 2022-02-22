package sku

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateSkuLogic {
	return UpdateSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSkuLogic) UpdateSku(req types.SkuStockListRequest) (*types.Response, error) {
	var rpcReq productservices.SkuStockListRequest
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	resp, err := l.svcCtx.ProductRpc.UpdateSku(l.ctx, &rpcReq)

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
