package follower

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetFollowerLogic {
	return GetFollowerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowerLogic) GetFollower(req types.MemberRequest) (resp *types.FollowerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
