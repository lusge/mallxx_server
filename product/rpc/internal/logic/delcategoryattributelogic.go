package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCategoryAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCategoryAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCategoryAttributeLogic {
	return &DelCategoryAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCategoryAttributeLogic) DelCategoryAttribute(in *pb.CategoryAttribute) (*pb.ProductResponse, error) {
	if !l.svcCtx.CategoryAttributeModel.Delete(in.Id) {
		return nil, merrorx.NewCodeError(500, "删除失败")
	}
	return &pb.ProductResponse{
		Code:   200,
		Detail: "删除成功",
	}, nil
}
