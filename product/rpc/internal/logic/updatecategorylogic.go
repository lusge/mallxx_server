package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *pb.Category) (*pb.ProductResponse, error) {
	isOk, err := l.svcCtx.CategoryModel.Update(in)
	if !isOk {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}

	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	return &pb.ProductResponse{
		Code:   200,
		Detail: "ok",
	}, nil
}
