package handler

import (
	"context"
	"fim_server/models"
	"fim_server/models/chat_models"
	"fim_server/models/mtype"
	"fim_server/models/user_models"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"fim_server/service/rpc/file/file"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
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
	MessageId     uint            `json:"message_id"`
	IsMe          bool            `json:"is_me"`
	SendUserId    models.UserInfo `json:"send_user_id"`
	ReceiveUserId models.UserInfo `json:"receive_user_id"`
	Message       mtype.Message   `json:"message"`
	CreatedAt     time.Time       `json:"created_at"`
}

// MessageInsertDatabaseChatModel 消息入库
func MessageInsertDatabaseChatModel(receiveUserId uint, sendUserId uint, message mtype.Message, db *gorm.DB) uint {
	switch message.MessageType {
	case mtype.MessageTypeWithdraw:
		logs.Info("撤回消息自己不入库")
		return 0
	}

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
			SendTipErrorMessage(sendUser.Conn, "消息保存失败"+err.Error())
		}
	}
	return chatModel.ID
}

// SendMessageByUser 给谁发消息
func SendMessageByUser(svcCtx *svc.ServiceContext, receiveUserId uint, sendUserId uint, message mtype.Message, messageId uint) {

	receiveUser, ok1 := UserWebsocketMap[receiveUserId]
	sendUser, _ := UserWebsocketMap[sendUserId]

	userBaseInfo, err := svcCtx.UserRpc.UserBaseInfo(context.Background(), &user_rpc.UserBaseInfoRequest{UserId: uint32(
		receiveUserId)})
	if err != nil {
		logs.Error(err)
		return
	}

	resp := chatResponse{
		ReceiveUserId: models.UserInfo{
			ID:     receiveUserId,
			Name:   userBaseInfo.Name,
			Avatar: userBaseInfo.Avatar,
		},
		SendUserId: models.UserInfo{
			ID:     sendUserId,
			Name:   sendUser.UserInfo.Name,
			Avatar: sendUser.UserInfo.Avatar,
		},

		MessageId: messageId,
		Message:   message,
		CreatedAt: time.Now(),
	}

	resp.IsMe = true
	sendUser.Conn.WriteMessage(websocket.TextMessage, conv.Marshal(resp)) // 给自己发
	if ok1 && receiveUserId != sendUserId {
		resp.IsMe = false
		receiveUser.Conn.WriteMessage(websocket.TextMessage, conv.Marshal(resp)) // 接收方
	}
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

			// 判断文件类型
			switch request.Message.MessageType {
			case mtype.MessageTypeFile:
				fileResponse, err1 := svcCtx.FileRpc.FileInfo(context.Background(), &file.FileInfoRequest{
					FileId: request.Message.MessageFile.Src,
				})
				if err1 != nil {
					SendTipErrorMessage(conn, err1.Error())
					continue
				}
				request.Message.MessageFile.Title = fileResponse.Name
				request.Message.MessageFile.Size = fileResponse.Size
				request.Message.MessageFile.Ext = fileResponse.Ext
			case mtype.MessageTypeWithdraw:
				var messageModel chat_models.ChatModel
				err1 := svcCtx.DB.Take(&messageModel, request.Message.MessageWithdraw.MessageId).Error
				if err1 != nil {
					SendTipErrorMessage(conn, "消息不存在")
					continue
				}
				if messageModel.SendUserId != req.UserId {
					SendTipErrorMessage(conn, "只能撤回自己的消息")
					continue
				}
				if time.Now().Sub(messageModel.CreatedAt) >= time.Minute*2 {
					SendTipErrorMessage(conn, "只能撤回2分钟内的消息")
					continue
				}

				// 撤回消息
				var content = "撤回了一条消息"
				if userInfo.UserConfigModel.RecallMessage != nil {
					content = *userInfo.UserConfigModel.RecallMessage
				}
				originMessage := messageModel.Message // 解决循环引用
				originMessage.MessageWithdraw = nil   // 撤回消息不能在撤回
				svcCtx.DB.Model(&messageModel).Updates(chat_models.ChatModel{
					MessagePreview: "[撤回消息] - " + content,
					Message: mtype.Message{
						MessageType: mtype.MessageTypeWithdraw, // 撤回消息
						MessageWithdraw: &mtype.MessageWithdraw{
							MessageId:     request.Message.MessageWithdraw.MessageId,
							Content:       "你" + content, // 撤回消息内容
							MessageOrigin: &originMessage,
						},
					},
				})
			case mtype.MessageTypeReply:
				if request.Message.MessageReply != nil || request.Message.MessageReply.MessageId == 0 {
					SendTipErrorMessage(conn, "撤回消息不能为空")
					continue
				}
				var messageModel chat_models.ChatModel
				err1 := svcCtx.DB.Take(&messageModel, request.Message.MessageReply.MessageId).Error
				if err1 != nil {
					SendTipErrorMessage(conn, "消息不存在")
					continue
				}

				if messageModel.MessageType == mtype.MessageTypeWithdraw {
					SendTipErrorMessage(conn, "该消息已撤回")
					continue
				}

				userBaseInfo, err2 := svcCtx.UserRpc.UserBaseInfo(context.Background(),
					&user_rpc.UserBaseInfoRequest{UserId: uint32(
						messageModel.SendUserId)})
				if err2 != nil {
					logs.Error(err2)
					continue
				}

				if !(messageModel.SendUserId == req.UserId && messageModel.ReceiveUserId == request.ReceiveId ||
					messageModel.SendUserId == request.ReceiveId && messageModel.ReceiveUserId == req.UserId) {
					SendTipErrorMessage(conn, "只能回复自己或者对方的消息")
					continue
				}
				request.Message.MessageReply.Message = &messageModel.Message
				request.Message.MessageReply.UserId = messageModel.SendUserId
				request.Message.MessageReply.Name = userBaseInfo.Name
				request.Message.MessageReply.OriginMessage = messageModel.CreatedAt
			case mtype.MessageTypeQuote:
				if request.Message.MessageQuote != nil || request.Message.MessageQuote.MessageId == 0 {
					SendTipErrorMessage(conn, "引用消息不能为空")
					continue
				}
				var messageModel chat_models.ChatModel
				err1 := svcCtx.DB.Take(&messageModel, request.Message.MessageQuote.MessageId).Error
				if err1 != nil {
					SendTipErrorMessage(conn, "消息不存在")
					continue
				}

				if messageModel.MessageType == mtype.MessageTypeWithdraw {
					SendTipErrorMessage(conn, "该消息已撤回")
					continue
				}

				userBaseInfo, err2 := svcCtx.UserRpc.UserBaseInfo(context.Background(),
					&user_rpc.UserBaseInfoRequest{UserId: uint32(
						messageModel.SendUserId)})
				if err2 != nil {
					logs.Error(err2)
					continue
				}

				if !(messageModel.SendUserId == req.UserId && messageModel.ReceiveUserId == request.ReceiveId ||
					messageModel.SendUserId == request.ReceiveId && messageModel.ReceiveUserId == req.UserId) {
					SendTipErrorMessage(conn, "只能回复自己或者对方的消息")
					continue
				}
				request.Message.MessageQuote.Message = &messageModel.Message
				request.Message.MessageQuote.UserId = messageModel.SendUserId
				request.Message.MessageQuote.Name = userBaseInfo.Name
				request.Message.MessageQuote.OriginMessage = messageModel.CreatedAt
			}

			// 消息入库
			id := MessageInsertDatabaseChatModel(request.ReceiveId, req.UserId, request.Message, svcCtx.DB)
			SendMessageByUser(svcCtx, request.ReceiveId, req.UserId, request.Message, id)
		}
	}
}
