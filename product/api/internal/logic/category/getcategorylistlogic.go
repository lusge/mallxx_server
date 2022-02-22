package category

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryListLogic {
	return GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req types.ListRequest) (*types.CategoryResponse, error) {
	rpcReq := new(productservices.ListRequest)
	copier.Copy(rpcReq, &req)
	resp, err := l.svcCtx.ProductRpc.GetCategoryList(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	var data []*types.Category
	copier.Copy(&data, resp.Data)
	return &types.CategoryResponse{
		Code:   resp.Code,
		Data:   data,
		Detail: resp.Detail,
	}, nil
}
