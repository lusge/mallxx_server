package svc

import (
	"mallxx_server/admin/models"
	"mallxx_server/admin/rpc/internal/config"
	"mallxx_server/common/mysqlx"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	AdminModel  models.Admin
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		AdminModel:  *models.NewAdmin(engine),
		RedisClient: redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass),
	}
}
