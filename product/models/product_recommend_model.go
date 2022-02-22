package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type ProductRecommend struct {
	engine *xorm.Engine
	table  string
}

func NewProductRecommend(engine *xorm.Engine) *ProductRecommend {
	return &ProductRecommend{
		table:  "product_recommend",
		engine: engine,
	}
}

func (m *ProductRecommend) SetSort(id int64, sort int) bool {
	_, err := m.engine.Table(m.table).Where("id = ?", id).Update(map[string]interface{}{
		"sort": sort,
	})

	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (m *ProductRecommend) Insert(ids []int64) bool {
	productModel := NewProduct(m.engine)
	for _, id := range ids {
		product := productModel.FindOne(id)
		if product != nil {
			m.engine.Table(m.table).Insert(&pb.ProductRecommend{
				ProductId:   product.Id,
				ProductName: product.Name,
				Sort:        0,
			})
		}
	}
	return true
}

func (m *ProductRecommend) Delete(ids []int64) bool {

	for _, id := range ids {
		m.engine.Table(m.table).Where("id = ?", id).Delete()
	}
	return true
}

//return product ids
func (m *ProductRecommend) FindPageIds(size, start int32) []int64 {
	ids := make([]int64, 0)
	list := make([]pb.ProductRecommend, 0)
	m.engine.Limit(int(size), int(start)).Find(&list)

	for _, value := range list {
		ids = append(ids, value.ProductId)
	}
	return ids
}

func (m *ProductRecommend) FindPage(size, start int32) []*pb.ProductRecommend {
	list := make([]*pb.ProductRecommend, 0)
	m.engine.Limit(int(size), int(start)).Find(&list)

	return list
}

func (m *ProductRecommend) Count() int64 {
	total, err := m.engine.Table(m.table).Count()
	if err != nil {
		logx.Error(err)
	}

	return total
}
