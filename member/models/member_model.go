package models

import (
	"mallxx_server/member/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type Member struct {
	engine *xorm.Engine
	table  string
}

func NewMember(engine *xorm.Engine) *Member {
	return &Member{
		table:  "member",
		engine: engine,
	}
}

func (m *Member) FindOneByUsername(username string) *pb.Member {
	member := new(pb.Member)
	isFind, err := m.engine.Table(m.table).Where("username = ?", username).Get(member)
	if err != nil {
		logx.Error(err)
		return nil
	}

	if isFind == false {
		return nil
	}
	return member
}

func (m *Member) Register(username, password string) (bool, *pb.Member) {
	info := new(pb.Member)
	info.AddTime = time.Now().Unix()
	info.Username = username
	info.Password = password
	info.MemberLevel = 1
	affected, err := m.engine.Table(m.table).InsertOne(info)
	if err != nil {
		logx.Error(err)
		return false, nil
	}

	if affected <= 0 {
		return false, nil
	}

	return true, info
}

func (m *Member) FindById(id int64) *pb.Member {
	info := new(pb.Member)
	_, err := m.engine.Where("id = ?", id).Get(info)
	if err != nil {
		logx.Error(err)
		return nil
	}

	return info
}
