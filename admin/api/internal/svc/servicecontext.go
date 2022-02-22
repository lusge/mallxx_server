package svc

import (
	"context"
	"fmt"
	"mallxx_server/admin/api/internal/config"
	"mallxx_server/admin/rpc/adminservice"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config   config.Config
	AdminRpc adminservice.AdminService
}

func timeInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	stime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Printf("调用 %s 方法 耗时: %v\n", method, time.Now().Sub(stime))
	return nil
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		AdminRpc: adminservice.NewAdminService(zrpc.MustNewClient(c.AdminRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor))),
	}
}
