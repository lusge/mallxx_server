package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type ProductLadder struct {
	engine *xorm.Engine
	table  string
}

func NewProductLadder(engine *xorm.Engine) *ProductLadder {
	return &ProductLadder{
		table:  "product_ladder",
		engine: engine,
	}
}

func (m *ProductLadder) FindAllByProductId(productId int64) []*pb.ProductLadder {
	list := make([]*pb.ProductLadder, 0)
	err := m.engine.Table(m.table).Where("product_id = ?", productId).Find(&list)

	if err != nil {
		logx.Error(err)
	}
	return list
}

func (m *ProductLadder) DeleteByProductIdAndSession(session *xorm.Session, productId int64) bool {
	_, err := session.Table(m.table).Where("product_id = ?", productId).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}

func (m *ProductLadder) InsertListBySession(session *xorm.Session, productId int64, list []*pb.ProductLadder) bool {

	for _, v := range list {
		v.ProductId = productId
		affected, err := session.Table(m.table).InsertOne(v)
		if err != nil {
			logx.Error(err)
			return false
		}

		if affected <= 0 {
			logx.Error("Not a single line was Insert")
			return false
		}
	}

	return true
}
