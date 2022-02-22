package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiRecommendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiRecommendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiRecommendListLogic {
	return &GetApiRecommendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  product recommend
func (l *GetApiRecommendListLogic) GetApiRecommendList(in *pb.ProductListRequest) (*pb.ProductListRespone, error) {
	start, size := utils.Pager(in.PageNum, in.PageSize)

	ids := l.svcCtx.ProductRecommendModel.FindPageIds(size, start)

	data := l.svcCtx.ProductModel.FindByIds(ids)

	Total := l.svcCtx.ProductRecommendModel.Count()

	return &pb.ProductListRespone{
		Code:   200,
		Detail: "ok",
		Data:   data,
		Total:  Total,
	}, nil
}
