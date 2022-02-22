package category

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryInfoLogic {
	return GetCategoryInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryInfoLogic) GetCategoryInfo(req types.IdRequest) (*types.CategoryInfoResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetCategoryInfo(l.ctx, &productservices.Category{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var info types.Category
	copier.Copy(&info, resp.Data)

	return &types.CategoryInfoResponse{
		Data:   &info,
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
