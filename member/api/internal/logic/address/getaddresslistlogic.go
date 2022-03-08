package address

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"
	"mallxx_server/member/rpc/memberservice"

	"github.com/jinzhu/copier"
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

func (l *GetAddressListLogic) GetAddressList(req types.MemberRequest) (*types.ReceiveAddressListResponse, error) {
	resp, err := l.svcCtx.MemberRpc.GetAddress(l.ctx, &memberservice.MemberRequest{
		Uid: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	var data []*types.ReceiveAddress
	copier.Copy(&data, resp.Data)

	return &types.ReceiveAddressListResponse{
		Code:   resp.Code,
		Data:   data,
		Detail: resp.Detail,
	}, nil
}
