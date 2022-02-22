package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryAttributeListLogic {
	return &GetCategoryAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CategoryAttribute
func (l *GetCategoryAttributeListLogic) GetCategoryAttributeList(in *pb.ListRequest) (*pb.CategoryAttributeListResponse, error) {
	start, size := utils.Pager(in.PageNum, in.PageSize)

	list := l.svcCtx.CategoryAttributeModel.FindByLimit(start, size)
	total := l.svcCtx.CategoryAttributeModel.Count()

	return &pb.CategoryAttributeListResponse{
		Code:   200,
		Detail: "ok",
		Data:   list,
		Total:  total,
	}, nil
}
