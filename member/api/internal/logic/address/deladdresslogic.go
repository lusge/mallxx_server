package address

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"
	"mallxx_server/member/rpc/memberservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelAddressLogic {
	return DelAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAddressLogic) DelAddress(req types.ReceiveAddressRequest) (*types.Response, error) {
	userId := l.ctx.Value("uid").(int64)
	if userId <= 0 {
		return nil, merrorx.NewCodeError(500, "参数错误")
	}

	resp, err := l.svcCtx.MemberRpc.DelAddress(l.ctx, &memberservice.ReceiveAddressRequest{
		Id:       req.Id,
		MemberId: userId,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
