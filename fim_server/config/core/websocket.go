package core

import (
	"fim_server/utils/stores/logs"
	"github.com/gorilla/websocket"
	"net/http"
)

func Websocket(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 鉴权   true 放行 | false 拦截
		},
	}
	// http upgrade websocket
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Error("websocket升级失败", err)
		return nil
	}
	return conn
}
