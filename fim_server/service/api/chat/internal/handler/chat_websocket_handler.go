package handler

import (
	"context"
	"fim_server/models"
	"fim_server/models/chat_models"
	"fim_server/models/mtype"
	"fim_server/models/user_models"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
	"time"

	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

type UserWebsocketInfo struct {
	Conn     *websocket.Conn // 用户ws连接对象
	UserInfo user_models.UserModel
}

var UserWebsocketMap = make(map[uint]UserWebsocketInfo)

type chatRequest struct {
	ReceiveId uint          `json:"receive_id"`
	Message   mtype.Message `json:"message"`
}
type chatResponse struct {
	SendUserId    models.UserInfo `json:"send_user_id"`
	ReceiveUserId models.UserInfo `json:"receive_user_id"`
	Message       mtype.Message   `json:"message"`
	CreatedAt     time.Time       `json:"created_at"`
}

// MessageInsertDatabaseChatModel 消息入库
func MessageInsertDatabaseChatModel(receiveUserId uint, sendUserId uint, message mtype.Message, db *gorm.DB) {
	chatModel := chat_models.ChatModel{
		SendUserId:    sendUserId,
		ReceiveUserId: receiveUserId,
		MessageType:   message.MessageType,
		Message:       message,
	}
	chatModel.MessagePreview = chatModel.MessagePreviewMethod()
	err := db.Create(&chatModel).Error
	if err != nil {
		sendUser, ok := UserWebsocketMap[sendUserId]
		if ok {
			SendTipErrorMessage(sendUser.Conn, "消息保存失败")
		}
	}
}

// SendMessageByUser 给谁发消息
func SendMessageByUser(receiveUserId uint, sendUserId uint, message mtype.Message) {

	receiveUser, ok1 := UserWebsocketMap[receiveUserId]
	sendUser, ok2 := UserWebsocketMap[sendUserId]
	if !ok1 || !ok2 {
		return
	}

	resp := chatResponse{
		ReceiveUserId: models.UserInfo{
			ID:     receiveUserId,
			Name:   receiveUser.UserInfo.Name,
			Avatar: receiveUser.UserInfo.Avatar,
		},
		SendUserId: models.UserInfo{
			ID:     sendUserId,
			Name:   sendUser.UserInfo.Name,
			Avatar: sendUser.UserInfo.Avatar,
		},
		Message:   message,
		CreatedAt: time.Now(),
	}
	byteData := conv.Marshal(resp)
	receiveUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

// SendTipErrorMessage 系统提示
func SendTipErrorMessage(conn *websocket.Conn, message string) {
	resp := chatResponse{
		Message: mtype.Message{
			MessageType: mtype.MessageTypeTip,
			MessageTip: &mtype.MessageTip{
				Status:  "error",
				Content: message,
			},
		},
		CreatedAt: time.Now(),
	}
	conn.WriteMessage(websocket.TextMessage, conv.Marshal(resp))
}

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
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		// 获取用户信息
		res, err := svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
			UserId: uint32(req.UserId),
		})
		if err != nil {
			logs.Error("连接失败", err)
			response.Response(r, w, nil, err)
			return
		}

		var userInfo user_models.UserModel
		if !conv.Unmarshal(res.Data, &userInfo) {
			return
		}

		// 将用户信息存入map
		var userWebsocketInfo = UserWebsocketInfo{
			Conn:     conn,
			UserInfo: userInfo,
		}
		UserWebsocketMap[req.UserId] = userWebsocketInfo
		svcCtx.Redis.HSet("online", fmt.Sprint(req.UserId), req.UserId) // Redis存入在线用户

		friendRes, err := svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
			UserId: uint32(req.UserId),
		})
		if err != nil {
			response.Response(r, w, nil, logs.Error("获取好友列表失败", err))
			return
		}

		logs.Info("用户上线", req.UserId, userInfo.Name)

		for _, info := range friendRes.FriendList {
			friend, ok := UserWebsocketMap[uint(info.UserId)]
			if ok {
				// 好友上线了
				if friend.UserInfo.UserConfigModel.FriendOnline {
					friend.Conn.WriteMessage(websocket.TextMessage,
						[]byte(UserWebsocketMap[req.UserId].UserInfo.Name+"上线了"))
				}
			}

		}

		logs.Info(UserWebsocketMap)
		for {
			// 消息类型，消息，错误
			_, p, err := conn.ReadMessage()
			// 用户断开聊天
			if err != nil {
				conn.Close()
				delete(UserWebsocketMap, req.UserId)
				svcCtx.Redis.HDel("online", fmt.Sprint(req.UserId)) // Redis删除在线用户
				logs.Error("断开链接")
				logs.Info(UserWebsocketMap)
				break
			}

			// 处理消息
			var request chatRequest
			if !conv.Unmarshal(p, &request) {
				SendTipErrorMessage(conn, "参数解析失败")
				continue
			}

			// 自己和自己聊

			// 判断是否为好友
			_, err = svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{User1: uint32(req.UserId), User2: uint32(request.ReceiveId)})
			if err != nil {
				SendTipErrorMessage(conn, "你们不是好友哦")
				continue
			}
			SendMessageByUser(request.ReceiveId, req.UserId, request.Message)
			MessageInsertDatabaseChatModel(request.ReceiveId, req.UserId, request.Message, svcCtx.DB)
		}
	}
}
