package models

import (
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	brandCacheUsernamePrefix = "mallxxadmin:brand:id:%d"
)

type Brand struct {
	table  string
	engine *xorm.Engine
}

func NewBrand(engine *xorm.Engine) *Brand {
	return &Brand{
		table:  "brand",
		engine: engine,
	}
}

func (m *Brand) Insert(in *pb.Brand) bool {

	affected, err := m.engine.InsertOne(in)

	if affected <= 0 {
		return false
	}

	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}

func (m *Brand) Update(in *pb.Brand) bool {
	affected, err := m.engine.ID(in.Id).Update(in)
	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	return true
}

func (m *Brand) FindOne(id int64) *pb.Brand {
	brand := new(pb.Brand)

	has, err := m.engine.Where("id = ?", id).Get(brand)
	if err != nil {
		logx.Error(err)
		return nil
	}

	if !has {
		return nil
	}

	return brand
}

func (m *Brand) FindAll(name string, start, size int32) []*pb.Brand {
	list := make([]*pb.Brand, 0)
	err := m.engine.Where("name LIKE ?", "%"+name+"%").Limit(int(size), int(start)).Find(&list)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}

func (m *Brand) Count(name string) int64 {
	var total int64 = 0
	// fmt.Println(sql, params)
	total, err := m.engine.Table(m.table).Where("name LIKE ?", "%"+name+"%").Count()
	if err != nil {
		logx.Error(err)
	}

	return total
}

func (m *Brand) Delete(id int64) bool {
	_, err := m.engine.Table(m.table).Where("id = ?", id).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}

func (m *Brand) UpdateShowStatus(id int64, status int32) bool {
	if id <= 0 {
		return false
	}

	affected, err := m.engine.Table(m.table).Where("id = ?", id).Update(map[string]interface{}{
		"show_status": status,
	})
	if err != nil {
		logx.Error(err)
		return false
	}

	return affected > 0
}

func (m *Brand) UpdateFactoryStatus(id int64, status int32) bool {

	if id <= 0 {
		return false
	}

	affected, err := m.engine.Table(m.table).Where("id = ?", id).Update(map[string]interface{}{
		"factory_status": status,
	})
	if err != nil {
		logx.Error(err)
		return false
	}

	return affected > 0
}
