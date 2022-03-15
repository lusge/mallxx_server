package cart

import (
	"context"

	"mallxx_server/cart/api/internal/svc"
	"mallxx_server/cart/api/internal/types"
	"mallxx_server/cart/rpc/cartservice"
	"mallxx_server/common/merrorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetListLogic {
	return GetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListLogic) GetList(req types.CartListRequest) (*types.CartResponse, error) {
	userId := l.ctx.Value("uid").(int64)
	if userId <= 0 {
		return nil, merrorx.NewCodeError(500, "参数错误")
	}

	resp, err := l.svcCtx.CartRpc.GetCartList(l.ctx, &cartservice.CartListRequest{
		MemberId: req.MemberId,
	})

	if err != nil {
		return nil, err
	}

	var data []*types.Cart
	copier.Copy(&data, resp.Data)
	return &types.CartResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
