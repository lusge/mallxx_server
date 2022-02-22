package member

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMemberListLogic {
	return GetMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberListLogic) GetMemberList(req types.MemberRequest) (resp *types.MemberResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
