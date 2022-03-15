package svc

import (
	"mallxx_server/common/mysqlx"
	"mallxx_server/member/models"
	"mallxx_server/member/rpc/internal/config"
)

type ServiceContext struct {
	Config              config.Config
	MemberLevelModel    *models.MemberLevel
	ReceiveAddressModel *models.ReceiveAddress
	MemberModel         *models.Member
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		MemberLevelModel:    models.NewMemberLevel(engine),
		ReceiveAddressModel: models.NewReceiveAddress(engine),
		MemberModel:         models.NewMember(engine),
	}
}
