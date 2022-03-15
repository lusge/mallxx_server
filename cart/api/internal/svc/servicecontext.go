package svc

import (
	"mallxx_server/cart/api/internal/config"
	"mallxx_server/cart/rpc/cartservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	CartRpc cartservice.CartService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		CartRpc: cartservice.NewCartService(zrpc.MustNewClient(c.CartRpc)),
	}
}
