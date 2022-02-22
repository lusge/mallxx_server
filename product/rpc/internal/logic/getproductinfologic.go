package logic

import (
	"context"
	"fmt"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductInfoLogic {
	return &GetProductInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductInfoLogic) GetProductInfo(in *pb.ProductInfoRequest) (*pb.ProductInfoResponse, error) {
	product := l.svcCtx.ProductModel.FindOne(in.Id)

	if product == nil {
		return nil, merrorx.NewCodeError(500, "该商品不存在或以下架")
	}

	product.SkuStock = l.svcCtx.SkuStockModel.FindAllByProductId(in.Id)
	product.ProductAttrValue = l.svcCtx.ProductAttributeValueModel.FindAllByProductId(in.Id)
	product.ProductLadder = l.svcCtx.ProductLadderModel.FindAllByProductId(in.Id)
	product.ProductFullReduction = l.svcCtx.ProductFullReductionModel.FindAllByProductId(in.Id)
	product.MemberPrice = l.svcCtx.MemberPriceModel.FindByProductId(in.Id)
	product.ProductAttribute = l.svcCtx.ProductAttributeModel.FindByCategoryAttributeId(product.ProductAttributeCategoryId, 0)

	parameters := make([]*pb.ProductParameters, 0)
	parameters = append(parameters, &pb.ProductParameters{Key: "品牌", Value: product.BrandName},
		&pb.ProductParameters{Key: "分类", Value: product.ProductCategoryName},
	)

	for _, value := range product.ProductAttrValue {
		productAttribute := l.svcCtx.ProductAttributeModel.FindOne(value.ProductAttributeId)
		if productAttribute != nil {
			parameters = append(parameters, &pb.ProductParameters{Key: productAttribute.Name, Value: value.Value})
		}
	}

	product.Parameters = parameters
	fmt.Println(product)
	return &pb.ProductInfoResponse{
		Code:   200,
		Detail: "ok",
		Data:   product,
	}, nil
}
