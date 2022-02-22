package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCategoryLogic {
	return &DelCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCategoryLogic) DelCategory(in *pb.Category) (*pb.Response, error) {
	if l.svcCtx.CategoryModel.Delete(in.Id) {
		return &pb.Response{
			Code:   200,
			Detail: "删除成功",
		}, nil
	}
	return nil, merrorx.NewCodeError(500, "删除失败")
}
