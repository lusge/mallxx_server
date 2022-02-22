package brand

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteBrandLogic {
	return DeleteBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBrandLogic) DeleteBrand(req types.BrandRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.DeleteBrand(l.ctx, &productservices.BrandRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Id:       req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
