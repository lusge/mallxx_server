package address

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelAddressLogic {
	return DelAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAddressLogic) DelAddress(req types.ReceiveAddressRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
