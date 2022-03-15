package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *pb.Category) (*pb.ProductResponse, error) {
	if in.Name == "" {
		return nil, merrorx.NewCodeError(500, "名称不能为空")
	}
	lastId, isOk := l.svcCtx.CategoryModel.Insert(in)

	if !isOk || lastId <= 0 {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}

	return &pb.ProductResponse{
		Code:   200,
		Detail: "ok",
	}, nil
}
