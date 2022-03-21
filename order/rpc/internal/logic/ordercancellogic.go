package logic

import (
	"context"

	"mallxx_server/order/rpc/internal/svc"
	"mallxx_server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCancelLogic {
	return &OrderCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderCancelLogic) OrderCancel(in *pb.OrderInfoRequest) (*pb.OrderResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.OrderResponse{}, nil
}
