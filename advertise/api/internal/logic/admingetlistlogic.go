package logic

import (
	"context"

	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"
	"mallxx_server/advertise/rpc/advertiseservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminGetListLogic {
	return AdminGetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetListLogic) AdminGetList(req types.AdRequest) (resp *types.AdResponse, err error) {
	var rpcRequest advertiseservice.AdRequest
	copier.Copy(&rpcRequest, req)
	response, err := l.svcCtx.AdvertiseRpc.AdminGetList(l.ctx, &rpcRequest)
	if err != nil {
		return nil, err
	}

	data := make([]*types.Advertise, 0)
	copier.Copy(&data, response.Data)

	return &types.AdResponse{
		Code:   response.Code,
		Detail: response.Detail,
		Data:   data,
		Total:  response.Total,
	}, nil
}
