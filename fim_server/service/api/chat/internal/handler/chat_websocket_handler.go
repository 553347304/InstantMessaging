package handler

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/models/mtype"
	"fim_server/models/user_models"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"fim_server/service/rpc/file/file_rpc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	"net/http"
	"strings"
	"time"
	
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserWebsocketInfo struct {
	UserInfo           user_models.UserModel      // 用户ws连接对象
	Conn               *websocket.Conn            // 用户ws连接对象
	WebsocketClientMap map[string]*websocket.Conn // 用户管理所有客户端
}

var UserOnlineWebsocketMap = make(map[uint64]*UserWebsocketInfo)

type chatRequest struct {
	ReceiveId uint64        `json:"receive_id"`
	Type      mtype.Int8    `json:"type"`
	Message   mtype.Message `json:"message"`
}
type chatResponse struct {
	MessageId   uint           `json:"message_id"`
	IsMe        bool           `json:"is_me"`
	SendUser    mtype.UserInfo `json:"send_user"`
	ReceiveUser mtype.UserInfo `json:"receive_user"`
	Type        mtype.Int8     `json:"type"`
	Message     mtype.Message  `json:"message"`
	CreatedAt   time.Time      `json:"created_at"`
}
type Message struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Conn    *websocket.Conn
	Send    *UserWebsocketInfo
	Receive *UserWebsocketInfo
	Req     types.ChatRequest
	Request chatRequest
	Err     error
}

