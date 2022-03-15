package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListWithChildrenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryListWithChildrenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListWithChildrenLogic {
	return &GetCategoryListWithChildrenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryListWithChildrenLogic) GetCategoryListWithChildren(in *pb.ProductEmptyRequest) (*pb.CategoryResponse, error) {
	// todo: add your logic here and delete this line
	list := l.svcCtx.CategoryModel.FindListWithChildren()
	return &pb.CategoryResponse{
		Code:   200,
		Detail: "ok",
		Data:   list,
	}, nil
}
