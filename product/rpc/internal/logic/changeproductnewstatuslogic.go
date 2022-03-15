package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeProductNewStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeProductNewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeProductNewStatusLogic {
	return &ChangeProductNewStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeProductNewStatusLogic) ChangeProductNewStatus(in *pb.ProductChangeStatusRequest) (*pb.ProductResponse, error) {
	if l.svcCtx.ProductModel.ChangeStatus(in.Id, map[string]interface{}{
		"new_status": in.Status,
	}) == false {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}
	return &pb.ProductResponse{
		Code:   200,
		Detail: "更新成功",
	}, nil
}
