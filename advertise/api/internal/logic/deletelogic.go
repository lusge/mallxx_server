package logic

import (
	"context"

	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"
	"mallxx_server/advertise/rpc/advertiseservice"

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

func (l *DeleteLogic) Delete(req types.AdOneRequest) (resp *types.Response, err error) {
	rpcResp, err := l.svcCtx.AdvertiseRpc.Delete(l.ctx, &advertiseservice.AdOneRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   rpcResp.Code,
		Detail: rpcResp.Detail,
	}, nil
}
