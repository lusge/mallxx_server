package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryInfoLogic {
	return &GetCategoryInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryInfoLogic) GetCategoryInfo(in *pb.Category) (*pb.CategoryInfoResponse, error) {
	category := l.svcCtx.CategoryModel.FindById(in.Id)

	if category == nil {
		return nil, merrorx.NewCodeError(500, "分类不存在")
	}

	ids := l.svcCtx.ProductCategoryAttributeRelationModel.FindByCategoryIdReturnIds(category.Id)

	category.ProductAttributeIds = ids

	return &pb.CategoryInfoResponse{
		Code:   200,
		Detail: "ok",
		Data:   category,
	}, nil
}
