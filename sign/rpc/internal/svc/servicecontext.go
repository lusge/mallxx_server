package svc

import (
	"mallxx_server/common/mysqlx"
	"mallxx_server/sign/models"
	"mallxx_server/sign/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config           config.Config
	RedisClient      *redis.Redis
	MemberModel      *models.Member
	MemberLevelModel *models.MemberLevel
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	reClient := redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass)
	return &ServiceContext{
		Config:           c,
		RedisClient:      redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass),
		MemberModel:      models.NewMember(engine, reClient),
		MemberLevelModel: models.NewMemberLevel(engine),
	}
}
