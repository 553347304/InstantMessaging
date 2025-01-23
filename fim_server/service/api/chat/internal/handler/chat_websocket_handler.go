package handler

import (
	"context"
	"fim_server/common/service/redis_service"
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
	UserInfo           user_models.UserModel      // 用户ws连接对象
	Conn               *websocket.Conn            // 用户ws连接对象
	WebsocketClientMap map[string]*websocket.Conn // 用户管理所有客户端
}

var UserOnlineWebsocketMap = make(map[uint]*UserWebsocketInfo)

type chatRequest struct {
	ReceiveId uint               `json:"receive_id"`
	Message   mtype.MessageArray `json:"message"`
}
type chatResponse struct {
	MessageId     uint               `json:"message_id"`
	IsMe          bool               `json:"is_me"`
	SendUserId    mtype.UserInfo     `json:"send_user_id"`
	ReceiveUserId mtype.UserInfo     `json:"receive_user_id"`
	Message       mtype.MessageArray `json:"message"`
	CreatedAt     time.Time          `json:"created_at"`
}

// MessageInsertDatabaseChatModel 消息入库
func MessageInsertDatabaseChatModel(receiveUserId uint, sendUserId uint, message mtype.MessageArray, db *gorm.DB) uint {
	
	chatModel := chat_models.ChatModel{
		SendUserId:    sendUserId,
		ReceiveUserId: receiveUserId,
		Message:       message,
	}
	
	chatModel.Preview = chatModel.PreviewMethod()
	err := db.Create(&chatModel).Error
	if err != nil {
		sendUser, ok := UserOnlineWebsocketMap[sendUserId]
		if ok {
			SendTipErrorMessage(sendUser.Conn, "消息保存失败"+err.Error())
		}
	}
	return chatModel.ID
}

// SendMessageByUser 给谁发消息
func SendMessageByUser(svcCtx *svc.ServiceContext, receiveUserId uint, sendUserId uint, message mtype.MessageArray, messageId uint) {
	
	receiveUser, ok1 := UserOnlineWebsocketMap[receiveUserId]
	sendUser, _ := UserOnlineWebsocketMap[sendUserId]
	
	// 用户信息
	userBaseInfo, err5 := redis_service.GetUserInfo(svcCtx.Redis, svcCtx.UserRpc, receiveUserId)
	if err5 != nil {
		logs.Error(err5)
		return
	}
	
	resp := chatResponse{
		ReceiveUserId: mtype.UserInfo{
			ID:     receiveUserId,
			Name:   userBaseInfo.Name,
			Avatar: userBaseInfo.Avatar,
		},
		SendUserId: mtype.UserInfo{
			ID:     sendUserId,
			Name:   sendUser.UserInfo.Name,
			Avatar: sendUser.UserInfo.Avatar,
		},
		
		MessageId: messageId,
		Message:   message,
		CreatedAt: time.Now(),
	}
	
	resp.IsMe = true
	sendUser.Conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(resp)) // 给自己发
	if ok1 && receiveUserId != sendUserId {
		resp.IsMe = false
		receiveUser.Conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(resp)) // 接收方
	}
}

// SendTipErrorMessage 系统提示
func SendTipErrorMessage(conn *websocket.Conn, message string) {
	resp := chatResponse{
		Message: mtype.MessageArray{
			{Type: mtype.MessageType.Tip, State: "error", Content: message},
		},
		CreatedAt: time.Now(),
	}
	conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(resp))
}

