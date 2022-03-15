package cart

import (
	"context"

	"mallxx_server/cart/api/internal/svc"
	"mallxx_server/cart/api/internal/types"
	"mallxx_server/cart/rpc/cartservice"
	"mallxx_server/common/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) ModifyLogic {
	return ModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyLogic) Modify(req types.CartModiyRequest) (*types.Response, error) {
	memberId := metadata.GetUidByCtx(l.ctx)
	resp, err := l.svcCtx.CartRpc.ModiryCart(l.ctx, &cartservice.CartModiyRequest{
		Id:       req.Id,
		Quantity: req.Quantity,
		MemberId: memberId,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
