package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecommendLogic {
	return &AddRecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRecommendLogic) AddRecommend(in *pb.RecommendRequest) (*pb.ProductResponse, error) {
	if l.svcCtx.ProductRecommendModel.Insert(in.ProductIds) == false {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}
	return &pb.ProductResponse{
		Code:   200,
		Detail: "新增成功",
	}, nil
}
