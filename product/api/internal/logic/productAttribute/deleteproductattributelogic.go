package productAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductAttributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteProductAttributeLogic {
	return DeleteProductAttributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductAttributeLogic) DeleteProductAttribute(req types.IdRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.DeleteProductAttribute(l.ctx, &productservices.ProductAttribute{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   200,
		Detail: resp.Detail,
	}, nil
}
