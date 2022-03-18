package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}

	CachePreifx struct {
		Address  string
		Member   string
		Follower string
		Level    string
	}
}
