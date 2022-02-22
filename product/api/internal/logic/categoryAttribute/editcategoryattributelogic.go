package categoryAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditCategoryAttributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditCategoryAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditCategoryAttributeLogic {
	return EditCategoryAttributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditCategoryAttributeLogic) EditCategoryAttribute(req types.CategoryAttribute) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.EditCategoryAttribute(l.ctx, &productservices.CategoryAttribute{
		Id:             req.Id,
		Name:           req.Name,
		AttributeCount: req.AttributeCount,
		ParamCount:     req.ParamCount,
	})
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
