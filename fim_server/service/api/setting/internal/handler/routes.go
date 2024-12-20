// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"fim_server/service/api/setting/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(src *rest.Server, serverCtx *svc.ServiceContext) {
	src.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/settings/open_login",
				Handler: open_login_infoHandler(serverCtx),
			},
		},
	)
}
