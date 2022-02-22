package logic

import (
	"context"

	"mallxx_server/admin/api/internal/svc"
	"mallxx_server/admin/api/internal/types"
	"mallxx_server/admin/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.AdmminLoginReqeust) (*types.AdminInfoResponse, error) {
	// todo: add your logic here and delete this line
	adminInfoResp, err := l.svcCtx.AdminRpc.Login(l.ctx, &pb.AdminInfo{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	adminInfo := new(types.AdminInfo)
	copier.Copy(adminInfo, adminInfoResp.Data)

	return &types.AdminInfoResponse{
		Data:   adminInfo,
		Code:   adminInfoResp.Code,
		Detail: adminInfoResp.Detail,
	}, nil
}
