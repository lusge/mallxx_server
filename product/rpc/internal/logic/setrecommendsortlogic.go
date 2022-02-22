package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRecommendSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRecommendSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRecommendSortLogic {
	return &SetRecommendSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRecommendSortLogic) SetRecommendSort(in *pb.RecommendSetSortRequest) (*pb.Response, error) {
	if l.svcCtx.ProductRecommendModel.SetSort(in.Id, int(in.Sort)) == false {
		return nil, merrorx.NewCodeError(500, "设置失败")
	}

	return &pb.Response{
		Code:   200,
		Detail: "设置成功",
	}, nil
}
