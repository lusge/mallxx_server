package category

import (
	"context"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeCategoryShowStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeCategoryShowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeCategoryShowStatusLogic {
	return ChangeCategoryShowStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeCategoryShowStatusLogic) ChangeCategoryShowStatus(req types.CategoryChangeStatus) (*types.Response, error) {
	resp, err := l.svcCtx.ProductRpc.ChangeCategoryShowStatus(l.ctx, &productservices.CategoryChangeStatus{
		Id:     req.Id,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code:   resp.Code,
		Detail: resp.Detail,
	}, nil
}
