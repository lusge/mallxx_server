package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeCategoryNavStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeCategoryNavStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeCategoryNavStatusLogic {
	return &ChangeCategoryNavStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeCategoryNavStatusLogic) ChangeCategoryNavStatus(in *pb.CategoryChangeStatus) (*pb.ProductResponse, error) {
	if l.svcCtx.CategoryModel.ChangeCategoryNavStatus(in.Id, in.Status) {
		return &pb.ProductResponse{
			Code:   200,
			Detail: "更改成功",
		}, nil
	}

	return &pb.ProductResponse{
		Code:   500,
		Detail: "更新失败",
	}, nil
}
