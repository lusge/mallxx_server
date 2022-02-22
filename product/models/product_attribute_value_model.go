package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	productAttributeValueCacheUsernamePrefix = "mallxxadmin:productAttributeValue:%s"
)

type ProductAttributeValue struct {
	engine *xorm.Engine
	table  string
}

func NewProductAttributeValue(engine *xorm.Engine) *ProductAttributeValue {
	return &ProductAttributeValue{
		table:  "product_attribute_value",
		engine: engine,
	}
}

func (m *ProductAttributeValue) FindAllByProductId(productId int64) []*pb.ProductAttributeValue {
	list := make([]*pb.ProductAttributeValue, 0)

	err := m.engine.Where("product_id = ?", productId).Find(&list)

	// fmt.Println(sql)
	if err != nil {
		logx.Error(err)
	}
	return list
}

func (m *ProductAttributeValue) InsertListBySession(session *xorm.Session, productId int64, list []*pb.ProductAttributeValue) bool {
	for _, value := range list {
		value.ProductId = productId
		_, err := session.InsertOne(value)
		if err != nil {
			logx.Error(err)
			return false
		}
	}
	return true
}

func (m *ProductAttributeValue) DeleteBySession(session *xorm.Session, productId int64) bool {
	_, err := session.Table(m.table).Where("product_id = ?", productId).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}
