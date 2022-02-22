package models

import (
	"mallxx_server/advertise/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	adVertisementCachePrefix = "mallxxadmin:brand:language:%s:id:%d"
)

type Advertise struct {
	table  string
	engine *xorm.Engine
}

func NewAdvertise(engine *xorm.Engine) *Advertise {
	return &Advertise{
		table:  "advertise",
		engine: engine,
	}
}

func (m *Advertise) Insert(in *pb.Advertise) bool {
	in.AddTime = time.Now().Format("2006-01-02 15:04:05")
	_, err := m.engine.InsertOne(in)
	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (m *Advertise) Update(in *pb.Advertise) bool {
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

func (m *Advertise) Delete(id int64) bool {
	//
	affected, err := m.engine.Table(m.table).Where("id = ?", id).Delete()
	if err != nil {
		logx.Error(err)
		return false
	}
	if affected <= 0 {
		return false
	}

	return true
}

func (m *Advertise) FindAll(pos int32, categoryId int64) []*pb.Advertise {
	list := make([]*pb.Advertise, 0)
	currnetTime := time.Now().Format("2006-01-02 15:04:05")

	db := m.engine.Where("start_time <= ?", currnetTime)
	db.And("end_time >= ?", currnetTime)

	if pos >= 0 {
		db.And("pos = ?", pos)
	}

	if categoryId > 0 {
		db.And("category_id = ?", categoryId)
	}

	err := db.Find(&list)
	if err != nil {
		logx.Error(err)
	}

	return list
}

func (m *Advertise) Find(in *pb.AdRequest, start, pageSize int32) []*pb.Advertise {
	list := make([]*pb.Advertise, 0)
	db := m.engine.Where("1 = 1")
	if in.Pos > 0 {
		db.And("pos = ?", in.Pos)
	}

	if len(in.Name) > 0 {
		db.And("name LIKE ?", "%"+in.Name+"%")
	}

	if len(in.StartTime) > 0 {
		db.And("start_time <= ?", in.StartTime)
	}

	if len(in.EndTime) > 0 {
		db.And("end_time >= ?", in.EndTime)
	}

	db.Limit(int(pageSize), int(start))

	err := db.Find(&list)

	if err != nil {
		logx.Error(err)
	}
	return list
}

func (m *Advertise) Count(in *pb.AdRequest) int64 {
	db := m.engine.Where("1 = 1")
	if in.Pos > 0 {
		db.And("pos = ?", in.Pos)
	}

	if len(in.Name) > 0 {
		db.And("name LIKE ?", "%"+in.Name+"%")
	}

	if len(in.StartTime) > 0 {
		db.And("start_time <= ?", in.StartTime)
	}

	if len(in.EndTime) > 0 {
		db.And("end_time >= ?", in.EndTime)
	}

	total, err := db.Count()
	if err != nil {
		logx.Error(err)
	}
	return total
}

func (m *Advertise) FindOne(id int64) *pb.Advertise {
	ad := new(pb.Advertise)
	_, err := m.engine.Where("id = ?", id).Get(ad)
	if err != nil {
		logx.Error(err)
		return nil
	}
	return ad
}
