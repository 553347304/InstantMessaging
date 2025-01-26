package main

import (
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"
	"net/http"
	"os"
	
	"github.com/gorilla/websocket"
)

var offerClient *websocket.Conn
var answerClient *websocket.Conn

func checkStart() {
	if offerClient != nil && answerClient != nil {
		offerClient.WriteJSON(map[string]string{
			"type": "create_offer",
		})
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	
	conn := src.Client().Websocket(w, r)
	
	for {
		var obj map[string]any
		err := conn.ReadJSON(&obj)
		if err != nil {
			logs.Info("Error reading JSON:", err)
			break
		}
		
		logs.Info("recv:", obj)
		
		switch obj["type"] {
		case "connect":
			if offerClient == nil {
				offerClient = conn
				conn.WriteJSON(map[string]interface{}{
					"type":    "connect",
					"code":    200,
					"message": "connect success",
				})
				checkStart()
			} else if answerClient == nil {
				answerClient = conn
				conn.WriteJSON(map[string]interface{}{
					"type":    "connect",
					"code":    200,
					"message": "connect success",
				})
				checkStart()
			} else {
				conn.WriteJSON(map[string]interface{}{
					"type":    "connect",
					"code":    -1,
					"message": "connect failed",
				})
				conn.Close()
			}
		case "offer":
			if answerClient != nil {
				answerClient.WriteJSON(obj)
			}
		case "answer":
			if offerClient != nil {
				offerClient.WriteJSON(obj)
			}
		case "offer_ice":
			if answerClient != nil {
				answerClient.WriteJSON(obj)
			}
		case "answer_ice":
			if offerClient != nil {
				offerClient.WriteJSON(obj)
			}
		}
	}
	
	if conn == offerClient {
		logs.Info("remove offerClient")
		offerClient = nil
	} else if conn == answerClient {
		logs.Info("remove answerClient")
		answerClient = nil
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	logs.Info(123)
	byteData, err := os.ReadFile("index.html")
	if err != nil {
		logs.Error(err)
		return
	}
	w.Write(byteData)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ws", wsHandler)
	logs.Info("Server running on :9004")
	logs.Fatal(http.ListenAndServe(":9004", nil))
}

