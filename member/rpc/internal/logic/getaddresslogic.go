package logic

import (
	"context"

	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressLogic {
	return &GetAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressLogic) GetAddress(in *pb.MemberRequest) (*pb.ReceiveAddressListResponse, error) {
	data := l.svcCtx.ReceiveAddressModel.FindAllByMemberId(in.Uid)

	return &pb.ReceiveAddressListResponse{
		Data:   data,
		Code:   200,
		Detail: "ok",
	}, nil
}
