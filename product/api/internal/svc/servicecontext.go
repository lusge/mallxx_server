package svc

import (
	"mallxx_server/product/api/internal/config"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productservices.ProductServices
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productservices.NewProductServices(zrpc.MustNewClient(c.ProductRpc)),
	}
}
