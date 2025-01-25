package admin

import (
	"fim_server/models/setting_models"
	"net/http"

	"fim_server/service/api/setting/internal/logic/admin"
	"fim_server/service/api/setting/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func SettingInfoUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setting_models.ConfigModel
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := admin.NewSettingInfoUpdateLogic(r.Context(), svcCtx)
		resp, err := l.SettingInfoUpdate(&req)
		response.Response(r, w, resp, err)
	}
}
