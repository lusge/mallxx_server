package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkuLogic {
	return &UpdateSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// sku
func (l *UpdateSkuLogic) UpdateSku(in *pb.SkuStockListRequest) (*pb.Response, error) {
	isOk := l.svcCtx.SkuStockModel.UpdateList(in.List)
	if !isOk {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}

	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
