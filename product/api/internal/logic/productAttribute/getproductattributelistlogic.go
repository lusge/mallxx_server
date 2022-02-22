package productAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductAttributeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductAttributeListLogic {
	return GetProductAttributeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductAttributeListLogic) GetProductAttributeList(req types.AttrRequest) (*types.ProductAttributeListResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetProductAttributeList(l.ctx, &productservices.AttrRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Cid:      req.Cid,
		Type:     req.Type,
	})

	if err != nil {
		return nil, err
	}

	var data []*types.ProductAttribute
	copier.Copy(&data, resp.Data)

	return &types.ProductAttributeListResponse{
		Code:   200,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
