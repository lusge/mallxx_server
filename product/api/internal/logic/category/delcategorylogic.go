package category

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelCategoryLogic {
	return DelCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCategoryLogic) DelCategory(req types.IdRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.DelCategory(l.ctx, &productservices.Category{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
