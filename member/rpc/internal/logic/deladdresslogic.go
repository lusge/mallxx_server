package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAddressLogic {
	return &DelAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAddressLogic) DelAddress(in *pb.ReceiveAddressRequest) (*pb.MemberResponse, error) {
	if !l.svcCtx.ReceiveAddressModel.Delete(in.MemberId, in.Id) {
		return nil, merrorx.NewCodeError(500, "删除失败")
	}
	return &pb.MemberResponse{
		Code:   200,
		Detail: "删除成功",
	}, nil
}
