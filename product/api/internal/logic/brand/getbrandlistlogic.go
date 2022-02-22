package brand

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBrandListLogic {
	return GetBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBrandListLogic) GetBrandList(req types.BrandRequest) (*types.BrandListResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetBrandList(l.ctx, &productservices.BrandRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Id:       req.Id,
	})

	if err != nil {
		return nil, err
	}

	var data []*types.Brand
	copier.Copy(&data, resp.Data)
	return &types.BrandListResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
