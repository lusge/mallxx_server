package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddressLogic) UpdateAddress(in *pb.ReceiveAddress) (*pb.MemberResponse, error) {
	if !l.svcCtx.ReceiveAddressModel.Update(in) {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}
	return &pb.MemberResponse{
		Code:   200,
		Detail: "更新失败",
	}, nil
}
