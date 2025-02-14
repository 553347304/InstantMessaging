// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"fim_server/service/api/log/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(src *rest.Server, serverCtx *svc.ServiceContext) {
	src.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/log/logs",
					Handler: LogListHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/log/logs",
					Handler: LogRemoveHandler(serverCtx),
				},
			}...,
		),
	)
}
