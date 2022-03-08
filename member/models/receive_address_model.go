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

func (m *ReceiveAddress) FindAll(memberId int64) []*pb.ReceiveAddress {
	list := make([]*pb.ReceiveAddress, 0)
	err := m.engine.Where("member_id = ?", memberId).Find(&list)

	if err != nil {
		logx.Error(err)
	}

	return list
}
