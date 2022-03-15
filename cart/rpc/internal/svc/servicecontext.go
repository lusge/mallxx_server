package svc

import (
	"fmt"
	"mallxx_server/cart/models"
	"mallxx_server/cart/rpc/internal/config"
	"mallxx_server/common/mysqlx"
	"mallxx_server/member/rpc/memberservice"
	"mallxx_server/product/rpc/productservices"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productservices.ProductServices
	MemberRpc  memberservice.MemberService
	CartModel  *models.Cart
	Redis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println("host", c.ProductRpc.Etcd.Hosts, "key", c.ProductRpc.Etcd.Key, "host", c.MemberRpc.Etcd.Hosts, "key", c.MemberRpc.Etcd.Key, "ccccc")

	engine := mysqlx.InitMysql(c.Mysql.DataSource)
	redis := redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass)
	return &ServiceContext{
		Config:     c,
		CartModel:  models.NewCart(engine),
		ProductRpc: productservices.NewProductServices(zrpc.MustNewClient(c.ProductRpc)),
		MemberRpc:  memberservice.NewMemberService(zrpc.MustNewClient(c.MemberRpc)),
		Redis:      redis,
	}

}
