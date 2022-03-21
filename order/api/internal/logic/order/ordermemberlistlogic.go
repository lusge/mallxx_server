package order

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
	"mallxx_server/order/rpc/orderservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderMemberListLogic {
	return OrderMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderMemberListLogic) OrderMemberList(req types.OrderMemberRequest) (*types.OrderListResponse, error) {
	memberId := l.ctx.Value("uid").(int64)
	if memberId <= 0 {
		return nil, merrorx.NewCodeError(401, "参数错误")
	}
	resp, err := l.svcCtx.OrcerRpc.OrderMemberList(l.ctx, &orderservice.OrderMemberRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		MemberId: memberId,
		Type:     req.Type,
	})

	if err != nil {
		return nil, err
	}

	var data []*types.Order
	copier.Copy(&data, resp.Data)

	return &types.OrderListResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
