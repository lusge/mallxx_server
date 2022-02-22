package sign

import (
	"context"

	"mallxx_server/sign/api/internal/svc"
	"mallxx_server/sign/api/internal/types"
	"mallxx_server/sign/rpc/signservice"

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

func (l *LoginLogic) Login(req types.MemberLoginRequest) (*types.MemberResponse, error) {
	resp, err := l.svcCtx.SignRpc.Login(l.ctx, &signservice.MemberLoginRequest{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	})

	if err != nil {
		return nil, err
	}

	var data types.MemberData
	copier.Copy(&data, resp.Data)
	return &types.MemberResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
