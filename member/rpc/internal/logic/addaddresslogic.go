package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAddressLogic) AddAddress(in *pb.ReceiveAddress) (*pb.ReceiveAddressResponse, error) {
	if l.svcCtx.ReceiveAddressModel.Insert(in) == false {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}

	return &pb.ReceiveAddressResponse{
		Code:   200,
		Detail: "ok",
		Data:   in,
	}, nil
}
