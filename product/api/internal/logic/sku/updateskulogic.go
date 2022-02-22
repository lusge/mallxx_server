package sku

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

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

func (l *UpdateSkuLogic) UpdateSku(req types.SkuStockListRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
