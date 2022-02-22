package brand

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBrandShowStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBrandShowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateBrandShowStatusLogic {
	return UpdateBrandShowStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBrandShowStatusLogic) UpdateBrandShowStatus(req types.BrandStatusRequest) (*types.Response, error) {

	resp, err := l.svcCtx.ProductRpc.UpdateBrandShowStatus(l.ctx, &productservices.BrandStatusRequest{
		Id:     req.Id,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
