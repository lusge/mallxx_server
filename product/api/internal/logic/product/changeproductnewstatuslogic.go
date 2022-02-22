package product

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeProductNewStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeProductNewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeProductNewStatusLogic {
	return ChangeProductNewStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeProductNewStatusLogic) ChangeProductNewStatus(req types.ProductChangeStatusRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.ChangeProductNewStatus(l.ctx, &productservices.ProductChangeStatusRequest{
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
