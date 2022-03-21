package svc

import (
	"mallxx_server/order/api/internal/config"
	"mallxx_server/order/rpc/orderservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrcerRpc orderservice.Orderservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrcerRpc: orderservice.NewOrderservice(zrpc.MustNewClient(c.OrderRpc)),
	}
}
