package member

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"
	"mallxx_server/member/rpc/memberservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemverLevelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemverLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMemverLevelListLogic {
	return GetMemverLevelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemverLevelListLogic) GetMemverLevelList(req types.EmptyRequest) (*types.MemverLeveListResponse, error) {
	resp, err := l.svcCtx.MemberRpc.GetMemverLevelList(l.ctx, &memberservice.EmptyRequest{})

	if err != nil {
		return nil, err
	}

	var data []*types.MemberLevel

	copier.Copy(&data, resp.Data)

	return &types.MemverLeveListResponse{
		Code:   200,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
