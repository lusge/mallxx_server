package order

import (
	"context"

	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
	"mallxx_server/order/rpc/orderservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderInfoLogic {
	return OrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderInfoLogic) OrderInfo(req types.OrderInfoRequest) (*types.OrderInfoRsponse, error) {
	resp, err := l.svcCtx.OrcerRpc.OrderInfo(l.ctx, &orderservice.OrderInfoRequest{
		OrderId: req.OrderId,
	})

	if err != nil {
		return nil, err
	}

	var data types.Order
	copier.Copy(&data, resp.Data)
	return &types.OrderInfoRsponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
