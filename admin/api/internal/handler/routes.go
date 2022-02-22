// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"mallxx_server/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/token/verify",
				Handler: verifyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: GetInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin"),
	)
}