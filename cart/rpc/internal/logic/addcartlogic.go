package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"mallxx_server/cart/rpc/internal/svc"
	"mallxx_server/cart/rpc/pb"
	"mallxx_server/common/merrorx"
	"mallxx_server/common/redisx"
	"mallxx_server/member/rpc/memberservice"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartLogic) AddCart(in *pb.CartAddRequest) (*pb.CartResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Cart, in.MemberId)
	var list []*pb.Cart
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if len(list) <= 0 {
		info, err := l.generateCartInfo(in)
		if err != nil {
			return nil, err
		}
		if !l.svcCtx.CartModel.Insert(info) {
			return nil, merrorx.NewCodeError(500, "添加失败")
		}

		list = append(list, info)
	} else {
		flag := true
		for _, value := range list {
			if value.SkuId == in.SkuId {
				flag = false
				value.Quantity += in.Quantity
				l.svcCtx.CartModel.ModifyQuantity(value.Id, value.Quantity)
			}
		}

		if flag {
			info, err := l.generateCartInfo(in)
			if err != nil {
				return nil, err
			}
			if !l.svcCtx.CartModel.Insert(info) {
				return nil, merrorx.NewCodeError(500, "添加失败")
			}
			list = append(list, info)
		}
	}

	js, _ := json.Marshal(list)
	err := l.svcCtx.Redis.Set(key, string(js))

	if err != nil {
		return nil, merrorx.NewCodeError(500, "redis 错误")
	}

	return &pb.CartResponse{
		Code:   200,
		Detail: "添加成功",
	}, nil
}

func (l *AddCartLogic) generateCartInfo(in *pb.CartAddRequest) (*pb.Cart, error) {
	productRet, err := l.svcCtx.ProductRpc.GetProductInfo(l.ctx, &productservices.ProductInfoRequest{
		Id: in.ProductId,
	})

	if err != nil {
		return nil, err
	}

	if productRet.Code != 200 || productRet.Data == nil {
		return nil, merrorx.NewCodeError(500, "商品不存在")
	}

	product := productRet.Data

	sku := l.getSkuStock(productRet.Data.SkuStock, in.SkuId)
	if sku == nil {
		return nil, merrorx.NewCodeError(500, "SKU 不存在")
	}

	currnetTime := time.Now().Format("2006-01-02 15:04:05")

	memberRet, err := l.svcCtx.MemberRpc.GetMemberInfo(l.ctx, &memberservice.MemberRequest{
		Uid: in.MemberId,
	})

	if memberRet.Code != 200 || memberRet.Data == nil {
		return nil, merrorx.NewCodeError(500, "用户不存在")
	}

	memberInfo := memberRet.Data.Member

	info := &pb.Cart{
		ProductId:       in.ProductId,
		Quantity:        in.Quantity,
		MemberId:        in.MemberId,
		SkuId:           in.SkuId,
		MemberNickname:  memberInfo.Nickname,
		Price:           float64(sku.Price), //sku价格
		ProductPic:      product.Pic,
		ProductName:     product.Name,
		ProductSubTitle: product.SubTitle,
		SkuCode:         sku.SkuCode,
		DeleteStatus:    0,
		CategoryId:      product.ProductCategoryId,
		BrandName:       product.BrandName,
		ProductSn:       product.ProductSn,
		ProductAttr:     sku.SpData,
		CreateDate:      currnetTime,
		ModifyDate:      currnetTime,
	}
	return info, nil
}

func (l *AddCartLogic) getSkuStock(skuList []*productservices.SkuStock, sku_id int64) *productservices.SkuStock {
	for _, sku := range skuList {
		if sku.Id == sku_id {
			return sku
		}
	}
	return nil
}
