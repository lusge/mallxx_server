package product

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductInfoLogic {
	return GetProductInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductInfoLogic) GetProductInfo(req types.ProductInfoRequest) (*types.ProductInfoResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetProductInfo(l.ctx, &productservices.ProductInfoRequest{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	var data types.Product
	copier.Copy(&data, resp.Data)
	return &types.ProductInfoResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
