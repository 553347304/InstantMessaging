{{.head}}

package src

import (
	{{if .notStream}}"context"{{end}}

	{{.imports}}
)

type {{.src}}Server struct {
	svcCtx *svc.ServiceContext
	{{.unimplementedServer}}
}

func New{{.src}}Server(svcCtx *svc.ServiceContext) *{{.src}}Server {
	return &{{.src}}Server{
		svcCtx: svcCtx,
	}
}

{{.funcs}}
