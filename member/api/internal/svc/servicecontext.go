package svc

import (
	"mallxx_server/member/api/internal/config"
	"mallxx_server/member/rpc/memberservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	MemberRpc memberservice.MemberService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MemberRpc: memberservice.NewMemberService(zrpc.MustNewClient(c.MemberRpc)),
	}
}
