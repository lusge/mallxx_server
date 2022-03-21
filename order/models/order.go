package models

import "xorm.io/xorm"

type Order struct {
	engine *xorm.Engine
	table  string
}

func NewOrder(engine *xorm.Engine) *Order {
	return &Order{
		table:  "order",
		engine: engine,
	}
}
