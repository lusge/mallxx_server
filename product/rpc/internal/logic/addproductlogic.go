package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductLogic) AddProduct(in *pb.Product) (*pb.Response, error) {
	isOk := l.svcCtx.ProductModel.Insert(in)
	if !isOk {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}
	return &pb.Response{
		Code:   200,
		Detail: "新增成功",
	}, nil
}
