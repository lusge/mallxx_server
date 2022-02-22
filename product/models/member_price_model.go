package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	memberPriceCacheUsernamePrefix = "mallxxadmin:member_price:language:%s:id:%d"
)

type MemberPrice struct {
	engine *xorm.Engine
	table  string
}

func NewMemberPrice(engine *xorm.Engine) *MemberPrice {
	return &MemberPrice{
		table:  "member_price",
		engine: engine,
	}
}

func (m *MemberPrice) FindByProductId(productId int64) []*pb.MemberPrice {
	list := make([]*pb.MemberPrice, 0)
	err := m.engine.Where("product_id = ?", productId).Find(&list)

	if err != nil {
		logx.Error(err)
	}

	return list
}

func (m *MemberPrice) InsertListBySession(session *xorm.Session, productId int64, list []*pb.MemberPrice) bool {
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

func (m *MemberPrice) DeleteBySession(session *xorm.Session, productId int64) bool {
	_, err := session.Table(m.table).Where("product_id = ?", productId).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}
