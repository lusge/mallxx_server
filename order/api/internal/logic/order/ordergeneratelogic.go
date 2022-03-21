package order

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
	"mallxx_server/order/rpc/orderservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderGenerateLogic {
	return OrderGenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderGenerateLogic) OrderGenerate(req types.GenerateOrderRequest) (*types.GenerateOrderResponse, error) {
	var rpcReq orderservice.GenerateOrderRequest
	err := copier.Copy(&rpcReq, req)

	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	resp, err := l.svcCtx.OrcerRpc.OrderGenerate(l.ctx, &rpcReq)

	if err != nil {
		return nil, err
	}

	var firm []*types.OrderFirm
	var item types.Order
	copier.Copy(&firm, resp.Firm)
	copier.Copy(&item, resp.Item)

	return &types.GenerateOrderResponse{
		Code:      resp.Code,
		Detail:    resp.Detail,
		AddressId: resp.AddressId,
		Firm:      firm,
		Item:      &item,
	}, nil
}
