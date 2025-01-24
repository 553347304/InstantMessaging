// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	admin "fim_server/service/api/setting/internal/handler/admin"
	"fim_server/service/api/setting/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(src *rest.Server, serverCtx *svc.ServiceContext) {
	src.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/setting/info",
				Handler: SettingInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/setting/open_login",
				Handler: open_login_infoHandler(serverCtx),
			},
		},
	)

	src.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/api/setting/info",
					Handler: admin.SettingInfoUpdateHandler(serverCtx),
				},
			}...,
		),
	)
}
