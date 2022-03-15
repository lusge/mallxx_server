package address

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"
	"mallxx_server/member/rpc/memberservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddAddressLogic {
	return AddAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddressLogic) AddAddress(req types.ReceiveAddress) (*types.ReceiveAddressResponse, error) {

	userId := l.ctx.Value("uid").(int64)
	if userId <= 0 {
		return nil, merrorx.NewCodeError(500, "参数错误")
	}

	var rpcReq memberservice.ReceiveAddress
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, "参数错误")
	}

	rpcReq.MemberId = userId

	resp, err := l.svcCtx.MemberRpc.AddAddress(l.ctx, &rpcReq)

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
