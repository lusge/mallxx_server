package order

import (
	"context"

	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
	"mallxx_server/order/rpc/orderservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCancelLogic {
	return OrderCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCancelLogic) OrderCancel(req types.OrderInfoRequest) (*types.OrderResponse, error) {
	resp, err := l.svcCtx.OrcerRpc.OrderCancel(l.ctx, &orderservice.OrderInfoRequest{
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
