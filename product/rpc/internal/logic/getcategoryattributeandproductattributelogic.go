package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryAttributeAndProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryAttributeAndProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryAttributeAndProductAttributeLogic {
	return &GetCategoryAttributeAndProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryAttributeAndProductAttributeLogic) GetCategoryAttributeAndProductAttribute(in *pb.ProductEmptyRequest) (*pb.CategoryAttrAndProductAttrListResponse, error) {
	categoryAttrList := l.svcCtx.CategoryAttributeModel.FindAll()

	data := make([]*pb.CategoryAttrAndProductAttr, 0)
	for _, value := range categoryAttrList {
		productAttrList := l.svcCtx.ProductAttributeModel.FindByCategoryAttributeId(value.Id, 1)
		data = append(data, &pb.CategoryAttrAndProductAttr{
			CategoryAttr: value,
			Children:     productAttrList,
		})
	}

	return &pb.CategoryAttrAndProductAttrListResponse{
		Code:   200,
		Data:   data,
		Detail: "ok",
	}, nil
}
