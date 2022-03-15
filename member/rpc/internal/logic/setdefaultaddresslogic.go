package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDefaultAddressLogic {
	return &SetDefaultAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetDefaultAddressLogic) SetDefaultAddress(in *pb.ReceiveAddressRequest) (*pb.MemberResponse, error) {
	if !l.svcCtx.ReceiveAddressModel.SetDefault(in.MemberId, in.Id) {
		return nil, merrorx.NewCodeError(500, "设置成功")
	}
	return &pb.MemberResponse{
		Code:   200,
		Detail: "设置失败",
	}, nil
}
