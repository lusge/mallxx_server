package models

import (
	"mallxx_server/cart/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	RedisCacheKeyPrefix = "mallxx:cart:memberid:%d"
)

type Cart struct {
	engine *xorm.Engine
	table  string
}

func NewCart(engine *xorm.Engine) *Cart {
	return &Cart{
		table:  "cart",
		engine: engine,
	}
}

func (m *Cart) Insert(in *pb.Cart) bool {
	_, err := m.engine.InsertOne(in)
	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (m *Cart) Modify(in *pb.CartModiyRequest, memberId int64) bool {
	// key := fmt.Sprintf(RedisCacheKeyPrefix, memberId)
	// fmt.Println(key, "cart model")
	// cartString, err := m.redis.Get(key)

	// if err != nil {
	// 	logx.Error(err)
	// 	return false
	// }

	// list := make([]*pb.Cart, 0)

	// if len(cartString) > 0 {
	// 	err = json.Unmarshal([]byte(cartString), &list)
	// 	if err != nil {
	// 		logx.Error(err)
	// 		return false
	// 	}

	// 	for _, value := range list {
	// 		if value.Id == in.Id {
	// 			value.Quantity += in.Quantity
	// 			m.ModifyQuantity(value.Id, value.Quantity)
	// 		}
	// 	}

	// 	js, _ := json.Marshal(list)
	// 	m.redis.Set(key, string(js))
	// }
	return true
}

func (m *Cart) ModifyQuantity(id int64, quantity int32) bool {

	_, err := m.engine.Table(m.table).Where("id = ?", id).Update(map[string]interface{}{
		"quantity":    quantity,
		"modify_date": time.Now().Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (m *Cart) Delete(id int64) bool {
	_, err := m.engine.Table(m.table).Where("id = ?", id).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (m *Cart) FindAllByMemberId(memberId int64) []*pb.Cart {
	list := make([]*pb.Cart, 0)
	err := m.engine.Where("member_id = ?", memberId).Find(&list)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}

func (m *Cart) FindOneByProductIdWithSkuId(memberId, productId, skuId int64) *pb.Cart {
	cart := new(pb.Cart)

	has, err := m.engine.Where("member_id = ?", memberId).And("product_id = ?", productId).And("sku_id = ?", skuId).Get(cart)

	if err != nil {
		logx.Error(err)
		return nil
	}
	if !has {
		return nil
	}

	return cart
}
