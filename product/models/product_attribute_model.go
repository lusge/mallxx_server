package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type ProductAttribute struct {
	engine *xorm.Engine
	table  string
}

var (
	productAttributeCacheUsernamePrefix = "mallxxadmin:productattribute:%s:id:"
)

func NewProductAttribute(engine *xorm.Engine) *ProductAttribute {
	return &ProductAttribute{
		table:  "product_attribute",
		engine: engine,
	}
}

func (m *ProductAttribute) FindByCategoryAttributeId(categoryAttributeId int64, typ int32) []*pb.ProductAttribute {
	list := make([]*pb.ProductAttribute, 0)

	err := m.engine.Where("category_attribute_id = ?", categoryAttributeId).And("type = ?", typ).Find(&list)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}

func (m *ProductAttribute) FindByPage(cid int64, name string, types, start, size int32) []*pb.ProductAttribute {
	list := make([]*pb.ProductAttribute, 0)

	err := m.engine.Where("category_attribute_id = ?", cid).And("type = ?", types).And("name LIKE ?", "%"+name+"%").Limit(int(size), int(start)).Find(&list)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}

func (m *ProductAttribute) Count(cid int64, name string, types int32) int64 {
	var total int64 = 0
	total, err := m.engine.Table(m.table).Where("category_attribute_id = ?", cid).And("type = ?", types).And("name LIKE ?", "%"+name+"%").Count()
	if err != nil {
		logx.Error(err)
	}
	return total
}

func (m *ProductAttribute) DeleteById(id int64) bool {
	session := m.engine.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		logx.Error(err)
		return false
	}

	affected, err := session.Table(m.table).Where("id = ?", id).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	affected, err = session.Table(m.table+"_translations").Where("product_attribute_id = ?", id).Delete()

	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	err = session.Commit()
	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (m *ProductAttribute) Insert(in *pb.ProductAttribute) bool {
	_, err := m.engine.InsertOne(in)
	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}

func (m *ProductAttribute) Update(in *pb.ProductAttribute) bool {
	affected, err := m.engine.Where("id = ?", in.Id).Update(in)

	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	return true
}

func (m *ProductAttribute) FindOne(id int64) *pb.ProductAttribute {

	info := new(pb.ProductAttribute)
	has, err := m.engine.Where("id = ?", id).Get(info)

	if err != nil {
		logx.Error(err)
		return nil
	}

	if has == false {
		logx.Errorf("not found")
		return nil
	}

	return info
}
