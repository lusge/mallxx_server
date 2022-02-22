package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductApiListLogic {
	return &GetProductApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductApiListLogic) GetProductApiList(in *pb.ProductApiRequest) (*pb.ProductListRespone, error) {
	start, size := utils.Pager(in.PageNum, in.PageSize)
	list := l.svcCtx.ProductModel.FinaApiPage(in, size, start)
	return &pb.ProductListRespone{
		Code:   200,
		Data:   list,
		Detail: "ok",
	}, nil
}
