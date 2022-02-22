package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkuLogic {
	return &GetSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkuLogic) GetSku(in *pb.SkuStockRequest) (*pb.SkuStockListResponse, error) {
	list := l.svcCtx.SkuStockModel.FindAll(in.ProductId, in.Keyword)

	return &pb.SkuStockListResponse{
		Code:   200,
		Data:   list,
		Detail: "ok",
	}, nil
}
