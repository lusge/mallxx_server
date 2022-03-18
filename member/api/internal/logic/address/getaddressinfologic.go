package address

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/common/metadata"
	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"
	"mallxx_server/member/rpc/memberservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAddressInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAddressInfoLogic {
	return GetAddressInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAddressInfoLogic) GetAddressInfo(req types.ReceiveAddressRequest) (*types.ReceiveAddressResponse, error) {
	memberId := metadata.GetUidByCtx(l.ctx)
	if memberId <= 0 {
		return nil, merrorx.NewCodeError(401, "请登录")
	}

	resp, err := l.svcCtx.MemberRpc.GetAddressInfo(l.ctx, &memberservice.ReceiveAddressRequest{
		Id:       req.Id,
		MemberId: memberId,
	})

	if err != nil {
		return nil, err
	}

	var data types.ReceiveAddress
	copier.Copy(&data, resp.Data)
	return &types.ReceiveAddressResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
