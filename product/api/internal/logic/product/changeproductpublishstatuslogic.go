package product

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeProductPublishStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeProductPublishStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeProductPublishStatusLogic {
	return ChangeProductPublishStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeProductPublishStatusLogic) ChangeProductPublishStatus(req types.ProductChangeStatusRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.ChangeProductPublishStatus(l.ctx, &productservices.ProductChangeStatusRequest{
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
