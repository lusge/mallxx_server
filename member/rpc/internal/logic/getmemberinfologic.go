package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberInfoLogic {
	return &GetMemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberInfoLogic) GetMemberInfo(in *pb.MemberRequest) (*pb.MemberInfoResponse, error) {
	info := l.svcCtx.MemberModel.FindById(in.Uid)

	if info == nil {
		return nil, merrorx.NewCodeError(500, "用户不存在")
	}
	data := &pb.MemberData{
		Member:      info,
		MemberLevel: l.svcCtx.MemberLevelModel.FindOneById(info.MemberLevel),
	}
	return &pb.MemberInfoResponse{
		Code:   200,
		Detail: "ok",
		Data:   data,
	}, nil
}
