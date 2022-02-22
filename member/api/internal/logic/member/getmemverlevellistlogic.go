package member

import (
	"context"

	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

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

func (l *GetMemverLevelListLogic) GetMemverLevelList(req types.EmptyRequest) (resp *types.MemverLeveListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
