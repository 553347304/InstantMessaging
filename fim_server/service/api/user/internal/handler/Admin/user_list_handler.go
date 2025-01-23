package Admin

import (
	"net/http"

	"fim_server/service/api/user/internal/logic/Admin"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func UserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestPageInfo
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewUserListLogic(r.Context(), svcCtx)
		resp, err := l.UserList(&req)
		response.Response(r, w, resp, err)
	}
}
