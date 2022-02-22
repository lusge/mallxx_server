package recommend

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRecommendSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRecommendSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetRecommendSortLogic {
	return SetRecommendSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRecommendSortLogic) SetRecommendSort(req types.RecommendSetSortRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.SetRecommendSort(l.ctx, &productservices.RecommendSetSortRequest{
		Id:   req.Id,
		Sort: req.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
