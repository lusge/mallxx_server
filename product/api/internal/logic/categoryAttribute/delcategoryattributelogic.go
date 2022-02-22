package categoryAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCategoryAttributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCategoryAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelCategoryAttributeLogic {
	return DelCategoryAttributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCategoryAttributeLogic) DelCategoryAttribute(req types.CategoryAttribute) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.DelCategoryAttribute(l.ctx, &productservices.CategoryAttribute{
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
