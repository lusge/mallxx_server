package svc

import (
	"mallxx_server/sign/api/internal/config"
	"mallxx_server/sign/rpc/signservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	SignRpc signservice.SignService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		SignRpc: signservice.NewSignService(zrpc.MustNewClient(c.SignRpc)),
	}
}
