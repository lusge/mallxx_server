package product

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeProductVerifyStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeProductVerifyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeProductVerifyStatusLogic {
	return ChangeProductVerifyStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeProductVerifyStatusLogic) ChangeProductVerifyStatus(req types.ProductChangeStatusRequest) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.ChangeProductVerifyStatus(l.ctx, &productservices.ProductChangeStatusRequest{
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
