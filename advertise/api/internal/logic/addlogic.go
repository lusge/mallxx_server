package logic

import (
	"context"

	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"
	"mallxx_server/advertise/rpc/advertiseservice"

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

func (l *AddLogic) Add(req types.Advertise) (resp *types.Response, err error) {
	var rpcReq advertiseservice.Advertise
	copier.Copy(&rpcReq, req)
	response, err := l.svcCtx.AdvertiseRpc.Add(l.ctx, &rpcReq)

	return &types.Response{
		Code:   response.Code,
		Detail: response.Detail,
	}, nil
}
