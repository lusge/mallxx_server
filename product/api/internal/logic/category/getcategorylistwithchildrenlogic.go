package category

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListWithChildrenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListWithChildrenLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryListWithChildrenLogic {
	return GetCategoryListWithChildrenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListWithChildrenLogic) GetCategoryListWithChildren() (*types.CategoryResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetCategoryListWithChildren(l.ctx, &productservices.EmptyRequest{})
	if err != nil {
		return nil, merrorx.NewCodeError(500, "")
	}

	var data []*types.Category
	copier.Copy(&data, resp.Data)

	return &types.CategoryResponse{
		Code:   200,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
