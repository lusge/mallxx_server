package sign

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.MemberLoginRequest) (resp *types.MemberResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
