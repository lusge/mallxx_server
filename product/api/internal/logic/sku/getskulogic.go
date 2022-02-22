package sku

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

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

func (l *GetSkuLogic) GetSku(req types.SkuStockRequest) (resp *types.SkuStockListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
