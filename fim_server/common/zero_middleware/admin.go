package zero_middleware

import (
	"fim_server/service/server/response"
	"fim_server/utils/stores/conv"
	"net/http"
)

func IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	role := r.Header.Get("role")
	if role != "1" {
		response.Response(r, w, nil, conv.Type("权限验证失败").Error())
		return false
	}
	return true
}
