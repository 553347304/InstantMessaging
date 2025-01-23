package Admin

import (
	"net/http"

	"fim_server/service/api/user/internal/logic/Admin"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func UserCurtailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCurtailRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewUserCurtailLogic(r.Context(), svcCtx)
		resp, err := l.UserCurtail(&req)
		response.Response(r, w, resp, err)
	}
}
