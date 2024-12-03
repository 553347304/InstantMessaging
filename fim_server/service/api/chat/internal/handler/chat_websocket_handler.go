package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"

	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func ChatWebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		var upGrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// 鉴权 true表示放行，false表示拦截
				return true
			},
		}
		// 将http升级至websocket
		conn, err := upGrader.Upgrade(w, r, nil)
		fmt.Println(err)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		for {
			// 消息类型，消息，错误
			_, p, err := conn.ReadMessage()
			if err != nil {
				// 用户断开聊天
				fmt.Println(err)
				break
			}
			fmt.Println(string(p))
			// 发送消息
			conn.WriteMessage(websocket.TextMessage, []byte("xxx"))
		}
		defer conn.Close()
	}
}
