package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryAttributeInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryAttributeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryAttributeInfoLogic {
	return &GetCategoryAttributeInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryAttributeInfoLogic) GetCategoryAttributeInfo(in *pb.CategoryAttribute) (*pb.CategoryAttributeInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CategoryAttributeInfoResponse{}, nil
}
