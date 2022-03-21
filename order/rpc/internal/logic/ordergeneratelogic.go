package logic

import (
	"context"

	"mallxx_server/order/rpc/internal/svc"
	"mallxx_server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGenerateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderGenerateLogic {
	return &OrderGenerateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderGenerateLogic) OrderGenerate(in *pb.GenerateOrderRequest) (*pb.GenerateOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GenerateOrderResponse{}, nil
}
