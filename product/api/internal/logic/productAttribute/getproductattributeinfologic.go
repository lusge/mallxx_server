package productAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductAttributeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductAttributeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductAttributeInfoLogic {
	return GetProductAttributeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductAttributeInfoLogic) GetProductAttributeInfo(req types.IdRequest) (*types.ProductAttributeInfoResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetProductAttributeInfo(l.ctx, &productservices.ProductAttribute{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var data types.ProductAttribute
	copier.Copy(&data, resp.Data)

	return &types.ProductAttributeInfoResponse{
		Code:   200,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
