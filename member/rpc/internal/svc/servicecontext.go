package svc

import (
	"mallxx_server/common/mysqlx"
	"mallxx_server/member/models"
	"mallxx_server/member/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config              config.Config
	MemberLevelModel    *models.MemberLevel
	ReceiveAddressModel *models.ReceiveAddress
	MemberModel         *models.Member
	Redis               *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	redis := redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass)
	return &ServiceContext{
		Config:              c,
		MemberLevelModel:    models.NewMemberLevel(engine),
		ReceiveAddressModel: models.NewReceiveAddress(engine),
		MemberModel:         models.NewMember(engine),
		Redis:               redis,
	}
}
