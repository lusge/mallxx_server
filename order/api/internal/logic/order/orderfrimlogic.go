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

type OrderFrimLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderFrimLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderFrimLogic {
	return OrderFrimLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderFrimLogic) OrderFrim(req types.OrderFirmRequest) (*types.OrderFirmResponse, error) {
	var rpcReq orderservice.OrderFirmRequest
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	resp, err := l.svcCtx.OrcerRpc.OrderFrim(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}

	var data []*types.OrderFirm
	copier.Copy(&data, resp.Data)
	return &types.OrderFirmResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
