package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeProductRecommendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeProductRecommendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeProductRecommendStatusLogic {
	return &ChangeProductRecommendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeProductRecommendStatusLogic) ChangeProductRecommendStatus(in *pb.ProductChangeStatusRequest) (*pb.ProductResponse, error) {
	if l.svcCtx.ProductModel.ChangeStatus(in.Id, map[string]interface{}{
		"recommend_status": in.Status,
	}) == false {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}
	return &pb.ProductResponse{
		Code:   200,
		Detail: "更新成功",
	}, nil
}
