package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type CategoryAttribute struct {
	engine *xorm.Engine
	table  string
}

var (
	categoryAttributeCacheUsernamePrefix = "mallxxadmin:categoryattribute:language:%s:id:"
)

func NewCategoryAttribute(engine *xorm.Engine) *CategoryAttribute {
	return &CategoryAttribute{
		table:  "category_attribute",
		engine: engine,
	}
}

//Find all category attribute
func (m *CategoryAttribute) FindAll() []*pb.CategoryAttribute {
	list := make([]*pb.CategoryAttribute, 0)

	err := m.engine.Find(&list)

	if err != nil {
		logx.Error(err)
		return nil
	}
	return list
}

//Find one by Category Attribute Id
func (m *CategoryAttribute) FindOne(id int64) *pb.CategoryAttribute {
	attr := new(pb.CategoryAttribute)

	has, err := m.engine.Where("id = ?", id).Get(attr)

	if !has || err != nil {
		return nil
	}

	return attr
}

//Find by Limit
func (m *CategoryAttribute) FindByLimit(start, size int32) []*pb.CategoryAttribute {
	list := make([]*pb.CategoryAttribute, 0)

	err := m.engine.Limit(int(size), int(start)).Find(&list)

	if err != nil {
		logx.Error(err)
		return nil
	}
	return list
}

//Find Count
//Return Total
func (m *CategoryAttribute) Count() int64 {
	var total int64 = 0
	total, err := m.engine.Table(m.table).Count()
	if err != nil {
		logx.Error(err)
	}

	return total
}

func (m *CategoryAttribute) Delete(id int64) bool {
	affected, err := m.engine.Table(m.table).Where("id = ?", id).Delete()

	if affected <= 0 {
		return false
	}

	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}

//insert
func (m *CategoryAttribute) Insert(in *pb.CategoryAttribute) bool {
	_, err := m.engine.InsertOne(in)

	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}

//udpate
func (m *CategoryAttribute) Update(in *pb.CategoryAttribute) bool {
	aff, err := m.engine.Table(m.table).
		Where("id = ?", in.Id).
		Update(map[string]interface{}{
			"name": in.Name,
		})

	if err != nil {
		logx.Error(err)
		return false
	}

	if aff <= 0 {
		return false
	}

	return true
}
