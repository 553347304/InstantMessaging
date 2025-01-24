package handler

import (
	"net/http"

	"fim_server/service/api/setting/internal/logic"
	"fim_server/service/api/setting/internal/svc"
	"fim_server/service/api/setting/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func SettingInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Empty
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewSettingInfoLogic(r.Context(), svcCtx)
		resp, err := l.SettingInfo(&req)
		response.Response(r, w, resp, err)
	}
}
