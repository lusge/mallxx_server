package address

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetDefaultLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetDefaultLogic {
	return SetDefaultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetDefaultLogic) SetDefault(req types.ReceiveAddressRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
