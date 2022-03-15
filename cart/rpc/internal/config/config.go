package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	ProductRpc zrpc.RpcClientConf

	MemberRpc zrpc.RpcClientConf

	CachePreifx struct {
		Cart string
	}
}