func (m *Message) Error(err string) error {
	m.Err = conv.Type(err).Error()
	return m.Err
}
func (m *Message) InsertDatabase() error {
	chatModel := chat_models.ChatModel{
		SendUserId:    m.Req.UserId,
		ReceiveUserId: m.Request.ReceiveId,
		Type:          m.Request.Type,
		Message:       m.Request.Message,
	}
	chatModel.Preview = chatModel.Message.GetPreview(m.Request.Type)
	
	err := m.svcCtx.DB.Create(&chatModel).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *Message) sendMessage(conn *websocket.Conn, t mtype.Int8, message mtype.Message, isMe bool) {
	conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(chatResponse{
		SendUser: mtype.UserInfo{
			UserId:   m.Send.UserInfo.ID,
			Username: m.Send.UserInfo.Username,
			Avatar:   m.Send.UserInfo.Avatar,
		},
		ReceiveUser: mtype.UserInfo{
			UserId:   m.Receive.UserInfo.ID,
			Username: m.Receive.UserInfo.Username,
			Avatar:   m.Receive.UserInfo.Avatar,
		},
		Type:      t,
		Message:   message,
		IsMe:      isMe,
		CreatedAt: time.Now(),
	}))
}
func (m *Message) SendUser(t mtype.Int8, message mtype.Message) {
	m.sendMessage(m.Conn, t, message, true)
}
func (m *Message) SendReceive(t mtype.Int8, message mtype.Message) {
	m.sendMessage(m.Receive.Conn, t, message, false)
}
func (m *Message) SendALLMember() error {
	err := m.InsertDatabase()
	if err != nil {
		return m.Error(err.Error())
	}
	_, ok1 := UserOnlineWebsocketMap[m.Request.ReceiveId] // 对付是否在线
	m.SendUser(m.Request.Type, m.Request.Message)         // 给自己发
	if ok1 && m.Request.ReceiveId != m.Req.UserId {
		m.SendReceive(m.Request.Type, m.Request.Message) // 给对方发
	}
	return nil
}
func (m *Message) SendOnlineGroup(wsMap map[string]*websocket.Conn, resp chatResponse) {
	for _, conn := range wsMap {
		conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(resp))
	}
}
func (m *Message) SendError(content string) {
	m.SendUser(mtype.MessageType.Error, mtype.Message{
		MessageError: &mtype.MessageError{
			Status:  "error",
			Content: content,
		},
	})
}
func (m *Message) isMessageFile() error {
	r := m.Request.Message.MessageFile
	if !strings.Contains(r.Src, "/") {
		return m.Error("请上传文件")
	}
	fileRpc, err := m.svcCtx.FileRpc.FileInfo(m.ctx, &file_rpc.FileInfoRequest{FileId: r.Src})
	if err != nil {
		return err
	}
	r.Title = fileRpc.Name
	r.Size = fileRpc.Size
	r.Ext = fileRpc.Ext
	return nil
}
func (m *Message) isMessageWithdraw() error {
	r := m.Request.Message.MessageWithdraw
	if r.MessageID == 0 {
		return m.Error("撤回消息id为空")
	}
	var chatModel chat_models.ChatModel
	
	err := m.svcCtx.DB.Take(&chatModel, r.MessageID).Error
	if err != nil {
		return m.Error("原消息不存在")
	}
	if chatModel.Type == mtype.MessageType.IsWithdraw {
		return m.Error("消息已经被撤回了")
	}
	if m.Req.UserId != m.Request.ReceiveId {
		return m.Error("只能撤回自己的消息")
	}
	// 自己撤回
	if m.Req.UserId == m.Request.ReceiveId {
		now := time.Now()
		if now.Sub(chatModel.CreatedAt) > 2*time.Minute {
			return m.Error("撤回消息时间超过两分钟")
		}
	}
	// 撤回消息
	m.svcCtx.DB.Model(&chatModel).Update("type", mtype.MessageType.IsWithdraw)
	r.Content = "你撤回了一条消息"
	return nil
}
func (m *Message) isMessageReply() error {
	r := m.Request.Message.MessageReply
	if r.MessageID == 0 {
		return m.Error("回复消息ID不能为空")
	}
	var chatModel chat_models.ChatModel
	err1 := m.svcCtx.DB.Take(&chatModel, r.MessageID).Error
	if err1 != nil {
		return m.Error("消息不存在")
	}
	if chatModel.Type == mtype.MessageType.IsWithdraw {
		return m.Error("消息已经被撤回了")
	}
	
	return nil
}
func (m *Message) isMessageVideoCall() error {
	r := m.Request.Message.MessageVideoCall
	// 先判断对方是否在线
	
	_, ok2 := UserOnlineWebsocketMap[m.Request.ReceiveId]
	if !ok2 {
		return m.Error("对方不在线")
	}
	
	key := fmt.Sprintf("%d_%d", m.Req.UserId, m.Request.ReceiveId)
	switch r.Flag {
	case 0:
		logs.Info("init")
		m.SendUser(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Flag: 1, Message: "等待对付接听"}})
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Flag: 2, Message: "等待对付接听"}})
	case 1: // 自己挂断
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Flag: 3, Message: "发起者已挂断"}})
	case 2: // 对方挂断
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Flag: 4, Message: "接收方拒绝视频通话"}})
	case 3: // 对方接受
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Flag: 5, Type: "create_offer", Message: "让发送者准备去发offer"}})
	case 4: // 我方正常挂断
		logs.Info("发起者已挂断", key)
		m.SendUser(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Flag: 6, Message: "发起者挂断了"}})
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Flag: 6, Message: "发起者挂断了"}})
		err := m.InsertDatabase() // 入库
		return err
	case 5: // 对方挂断
		logs.Info("对方正常挂断")
	}
	
	switch r.Type {
	case "offer":
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Type: "offer", Data: r.Data}})
		logs.Info("offer")
	case "answer":
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Type: "answer", Data: r.Data}})
		logs.Info("answer")
	case "offer_ice":
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Type: "offer_ice", Data: r.Data}})
		logs.Info("offer_ice")
	case "answer_ice":
		m.SendReceive(mtype.MessageType.VideoCall, mtype.Message{MessageVideoCall: &mtype.MessageVideoCall{Type: "offer_ice", Data: r.Data}})
		logs.Info("answer_ice")
	}
	return m.Error("视频通话")
}
func (m *Message) isMessage() error {
	if m.Request.Type == mtype.MessageType.File && m.Request.Message.MessageFile != nil {
		return m.isMessageFile() // 文件消息
	}
	if m.Request.Type == mtype.MessageType.Withdraw && m.Request.Message.MessageWithdraw != nil {
		return m.isMessageWithdraw() // 撤回消息
	}
	if m.Request.Type == mtype.MessageType.Reply && m.Request.Message.MessageReply != nil {
		return m.isMessageReply() // 回复消息
	}
	if m.Request.Type == mtype.MessageType.VideoCall && m.Request.Message.MessageVideoCall != nil {
		return m.isMessageVideoCall() // 回复消息
	}
	return nil
}
func (m *Message) isBan() error {
	is, _ := m.svcCtx.UserRpc.Curtail.IsCurtail(m.ctx, &user_rpc.ID{Id: m.Req.UserId})
	if is != nil && is.CurtailChat != "" {
		return m.Error(is.CurtailChat)
	}
	return nil
}
func (m *Message) Init(p []byte) error {
	if m.isBan() != nil || m.isMessage() != nil {
		return m.Error(m.Err.Error())
	}
	
	if !conv.Json().Unmarshal(p, &m.Request) {
		return m.Error("参数解析失败")
	}
	
	_, err := m.svcCtx.UserRpc.Friend.IsFriend(m.ctx, &user_rpc.IsFriendRequest{User1: uint32(m.Req.UserId), User2: uint32(m.Request.ReceiveId)})
	if err != nil {
		return m.Error(err.Error())
	}
	
	sendUser, ok1 := UserOnlineWebsocketMap[m.Req.UserId]
	if !ok1 {
		return m.Error("发送人不存在")
	}
	receiveUser, ok2 := UserOnlineWebsocketMap[m.Request.ReceiveId]
	if !ok2 {
		return m.Error("接收人不存在")
	}
	m.Send = sendUser
	m.Receive = receiveUser
	return nil
}
func (m *Message) Head() error {
	
	// 获取用户信息
	userRpc, err := m.svcCtx.UserRpc.User.UserInfo(m.ctx, &user_rpc.IdList{Id: []uint64{m.Req.UserId}})
	if err != nil {
		return m.Error(err.Error())
	}
	
	// 第一次进入  将用户信息存入map
	var user user_models.UserModel
	method.Struct().To(userRpc.Info, &user)
	UserOnlineWebsocketMap[m.Req.UserId] = &UserWebsocketInfo{
		Conn:     m.Conn,
		UserInfo: user,
		WebsocketClientMap: map[string]*websocket.Conn{
			m.Conn.RemoteAddr().String(): m.Conn,
		},
	}
	
	// Redis存入在线用户
	m.svcCtx.Redis.HSet("user_online", fmt.Sprint(m.Req.UserId), m.Req.UserId)
	
	friendRpc, err := m.svcCtx.UserRpc.Friend.FriendList(context.Background(), &user_rpc.ID{Id: m.Req.UserId})
	if err != nil {
		return m.Error(err.Error())
	}
	
	logs.Info("用户上线", m.Req.UserId, userRpc.Info.Username)
	
	for _, f := range friendRpc.FriendList {
		
		if f.Id == m.Req.UserId {
			continue // 剔除自己
		}
		
		friend, ok := UserOnlineWebsocketMap[f.Id]
		if ok {
			// 好友上线了
			if friend.UserInfo.UserConfigModel.FriendOnline {
				m.SendOnlineGroup(friend.WebsocketClientMap, chatResponse{
					Type: mtype.MessageType.Text,
					Message: mtype.Message{
						MessageText: &mtype.MessageText{
							Content: UserOnlineWebsocketMap[m.Req.UserId].UserInfo.Username + "上线了",
						},
					},
					CreatedAt: time.Now(),
				})
			}
		}
	}
	
	logs.Info(UserOnlineWebsocketMap)
	
	return nil
}
func (m *Message) Defer() {
	// 用户断开聊天
	defer func() {
		logs.Error("断开链接")
		m.Conn.Close()
		addr := m.Conn.RemoteAddr().String()
		userWsInfo, ok := UserOnlineWebsocketMap[m.Req.UserId]
		if ok {
			delete(userWsInfo.WebsocketClientMap, addr) // 删除退出的ws信息
		}
		delete(UserOnlineWebsocketMap, m.Req.UserId)
		m.svcCtx.Redis.HDel("user_online", fmt.Sprint(m.Req.UserId)) // Redis删除在线用户
	}()
}

func ChatWebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		
		m := Message{ctx: r.Context(), svcCtx: svcCtx, Req: req, Conn: src.Client().Websocket(w, r)}
		m.Defer()
		if m.Head() != nil {
			response.Response(r, w, nil, m.Err)
			return
		}
		
		for {
			// 消息类型，消息，错误
			_, p, err := m.Conn.ReadMessage()
			if err != nil {
				return
			}
			if m.Init(p) != nil || m.SendALLMember() != nil {
				m.SendError(m.Err.Error())
				continue
			}
		}
	}
}
