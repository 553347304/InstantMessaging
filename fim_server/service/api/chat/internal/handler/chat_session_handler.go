package handler

import (
	"fim_server/service/server/response"
	"net/http"

	"fim_server/service/api/chat/internal/logic"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatSessionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatSessionRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewChatSessionLogic(r.Context(), svcCtx)
		resp, err := l.ChatSession(&req)
		response.Response(r, w, resp, err)
	}
}
