package logic

import (
	"context"

	"mallxx_server/order/rpc/internal/svc"
	"mallxx_server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDeleteLogic) OrderDelete(in *pb.OrderInfoRequest) (*pb.OrderResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.OrderResponse{}, nil
}
