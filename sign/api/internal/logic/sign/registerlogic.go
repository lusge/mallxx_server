package sign

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/sign/api/internal/svc"
	"mallxx_server/sign/api/internal/types"
	"mallxx_server/sign/rpc/signservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.Member) (*types.MemberResponse, error) {
	var rpcReq signservice.Member

	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}
	resp, err := l.svcCtx.SignRpc.Register(l.ctx, &rpcReq)

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
