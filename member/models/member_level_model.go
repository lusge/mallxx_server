package models

import (
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

var (
	memberLevelCachePrefix = "mallxxadmin:member_level:language:%s:id:%d"
)

type MemberLevel struct {
	engine *xorm.Engine
	table  string
}

func NewMemberLevel(engine *xorm.Engine) *MemberLevel {
	return &MemberLevel{
		table:  "member_level",
		engine: engine,
	}
}

func (m *MemberLevel) FindAll() []*pb.MemberLevel {
	list := make([]*pb.MemberLevel, 0)

	err := m.engine.Find(&list)
	if err != nil {
		logx.Error(err)
	}

	return list
}

func (m *MemberLevel) FindOneById(id int64) *pb.MemberLevel {
	if id <= 0 {
		return nil
	}

	level := new(pb.MemberLevel)
	isFind, err := m.engine.Where("id = ?", id).Get(level)
	if err != nil {
		logx.Error(err)
		return nil
	}

	if isFind == false {
		return nil
	}

	return level
}
