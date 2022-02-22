package product

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeProductRecommendStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeProductRecommendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeProductRecommendStatusLogic {
	return ChangeProductRecommendStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeProductRecommendStatusLogic) ChangeProductRecommendStatus(req types.ProductChangeStatusRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.ChangeProductRecommendStatus(l.ctx, &productservices.ProductChangeStatusRequest{
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
