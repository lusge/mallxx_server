package logic

import (
	"context"

	"mallxx_server/order/rpc/internal/svc"
	"mallxx_server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderFrimLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderFrimLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderFrimLogic {
	return &OrderFrimLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderFrimLogic) OrderFrim(in *pb.OrderFirmRequest) (*pb.OrderFirmResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.OrderFirmResponse{}, nil
}
