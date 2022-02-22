package models

import "xorm.io/xorm"

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
