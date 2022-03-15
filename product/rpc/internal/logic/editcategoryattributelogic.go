package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditCategoryAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditCategoryAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditCategoryAttributeLogic {
	return &EditCategoryAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditCategoryAttributeLogic) EditCategoryAttribute(in *pb.CategoryAttribute) (*pb.ProductResponse, error) {
	if !l.svcCtx.CategoryAttributeModel.Insert(in) {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}
	return &pb.ProductResponse{
		Code:   200,
		Detail: "新增成功",
	}, nil
}
