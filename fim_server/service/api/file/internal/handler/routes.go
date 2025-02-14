// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	admin "fim_server/service/api/file/internal/handler/admin"
	"fim_server/service/api/file/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(src *rest.Server, serverCtx *svc.ServiceContext) {
	src.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/file/:name",
				Handler: ShowHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/file/upload",
				Handler: FileHandler(serverCtx),
			},
		},
	)

	src.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/file/admin/file",
					Handler: admin.FileListHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/file/admin/file",
					Handler: admin.FileDeleteHandler(serverCtx),
				},
			}...,
		),
	)
}
