package models

import (
	"fmt"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	skuStockCacheUsernamePrefix = "mallxxadmin:skuStock:%s"
)

type SkuStock struct {
	engine *xorm.Engine
	table  string
}

func NewSkuStock(engine *xorm.Engine) *SkuStock {
	return &SkuStock{
		table:  "sku_stock",
		engine: engine,
	}
}

func (m *SkuStock) FindByProductId(productId int64) *pb.SkuStock {
	skuStock := new(pb.SkuStock)

	has, err := m.engine.Where("product_id = ?", productId).Get(skuStock)
	if err != nil {
		logx.Error(err)
		return nil
	}

	if has == false {
		return nil
	}

	return skuStock
}

func (m *SkuStock) FindAllByProductId(productId int64) []*pb.SkuStock {
	list := make([]*pb.SkuStock, 0)
	// fmt.Println(sql)
	err := m.engine.Where("product_id = ?", productId).Find(&list)
	if err != nil {
		fmt.Println(err)
		logx.Error(err)
		return nil
	}

	return list
}

func (m *SkuStock) DeleteByProductIdAndSession(session *xorm.Session, productId int64) bool {
	affected, err := session.Table(m.table).Where("product_id = ?", productId).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	return true
}

func (m *SkuStock) InsertListBySession(session *xorm.Session, productId int64, in []*pb.SkuStock) bool {

	for _, v := range in {
		v.ProductId = productId
		_, err := session.Table(m.table).InsertOne(v)
		if err != nil {
			logx.Error(err)
			return false
		}

	}
	return true
}

func (m *SkuStock) FindAll(productId int64, skuCode string) []*pb.SkuStock {
	list := make([]*pb.SkuStock, 0)
	db := m.engine.Where("product_id = ?", productId)
	if len(skuCode) > 0 {
		db.And("sku_code LIKE ?", "%"+skuCode+"%")
	}
	// fmt.Println(sql)
	err := db.Find(&list)
	if err != nil {
		logx.Error(err)
	}

	return list
}

func (m *SkuStock) UpdateList(list []*pb.SkuStock) bool {
	for _, value := range list {
		m.engine.ShowSQL(true)
		_, err := m.engine.Where("id = ?", value.Id).Update(value)
		if err != nil {
			logx.Error(err)
			return false
		}
	}

	return true
}
