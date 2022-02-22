package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Product
func (l *GetProductListLogic) GetProductList(in *pb.ProductListRequest) (*pb.ProductListRespone, error) {
	productModel := l.svcCtx.ProductModel

	start, size := utils.Pager(in.PageNum, in.PageSize)

	list := productModel.FindAll(in, start, size)
	total := productModel.Count(in)
	return &pb.ProductListRespone{
		Code:   200,
		Detail: "ok",
		Data:   list,
		Total:  total,
	}, nil
}
