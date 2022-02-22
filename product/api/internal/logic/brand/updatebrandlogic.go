package brand

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateBrandLogic {
	return UpdateBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBrandLogic) UpdateBrand(req types.Brand) (*types.Response, error) {
	var brand productservices.Brand
	err := copier.Copy(&brand, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	resp, err := l.svcCtx.ProductRpc.UpdateBrand(l.ctx, &brand)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
