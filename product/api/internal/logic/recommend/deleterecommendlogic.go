package recommend

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteRecommendLogic {
	return DeleteRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRecommendLogic) DeleteRecommend(req types.RecommendRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.DeleteRecommend(l.ctx, &productservices.RecommendRequest{
		ProductIds: req.ProductIds,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
