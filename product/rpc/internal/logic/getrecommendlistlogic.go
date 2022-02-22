package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecommendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendListLogic {
	return &GetRecommendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecommendListLogic) GetRecommendList(in *pb.ProductListRequest) (*pb.RecommendListResponse, error) {
	start, size := utils.Pager(in.PageNum, in.PageSize)

	data := l.svcCtx.ProductRecommendModel.FindPage(size, start)

	Total := l.svcCtx.ProductRecommendModel.Count()
	return &pb.RecommendListResponse{
		Code:   200,
		Detail: "ok",
		Data:   data,
		Total:  Total,
	}, nil
}
