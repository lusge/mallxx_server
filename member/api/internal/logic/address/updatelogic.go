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

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.ReceiveAddress) (*types.Response, error) {
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

	resp, err := l.svcCtx.MemberRpc.UpdateAddress(l.ctx, &rpcReq)

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
