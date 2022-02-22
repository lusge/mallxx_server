package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type ProductFullReduction struct {
	engine *xorm.Engine
	table  string
}

func NewProductFullReduction(engine *xorm.Engine) *ProductFullReduction {
	return &ProductFullReduction{
		table:  "product_full_reduction",
		engine: engine,
	}
}

func (m *ProductFullReduction) FindAllByProductId(productId int64) []*pb.ProductFullReduction {
	list := make([]*pb.ProductFullReduction, 0)
	err := m.engine.Table(m.table).Where("product_id = ?", productId).Find(&list)
	if err != nil {
		logx.Error(err)
	}
	return list
}

func (m *ProductFullReduction) DeleteByProductIdAndSession(session *xorm.Session, productId int64) bool {
	_, err := session.Table(m.table).Where("product_id = ?", productId).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}

func (m *ProductFullReduction) InsertListBySession(session *xorm.Session, productId int64, list []*pb.ProductFullReduction) bool {

	for _, v := range list {
		v.ProductId = productId
		affected, err := session.Table(m.table).InsertOne(v)
		if err != nil {
			logx.Error(err)
			return false
		}

		if affected <= 0 {
			logx.Error("Not a single line was deleted")
			return false
		}
	}

	return true
}
