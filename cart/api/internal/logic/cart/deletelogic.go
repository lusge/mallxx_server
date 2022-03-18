package cart

import (
	"context"

	"mallxx_server/cart/api/internal/svc"
	"mallxx_server/cart/api/internal/types"
	"mallxx_server/cart/rpc/cartservice"
	"mallxx_server/common/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteLogic {
	return DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req types.CartDelRequest) (*types.Response, error) {
	memberId := metadata.GetUidByCtx(l.ctx)
	resp, err := l.svcCtx.CartRpc.DelCart(l.ctx, &cartservice.CartDelRequest{
		Ids:      req.Ids,
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
