package svc

import (
	"mallxx_server/advertise/api/internal/config"
	"mallxx_server/advertise/rpc/advertiseservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	AdvertiseRpc advertiseservice.AdvertiseService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AdvertiseRpc: advertiseservice.NewAdvertiseService(zrpc.MustNewClient(c.AdvertiseRpc)),
	}
}
