package logic

import (
	"context"

	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"
	"mallxx_server/advertise/rpc/advertiseservice"

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

func (l *UpdateLogic) Update(req types.Advertise) (resp *types.Response, err error) {
	var rpcReq advertiseservice.Advertise
	copier.Copy(&rpcReq, req)
	response, err := l.svcCtx.AdvertiseRpc.Update(l.ctx, &rpcReq)

	return &types.Response{
		Code:   response.Code,
		Detail: response.Detail,
	}, nil
}
