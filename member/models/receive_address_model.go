package models

import (
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type ReceiveAddress struct {
	engine *xorm.Engine
	table  string
}

func NewReceiveAddress(engine *xorm.Engine) *ReceiveAddress {
	return &ReceiveAddress{
		engine: engine,
		table:  "receive_address",
	}
}

func (m *ReceiveAddress) FindAllByMemberId(memberId int64) []*pb.ReceiveAddress {
	list := make([]*pb.ReceiveAddress, 0)
	err := m.engine.Where("member_id = ?", memberId).Find(&list)

	if err != nil {
		logx.Error(err)
	}
	return list
}

func (m *ReceiveAddress) FindAll() []*pb.ReceiveAddress {
	list := make([]*pb.ReceiveAddress, 0)
	err := m.engine.Find(&list)

	if err != nil {
		logx.Error(err)
	}
	return list
}

func (m *ReceiveAddress) FindOneByMemberId(memberId, id int64) *pb.ReceiveAddress {
	info := new(pb.ReceiveAddress)
	_, err := m.engine.Where("member_id = ?", memberId).And("id = ?", id).Get(info)

	if err != nil {
		logx.Error(err)
		return nil
	}
	return info
}

func (m *ReceiveAddress) FindOneById(id int64) *pb.ReceiveAddress {
	info := new(pb.ReceiveAddress)
	_, err := m.engine.Where("id = ?", id).Get(info)

	if err != nil {
		logx.Error(err)
		return nil
	}
	return info
}

func (m *ReceiveAddress) Insert(in *pb.ReceiveAddress) bool {
	affected, err := m.engine.InsertOne(in)

	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	return true
}

func (m *ReceiveAddress) SetDefault(memberId, id int64) bool {
	_, err := m.engine.Table(m.table).Where("member_id = ?", memberId).Update(map[string]interface{}{
		"default_status": 0,
	})

	if err != nil {
		logx.Error(err)
		return false
	}

	affected, err := m.engine.Table(m.table).Where("id = ?", id).And("member_id = ?", memberId).Update(map[string]interface{}{
		"default_status": 1,
	})

	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	return true
}

func (m *ReceiveAddress) Update(in *pb.ReceiveAddress) bool {
	_, err := m.engine.Where("member_id = ?", in.MemberId).And("id = ?", in.Id).Update(in)

	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (m *ReceiveAddress) Delete(memberId, id int64) bool {
	affected, err := m.engine.Table(m.table).Where("member_id = ?", memberId).And("id = ?", id).Delete()

	if err != nil {
		logx.Error(err)
		return false
	}

	if affected <= 0 {
		return false
	}

	return true
}
