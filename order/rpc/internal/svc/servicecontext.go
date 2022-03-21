package svc

import (
	"mallxx_server/cart/rpc/cartservice"
	"mallxx_server/common/mysqlx"
	"mallxx_server/member/rpc/memberservice"
	"mallxx_server/order/models"
	"mallxx_server/order/rpc/internal/config"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productservices.ProductServices
	MemberRpc  memberservice.MemberService
	CartRpc    cartservice.CartService
	OrderModel *models.Order
	Redis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	redis := redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass)
	return &ServiceContext{
		Config:     c,
		ProductRpc: productservices.NewProductServices(zrpc.MustNewClient(c.ProductRpc)),
		MemberRpc:  memberservice.NewMemberService(zrpc.MustNewClient(c.MemberRpc)),
		CartRpc:    cartservice.NewCartService(zrpc.MustNewClient(c.CartRpc)),
		Redis:      redis,
		OrderModel: models.NewOrder(engine),
	}
}
