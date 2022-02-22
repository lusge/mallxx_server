package logic

import (
	"context"

	"mallxx_server/admin/api/internal/svc"
	"mallxx_server/admin/api/internal/types"
	"mallxx_server/admin/rpc/adminservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetInfoLogic {
	return GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo() (resp *types.AdminInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("uid").(int64)
	rpcResp, err := l.svcCtx.AdminRpc.GetInfo(l.ctx, &adminservice.AdminInfo{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	adminInfo := new(types.AdminInfo)

	copier.Copy(adminInfo, rpcResp.Data)

	return &types.AdminInfoResponse{
		Code:   rpcResp.Code,
		Detail: rpcResp.Detail,
		Data:   adminInfo,
	}, nil
}
