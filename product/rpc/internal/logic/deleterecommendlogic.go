package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRecommendLogic {
	return &DeleteRecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRecommendLogic) DeleteRecommend(in *pb.RecommendRequest) (*pb.Response, error) {
	if l.svcCtx.ProductRecommendModel.Delete(in.ProductIds) == false {
		return nil, merrorx.NewCodeError(500, "删除失败")
	}

	return &pb.Response{
		Code:   200,
		Detail: "删除成功",
	}, nil
}
