package models

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type ProductCategoryAttributeRelation struct {
	engine *xorm.Engine
	Table  string
}

type ProductCategoryAttributeRelationTable struct {
	Id                 int64
	ProductCategoryId  int64 //category_id
	ProductAttributeId int64
}

func NewProductCategoryAttributeRelation(engine *xorm.Engine) *ProductCategoryAttributeRelation {
	return &ProductCategoryAttributeRelation{
		Table:  "product_category_attribute_relation",
		engine: engine,
	}
}

//
func (m *ProductCategoryAttributeRelation) FindByCategoryId(categoryId int64) []ProductCategoryAttributeRelationTable {
	list := make([]ProductCategoryAttributeRelationTable, 0)

	err := m.engine.Table(m.Table).Where("product_category_id = ?", categoryId).Find(&list)

	if err != nil {
		return nil
	}
	return list
}

//find by cateogry return ProductAttributeId ids
func (m *ProductCategoryAttributeRelation) FindByCategoryIdReturnIds(categoryId int64) []int64 {
	ids := make([]int64, 0)
	list := m.FindByCategoryId(categoryId)

	if list == nil {
		return ids
	}

	for _, value := range list {
		ids = append(ids, value.ProductAttributeId)
	}

	return ids
}

func (m *ProductCategoryAttributeRelation) InsertListBySession(session *xorm.Session, categoryId int64, list []int64) bool {
	productCategoryAttributeRelationSql := fmt.Sprintf("INSERT INTO product_category_attribute_relation(`product_category_id`, `product_attribute_id`) VALUES(?, ?)")

	for _, productAttrId := range list {
		_, err := session.Exec(productCategoryAttributeRelationSql, categoryId, productAttrId)
		if err != nil {
			logx.Error(err)
			return false
		}
	}

	return true
}

func (m *ProductCategoryAttributeRelation) DeleteByCategoryIdSession(session *xorm.Session, cateogryId int64) bool {
	_, err := session.Table(m.Table).Where("product_category_id = ?", cateogryId).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}
