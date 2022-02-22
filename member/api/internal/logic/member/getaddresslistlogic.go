package member

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAddressListLogic {
	return GetAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAddressListLogic) GetAddressList(req types.MemberRequest) (resp *types.ReceiveAddressListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
