package models

import (
	"mallxx_server/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type Admin struct {
	engine *xorm.Engine
	table  string
}

func NewAdmin(engine *xorm.Engine) *Admin {
	return &Admin{
		table:  "admin",
		engine: engine,
	}
}

var (
	adminCacheUsernamePrefix = "mallxxadmin:admin:username:"
)

func (m *Admin) Insert(*pb.AdminInfo) (*pb.AdminInfo, error) {

	return nil, nil
}

func (m *Admin) FindOne(id int64) *pb.AdminInfo {
	adminInfo := &pb.AdminInfo{}
	has, err := m.engine.Table(m.table).Where("id = ?", id).Get(adminInfo)
	if err != nil {
		logx.Error(err)
		return nil
	}
	if !has {
		return nil
	}

	return adminInfo
}

func (m *Admin) FindByUsername(username string) *pb.AdminInfo {
	adminInfo := &pb.AdminInfo{}
	has, err := m.engine.Table(m.table).Where("username = ?", username).Get(adminInfo)
	if err != nil {
		logx.Error(err)
		return nil
	}
	if !has {
		return nil
	}

	return adminInfo
}

func (m *Admin) DeleteByid(id int64) error {
	return nil
}

func (m *Admin) DeleteByUsername(username string) error {
	return nil
}

func (m *Admin) Update(*pb.AdminInfo) error {

	return nil
}
