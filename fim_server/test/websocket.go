package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(res http.ResponseWriter, req *http.Request) {
	// 服务升级
	conn, err := UP.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// 消息类型，消息，错误
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是：%s吗？", string(p))))
		fmt.Println(t, string(p))
	}
	defer conn.Close()
	fmt.Println("服务关闭")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

