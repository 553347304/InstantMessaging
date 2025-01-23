package service_method

import (
	"fim_server/service/server/response"
	"fim_server/utils/stores/conv"
	"net/http"
)

type authServiceInterface interface {
	IsAdmin(http.ResponseWriter, *http.Request) bool
}
type authService struct{}

//goland:noinspection GoExportedFuncWithUnexportedType
func Auth() authServiceInterface { return &authService{} }
func (l *authService) IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	role := r.Header.Get("role")
	if role != "1" {
		response.Response(r, w, nil, conv.Type("权限验证失败").Error())
		return false
	}
	return true
}
