package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandListLogic {
	return &GetBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Brand
func (l *GetBrandListLogic) GetBrandList(in *pb.BrandRequest) (*pb.BrandListResponse, error) {
	start, size := utils.Pager(in.PageNum, in.PageSize)
	count := l.svcCtx.BrandModel.Count(in.Keyword)
	list := l.svcCtx.BrandModel.FindAll(in.Keyword, start, size)

	return &pb.BrandListResponse{
		Code:   200,
		Data:   list,
		Detail: "ok",
		Total:  count,
	}, nil
}
