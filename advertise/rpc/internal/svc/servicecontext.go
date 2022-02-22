package svc

import (
	"mallxx_server/advertise/models"
	"mallxx_server/advertise/rpc/internal/config"
	"mallxx_server/common/mysqlx"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	AdvertiseModel models.Advertise
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		AdvertiseModel: *models.NewAdvertise(engine),
		RedisClient:    redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass),
	}
}
