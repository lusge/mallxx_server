package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductAttributeInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductAttributeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductAttributeInfoLogic {
	return &GetProductAttributeInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductAttributeInfoLogic) GetProductAttributeInfo(in *pb.ProductAttribute) (*pb.ProductAttributeInfoResponse, error) {
	info := l.svcCtx.ProductAttributeModel.FindOne(in.Id)

	return &pb.ProductAttributeInfoResponse{
		Code:   200,
		Detail: "ok",
		Data:   info,
	}, nil
}