func sendWsMapMessage(wsMap map[string]*websocket.Conn, data []byte) {
	for _, conn := range wsMap {
		conn.WriteMessage(websocket.TextMessage, data)
	}
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
		userResponse, err := svcCtx.UserRpc.User.UserInfo(context.Background(), &user_rpc.IdList{Id: []uint32{uint32(req.UserId)}})
		if err != nil {
			logs.Info("用户服务错误", err)
			response.Response(r, w, nil, err)
			return
		}
		var userConfigModel user_models.UserConfigModel
		conv.Json().Unmarshal(userResponse.Info.UserConfigModel, &userConfigModel)
		
		// 将用户信息存入map
		addr := conn.RemoteAddr().String()
		UserOnlineWebsocketMap[req.UserId] = &UserWebsocketInfo{
			Conn:     conn,
			UserInfo: conv.Struct(user_models.UserModel{}).Type(userResponse.Info),
			WebsocketClientMap: map[string]*websocket.Conn{
				conn.RemoteAddr().String(): conn,
			},
		}
		
		svcCtx.Redis.HSet("user_online", fmt.Sprint(req.UserId), req.UserId) // Redis存入在线用户
		
		friendResponse, err := svcCtx.UserRpc.Friend.FriendList(context.Background(), &user_rpc.ID{Id: uint32(req.UserId)})
		if err != nil {
			response.Response(r, w, nil, logs.Error("获取好友列表失败", err))
			return
		}
		
		logs.Info("用户上线", req.UserId, userResponse.Info.Name)
		
		for _, f := range friendResponse.FriendList {
			friend, ok := UserOnlineWebsocketMap[uint(f.Id)]
			if ok {
				// 好友上线了
				if friend.UserInfo.UserConfigModel.FriendOnline {
					sendWsMapMessage(friend.WebsocketClientMap, []byte(UserOnlineWebsocketMap[req.UserId].UserInfo.Name+"上线了"))
				}
			}
			
		}
		
		logs.Info(UserOnlineWebsocketMap)
		for {
			// 消息类型，消息，错误
			_, p, err := conn.ReadMessage()
			if err != nil {
				break
			}
			
			is, err1 := svcCtx.UserRpc.Curtail.IsCurtail(r.Context(), &user_rpc.ID{Id: uint32(req.UserId)})
			if err1 != nil || !is.CurtailChat.Is {
				SendTipErrorMessage(conn, is.CurtailChat.Error)
				continue
			}
			
			// 处理消息
			var request chatRequest
			if !conv.Json().Unmarshal(p, &request) {
				SendTipErrorMessage(conn, "参数解析失败")
				continue
			}
			
			// 自己和自己聊
			
			// 判断是否为好友
			_, err = svcCtx.UserRpc.Friend.IsFriend(context.Background(), &user_rpc.IsFriendRequest{User1: uint32(req.UserId), User2: uint32(request.ReceiveId)})
			if err != nil {
				SendTipErrorMessage(conn, "你们不是好友哦")
				continue
			}
			
			for _, m := range request.Message {
				// 文件消息
				if m.Type == mtype.MessageType.File {
					fileResponse, err1 := svcCtx.FileRpc.FileInfo(context.Background(), &file.FileInfoRequest{
						FileId: m.Content,
					})
					if err1 != nil {
						SendTipErrorMessage(conn, err1.Error())
						continue
					}
					// m.Title = fileResponse.Name
					m.Size = fileResponse.Size
				}
				// 撤回消息
				if m.Type == mtype.MessageType.Withdraw {
					var messageModel chat_models.ChatModel
					err1 := svcCtx.DB.Take(&messageModel, m.MessageId).Error
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
					if userConfigModel.RecallMessage != nil {
						content = *userConfigModel.RecallMessage
					}
					
					svcCtx.DB.Model(&messageModel).Updates(chat_models.ChatModel{
						Message: mtype.MessageArray{
							{Type: mtype.MessageType.Withdraw, Content: content, MessageId: m.MessageId},
						},
					})
				}
				// 回复消息
				if m.Type == mtype.MessageType.Reply {
					if m.MessageId == 0 {
						SendTipErrorMessage(conn, "撤回消息不能为空")
						continue
					}
					var messageModel chat_models.ChatModel
					err1 := svcCtx.DB.Take(&messageModel, m.MessageId).Error
					if err1 != nil {
						SendTipErrorMessage(conn, "消息不存在")
						continue
					}
					
					if messageModel.ID == conv.Type(mtype.MessageType.Withdraw).Uint() {
						SendTipErrorMessage(conn, "该消息已撤回")
						continue
					}
					
					if !(messageModel.SendUserId == req.UserId && messageModel.ReceiveUserId == request.ReceiveId ||
						messageModel.SendUserId == request.ReceiveId && messageModel.ReceiveUserId == req.UserId) {
						SendTipErrorMessage(conn, "只能回复自己或者对方的消息")
						continue
					}
				}
				// 消息入库
				id := MessageInsertDatabaseChatModel(request.ReceiveId, req.UserId, request.Message, svcCtx.DB)
				SendMessageByUser(svcCtx, request.ReceiveId, req.UserId, request.Message, id)
			}
		}
		
		// 用户断开聊天
		defer func() {
			logs.Error("断开链接")
			conn.Close()
			userWsInfo, ok := UserOnlineWebsocketMap[req.UserId]
			if ok {
				delete(userWsInfo.WebsocketClientMap, addr) // 删除退出的ws信息
			}
			delete(UserOnlineWebsocketMap, req.UserId)
			svcCtx.Redis.HDel("user_online", fmt.Sprint(req.UserId)) // Redis删除在线用户
		}()
	}
}
