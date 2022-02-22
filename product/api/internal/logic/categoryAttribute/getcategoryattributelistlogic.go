package categoryAttribute

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryAttributeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryAttributeListLogic {
	return GetCategoryAttributeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryAttributeListLogic) GetCategoryAttributeList(req types.ListRequest) (*types.CategoryAttributeListResponse, error) {
	resp, err := l.svcCtx.ProductRpc.GetCategoryAttributeList(l.ctx, &productservices.ListRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Pid:      req.Pid,
	})

	if err != nil {
		return nil, err
	}

	var data []*types.CategoryAttribute
	copier.Copy(&data, resp.Data)

	return &types.CategoryAttributeListResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
