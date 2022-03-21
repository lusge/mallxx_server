package logic

import (
	"context"

	"mallxx_server/order/rpc/internal/svc"
	"mallxx_server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderInfoLogic) OrderInfo(in *pb.OrderInfoRequest) (*pb.OrderInfoRsponse, error) {
	// todo: add your logic here and delete this line

	return &pb.OrderInfoRsponse{}, nil
}
