package categoryAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryAttributeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryAttributeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryAttributeInfoLogic {
	return GetCategoryAttributeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryAttributeInfoLogic) GetCategoryAttributeInfo(req types.CategoryAttribute) (*types.CategoryAttributeInfoResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetCategoryAttributeInfo(l.ctx, &productservices.CategoryAttribute{
		Id:             req.Id,
		Name:           req.Name,
		AttributeCount: req.AttributeCount,
		ParamCount:     req.ParamCount,
	})

	if err != nil {
		return nil, err
	}

	var data types.CategoryAttribute
	copier.Copy(&data, resp.Data)

	return &types.CategoryAttributeInfoResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
