package cart

import (
	"context"

	"mallxx_server/cart/api/internal/svc"
	"mallxx_server/cart/api/internal/types"
	"mallxx_server/cart/rpc/cartservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.CartAddRequest) (*types.Response, error) {

	var rpcReq cartservice.CartAddRequest
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.CartRpc.AddCart(l.ctx, &rpcReq)
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
