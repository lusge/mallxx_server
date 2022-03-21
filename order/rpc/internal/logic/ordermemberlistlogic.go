package logic

import (
	"context"

	"mallxx_server/order/rpc/internal/svc"
	"mallxx_server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderMemberListLogic {
	return &OrderMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderMemberListLogic) OrderMemberList(in *pb.OrderMemberRequest) (*pb.OrderListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.OrderListResponse{}, nil
}
