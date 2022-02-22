package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeCategoryShowStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeCategoryShowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeCategoryShowStatusLogic {
	return &ChangeCategoryShowStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeCategoryShowStatusLogic) ChangeCategoryShowStatus(in *pb.CategoryChangeStatus) (*pb.Response, error) {
	if l.svcCtx.CategoryModel.ChangeCategoryShowStatus(in.Id, in.Status) {
		return &pb.Response{
			Code:   200,
			Detail: "更改成功",
		}, nil
	}

	return &pb.Response{
		Code:   500,
		Detail: "更新失败",
	}, nil
}
