package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Category
func (l *GetCategoryListLogic) GetCategoryList(in *pb.ListRequest) (*pb.CategoryResponse, error) {
	pid := in.Pid

	start, size := utils.Pager(in.PageNum, in.PageSize)
	//结果查询
	list := l.svcCtx.CategoryModel.FindByLimit(start, size, pid)
	count := l.svcCtx.CategoryModel.Count(pid)

	return &pb.CategoryResponse{
		Code:   200,
		Detail: "ok",
		Data:   list,
		Total:  count,
	}, nil
}
