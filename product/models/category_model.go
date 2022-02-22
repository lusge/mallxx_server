package models

import (
	"fmt"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type Category struct {
	engine *xorm.Engine
	table  string
}

var (
	categoryCacheUsernamePrefix = "mallxxadmin:category:%s:id:%d"
)

func NewCategory(engine *xorm.Engine) *Category {
	return &Category{
		table:  "category",
		engine: engine,
	}
}

//返回分类列表
//params language 语言
//params start, size 分页
//params pid 是分级查询
//return []*pb.Category 返回proto对应的成员数据
func (m *Category) FindByLimit(start, size int32, pid int64) []*pb.Category {

	list := make([]*pb.Category, 0)
	err := m.engine.Where("parent_id = ?", pid).Limit(int(size), int(start)).Find(&list)

	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}

//统计数量
func (m *Category) Count(pid int64) int64 {
	count, err := m.engine.Table(m.table).Where("parent_id = ?", pid).Count()
	if err != nil {
		logx.Error(err)
		return 0
	}

	return count
}

// update nav status
// return bool
func (m *Category) ChangeCategoryNavStatus(id int64, status int32) bool {
	affected, err := m.engine.Table("category").Where("id = ?", id).Update(map[string]interface{}{
		"nav_status": status,
	})

	if err != nil {
		// logx.Fatal(err, 222)
		fmt.Println(err, id, status)
		return false
	}
	if affected <= 0 {
		return false
	}
	return true
}

// update show status
// return bool
func (m *Category) ChangeCategoryShowStatus(id int64, status int32) bool {
	affected, err := m.engine.Table(m.table).Where("id = ?", id).Update(map[string]interface{}{
		"show_status": status,
	})

	if err != nil {
		logx.Error(err, 11)
		return false
	}
	if affected <= 0 {
		return false
	}
	return true
}

//insert
//returns last inserted ID and whether the insertion was successful
func (m *Category) Insert(in *pb.Category) (int64, bool) {
	var lastId int64 = 0

	session := m.engine.NewSession()
	defer session.Close()

	// start mysql transaction
	session.Begin()

	//set level
	// level := in.Level
	m.setLevel(in.ParentId, &in.Level)
	// in.Level = level

	_, err := session.InsertOne(in)
	if err != nil {
		logx.Error(err)
		return 0, false
	}

	lastId = in.Id
	if len(in.ProductAttributeIds) <= 0 {
		session.Commit() //Commit transaction
		return lastId, true
	}

	//insert category and product attribute relation by session
	if NewProductCategoryAttributeRelation(m.engine).InsertListBySession(session, lastId, in.ProductAttributeIds) == false {
		session.Rollback()
		logx.Error("Ids insert falied")
		return 0, false
	}

	err = session.Commit() //Commit transaction
	if err != nil {
		logx.Error(err)
		return 0, false
	}

	return lastId, true
}

//set level
func (m *Category) setLevel(pid int64, level *int32) {
	if pid == 0 {
		*level = 0
	} else {
		category := m.FindById(pid)
		if category == nil {
			*level = 0
		} else {
			*level = category.Level + 1
		}
	}
}

func (m *Category) FindById(id int64) *pb.Category {
	category := new(pb.Category)
	_, err := m.engine.Where("id = ?", id).Get(category)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return category
}

//update
//return whether the insertion was successful
func (m *Category) Update(in *pb.Category) (bool, error) {
	id := in.Id
	if id <= 0 {
		return false, fmt.Errorf("Id 错误")
	}

	session := m.engine.NewSession()
	defer session.Close()

	session.Begin()

	//delete product category attribute by category id
	//Update relation table first
	pcr := NewProductCategoryAttributeRelation(m.engine)

	if pcr.DeleteByCategoryIdSession(session, id) == false {
		session.Rollback()
		return false, nil
	}

	//insert product category attribute
	if pcr.InsertListBySession(session, id, in.ProductAttributeIds) == false {
		session.Rollback()
		return false, nil
	}

	//insert category
	_, err := session.Where("id = ?", id).Update(in)
	if err != nil {
		session.Rollback()
		logx.Error(err)
		return false, nil
	}

	err = session.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

//delete category
//return param true
func (m *Category) Delete(id int64) bool {
	session := m.engine.NewSession()
	defer session.Close()
	session.Begin()

	//
	affected, err := session.Table(m.table).Where("id = ?", id).Delete()
	if err != nil {
		goto ERR
	}
	if affected <= 0 {
		return false
	}

	//delete product_category_attribute_relation
	_, err = session.Table("product_category_attribute_relation").Where("product_category_id = ?", id).Delete()
	if err != nil {
		goto ERR
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

ERR:
	logx.Error(err)
	session.Rollback()
	return false
}

func (m *Category) FindListWithChildren() []*pb.Category {
	list := make([]*pb.Category, 0)

	err := m.engine.Find(&list)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return list
}
