package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *pb.ProductInfoRequest) (*pb.Response, error) {
	isOk := l.svcCtx.ProductModel.ChangeStatus(in.Id, map[string]interface{}{
		"delete_status": 1,
	})

	if !isOk {
		return nil, merrorx.NewCodeError(500, "删除失败")
	}
	return &pb.Response{
		Code:   200,
		Detail: "删除成功",
	}, nil
}
