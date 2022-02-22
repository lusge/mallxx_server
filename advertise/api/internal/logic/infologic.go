package logic

import (
	"context"

	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"
	"mallxx_server/advertise/rpc/advertiseservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) InfoLogic {
	return InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req types.AdOneRequest) (resp *types.AdInfoResponse, err error) {
	rpcResp, err := l.svcCtx.AdvertiseRpc.GetInfo(l.ctx, &advertiseservice.AdOneRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var info types.Advertise
	copier.Copy(&info, rpcResp.Data)

	return &types.AdInfoResponse{
		Code:   200,
		Detail: "ok",
		Data:   &info,
	}, nil
}
