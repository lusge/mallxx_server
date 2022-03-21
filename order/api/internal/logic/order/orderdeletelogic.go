package order

import (
	"context"

	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
	"mallxx_server/order/rpc/orderservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderDeleteLogic {
	return OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req types.OrderInfoRequest) (*types.OrderResponse, error) {
	resp, err := l.svcCtx.OrcerRpc.OrderDelete(l.ctx, &orderservice.OrderInfoRequest{
		OrderId: req.OrderId,
	})

	if err != nil {
		return nil, err
	}

	return &types.OrderResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
