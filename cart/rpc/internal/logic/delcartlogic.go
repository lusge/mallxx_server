package logic

import (
	"context"

	"mallxx_server/cart/rpc/internal/svc"
	"mallxx_server/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartLogic {
	return &DelCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCartLogic) DelCart(in *pb.CartDelRequest) (*pb.CartResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CartResponse{}, nil
}
