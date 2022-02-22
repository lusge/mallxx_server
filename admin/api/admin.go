package main

import (
	"flag"
	"fmt"

	"mallxx_server/admin/api/internal/config"
	"mallxx_server/admin/api/internal/handler"
	"mallxx_server/admin/api/internal/svc"
	"mallxx_server/common/merrorx"
	"mallxx_server/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)

	server.Use(middleware.SetUserId)

	defer server.Stop()

	httpx.SetErrorHandler(merrorx.ErrorHandler)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
