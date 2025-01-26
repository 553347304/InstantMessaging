package handler

import (
	"context"
	"fim_server/common/service/redis_service"
	"fim_server/models"
	"fim_server/models/chat_models"
	"fim_server/models/mtype"
	"fim_server/models/user_models"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"fim_server/service/rpc/file/file"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fmt"
	"net/http"
	"time"
	
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserWebsocketInfo struct {
	UserInfo           user_models.UserModel      // 用户ws连接对象
	Conn               *websocket.Conn            // 用户ws连接对象
	WebsocketClientMap map[string]*websocket.Conn // 用户管理所有客户端
}

var UserOnlineWebsocketMap = make(map[uint]*UserWebsocketInfo)

type chatRequest struct {
	ReceiveId uint       `json:"receive_id"`
	Type      mtype.Int8 `json:"type"`
	Message   []byte     `json:"message"`
}
type chatResponse struct {
	MessageId     uint           `json:"message_id"`
	IsMe          bool           `json:"is_me"`
	SendUserId    mtype.UserInfo `json:"send_user_id"`
	ReceiveUserId mtype.UserInfo `json:"receive_user_id"`
	Message       []byte         `json:"message"`
	CreatedAt     time.Time      `json:"created_at"`
}

// SendMessageByUser 给谁发消息
func SendMessageByUser(svcCtx *svc.ServiceContext, receiveUserId uint, sendUserId uint, message mtype.Message, messageId uint) {
	
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

func sendWsMapMessage(wsMap map[string]*websocket.Conn, data []byte) {
	for _, conn := range wsMap {
		conn.WriteMessage(websocket.TextMessage, data)
	}
}

type Message struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Conn    *websocket.Conn
	Req     types.ChatRequest
	Content chatRequest
	Err     error
}
func (m *Message) Error(err string) error {
	m.Err = conv.Type(err).Error()
	return m.Err
}
func (m *Message) InsertDatabase(t mtype.Int8, message []byte) {
	chatModel := chat_models.ChatModel{
		SendUserId:    m.Req.UserId,
		ReceiveUserId: m.Content.ReceiveId,
		Message:       message,
		Type:          t,
		Preview:       mtype.GetPreview(t),
	}
	chatModel.Preview = chatModel
	
	err := m.svcCtx.DB.Create(&chatModel).Error
	if err != nil {
		sendUser, ok := UserOnlineWebsocketMap[sendUserId]
		if ok {
			SendTipErrorMessage(sendUser.Conn, "消息保存失败"+err.Error())
		}
	}
	return chatModel.ID
}
func (m *Message) SendErrorTip(content string) {
	m.Conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(mtype.MessageTip{
		Status:  "error",
		Content: content,
	}))
}

func (*Message) SendUser(content string) {
	
	// UserOnlineWebsocketMap[]
}
func (*Message) SendReceive(a interface{}) {
	
	// UserOnlineWebsocketMap[]
}

func (m *Message) MessageType(t mtype.Int8) {
	
	var data mtype.MessageFile
	conv.Json().Unmarshal(m.Content.Message, &data)
	fileRpc, err := m.svcCtx.FileRpc.FileInfo(m.ctx, &file.FileInfoRequest{FileId: data.Src})
	if err != nil {
		return err
	}
	
	type ChatModel struct {
		models.Model
		Type    mtype.Int8 `json:"type"`                   // 消息类型 0:成功|1:被撤回|2:删除
		Preview string     `gorm:"size:64" json:"preview"` // 消息预览
		Message []byte     `json:"message"`                // 消息内容
	}
	
	chat_models.ChatModel{
		Type: mtype.MessageType.File,
	}
	mtype.MessageFile{}
	request.Msg.FileMsg.Title = fileResponse.FileName
	request.Msg.FileMsg.Size = fileResponse.FileSize
	request.Msg.FileMsg.Type = fileResponse.FileType
	
	if err != nil {
		return
	}
	
	file.Title = fileResponse.Name
	file.Size
}
func (m *Message) MessageType() {
	
	if m.Content.Type != mtype.MessageType.File {
		
	}
	
	if t == mtype.MessageType.VideoCall {
		data := mtype.MessageVideoCall{
			Flag: 1,
		}
		
		switch data.Flag {
		case 0:
			m.Conn.WriteJSON(mtype.MessageVideoCall{
				Flag: 1,
			})
			m.SendReceive(mtype.MessageVideoCall{
				Flag: 2,
			})
		case 1:
			m.SendReceive(mtype.MessageVideoCall{
				Flag:    3,
				Message: "发起者已挂断",
			})
		case 2:
			// 接收者挂断
			m.SendReceive(mtype.MessageVideoCall{
				Flag:    3,
				Message: "用户拒绝了你的视频通话",
			})
		case 3:
			// 接收者接受
			m.SendReceive(mtype.MessageVideoCall{
				Flag:    5,
				Message: "建立连接",
				Type:    "create_offer",
			})
		case 4:
			// 通话中挂断
			
		}
		
		switch data.Message {
		case "offer":
			m.SendReceive(mtype.MessageVideoCall{
				Flag:    3,
				Message: "offer",
				Data:    data.Data,
			})
		case "answer":
			m.Conn.WriteJSON(mtype.MessageVideoCall{
				Message: "answer",
				Data:    data.Data,
			})
		case "offer_ice":
			m.SendReceive(mtype.MessageVideoCall{
				Flag:    3,
				Message: "offer_ice",
				Data:    data.Data,
			})
		case "answer_ice":
			m.SendReceive(mtype.MessageVideoCall{
				Flag:    3,
				Message: "answer_ice",
				Data:    data.Data,
			})
		}
		
		var receiveUserID uint = 1
		_, ok := UserOnlineWebsocketMap[receiveUserID]
		if !ok {
			m.SendUser("对方不在线")
		}
		
	}
	
	// UserOnlineWebsocketMap[]
}

func (m *Message) Init(p []byte) string {
	is, err1 := m.svcCtx.UserRpc.Curtail.IsCurtail(m.ctx, &user_rpc.ID{Id: uint32(req.UserId)})
	if err1 != nil || !is.CurtailChat.Is {
		return is.CurtailChat.Error
	}
	
	if !conv.Json().Unmarshal(p, &m.Content) {
		return "参数解析失败"
	}
	
	_, err := m.svcCtx.UserRpc.Friend.IsFriend(m.ctx, &user_rpc.IsFriendRequest{User1: uint32(m.Req.UserId), User2: uint32(m.Content.ReceiveId)})
	if err != nil {
		return err.Error()
	}
	
}

func ChatWebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		
		logs.Info(svcCtx)
		conn := src.Client().Websocket(w, r)
		logs.Info(conn)
		
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
			
			m := Message{
				ctx: r.Context(),
				Req: req,
			}
			m.Init(p)
			
			for _, m := range request.Message {
				// 文件消息
				
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
				
				ws.MessageType()
				
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
