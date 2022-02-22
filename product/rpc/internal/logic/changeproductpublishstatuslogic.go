package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeProductPublishStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeProductPublishStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeProductPublishStatusLogic {
	return &ChangeProductPublishStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeProductPublishStatusLogic) ChangeProductPublishStatus(in *pb.ProductChangeStatusRequest) (*pb.Response, error) {
	if l.svcCtx.ProductModel.ChangeStatus(in.Id, map[string]interface{}{
		"publish_status": in.Status,
	}) == false {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}
	return &pb.Response{
		Code:   200,
		Detail: "更新成功",
	}, nil
}
