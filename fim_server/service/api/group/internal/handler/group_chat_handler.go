package handler

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/file/file_rpc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
	"time"
)

type UserInfoWebsocket struct {
	UserInfo mtype.UserInfo
	ConnMap  map[string]*websocket.Conn
}

var UserOnlineMapWebsocket = make(map[uint64]*UserInfoWebsocket)

type ChatRequest struct {
	GroupId uint64          `json:"group_id"`
	Type    mtype.Int8    `json:"type"`
	Message mtype.Message `json:"message"`
}
type ChatResponse struct {
	UserId    uint64          `json:"user_id"`
	Name      string        `json:"name"`
	Avatar    string        `json:"avatar"`
	Type      mtype.Int8    `json:"type"`
	Message   mtype.Message `json:"message"`
	Id        uint64          `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	IsMe      bool          `json:"is_me"`
}
type Message struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Conn    *websocket.Conn
	Req     types.GroupChatRequest
	Request ChatRequest
	Member  group_models.GroupMemberModel
	Err     error
}

func (m *Message) Error(err string) error {
	m.Err = conv.Type(err).Error()
	return m.Err
}
func (m *Message) InsertDatabase() (uint64, error) {
	groupModel := group_models.GroupMessageModel{
		GroupId:    m.Request.GroupId,
		SendUserId: m.Req.UserId,
		Message:    m.Request.Message,
		MemberId:   m.Member.ID,
		Type:       m.Request.Type,
	}
	groupModel.Preview = groupModel.Message.GetPreview(m.Request.Type)
	err := m.svcCtx.DB.Create(&groupModel).Error
	if err != nil {
		return groupModel.ID, err
	}
	return groupModel.ID, nil
}
func (m *Message) SendError(message string) {
	m.Conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(ChatResponse{
		Type: mtype.MessageType.Error,
		Message: mtype.Message{
			MessageError: &mtype.MessageError{Status: "error", Content: message},
		},
		CreatedAt: time.Now(),
	}))
}
func (m *Message) SendALLMember() error {
	messageID, err := m.InsertDatabase()
	if err != nil {
		return m.Error(err.Error())
	}
	
	// 用户在线列表
	var userOnlineIdList []uint64
	for u, _ := range UserOnlineMapWebsocket {
		userOnlineIdList = append(userOnlineIdList, u)
	}
	
	// 群成员在线列表
	var groupMemberOnlineIdList []uint64
	m.svcCtx.DB.Model(&group_models.GroupMemberModel{}).
		Where("group_id = ? and user_id in ?", m.Request.GroupId, userOnlineIdList).
		Select("user_id").Scan(&groupMemberOnlineIdList)
	
	info, _ := UserOnlineMapWebsocket[m.Req.UserId]
	var chatResponse = ChatResponse{
		UserId:    m.Req.UserId,
		Name:      info.UserInfo.Username,
		Avatar:    info.UserInfo.Avatar,
		Type:      m.Request.Type,
		Message:   m.Request.Message,
		Id:        messageID,
		CreatedAt: time.Now(),
	}
	
	for _, u := range groupMemberOnlineIdList {
		wsUserInfo, ok := UserOnlineMapWebsocket[u]
		if !ok {
			continue
		}
		chatResponse.IsMe = wsUserInfo.UserInfo.UserId == m.Req.UserId
		
		for _, w2 := range wsUserInfo.ConnMap {
			w2.WriteMessage(websocket.TextMessage, conv.Json().Marshal(chatResponse))
		}
	}
	return nil
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
	var groupMessage group_models.GroupMessageModel
	err := m.svcCtx.DB.Take(&groupMessage, r.MessageID).Error
	if err != nil {
		return m.Error("原消息不存在")
	}
	if groupMessage.Type == mtype.MessageType.IsWithdraw {
		return m.Error("消息已经被撤回了")
	}
	
	// 管理员和群主撤回
	if m.Member.Role == 1 || m.Member.Role == 2 {
		var messageUserRole int8 = 3
		m.svcCtx.DB.Model(group_models.GroupMemberModel{}).
			Where("group_id = ? and user_id = ?", m.Request.GroupId, groupMessage.SendUserId).
			Select("role").Scan(&messageUserRole)
		if messageUserRole == 1 || (messageUserRole == 2 && groupMessage.SendUserId != m.Req.UserId) {
			return m.Error("管理员只能撤回自己或普通用户的消息")
		}
	}
	
	// 自己撤回
	if m.Req.UserId == groupMessage.SendUserId {
		now := time.Now()
		if now.Sub(groupMessage.CreatedAt) > 2*time.Minute {
			return m.Error("撤回消息时间超过两分钟")
		}
	}
	
	// 撤回消息
	m.svcCtx.DB.Model(&groupMessage).Update("type", mtype.MessageType.IsWithdraw)
	r.Content = "你撤回了一条消息"
	return nil
}
func (m *Message) isMessageReply() error {
	r := m.Request.Message.MessageReply
	if r.MessageID == 0 {
		return m.Error("回复消息ID不能为空")
	}
	var groupMessage group_models.GroupMessageModel
	err1 := m.svcCtx.DB.Take(&groupMessage, r.MessageID).Error
	if err1 != nil {
		return m.Error("消息不存在")
	}
	if groupMessage.Type == mtype.MessageType.IsWithdraw {
		return m.Error("消息已经被撤回了")
	}
	return nil
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
	return nil
}
func (m *Message) isBan() error {
	if m.Member.GroupModel.IsBan {
		return m.Error("当前群正在全员禁言中")
	}
	if m.Member.BanTime != nil {
		return m.Error("当前用户备禁言中")
	}
	return nil
}
func (m *Message) Init(p []byte) error {
	// 获取发送的消息
	if !conv.Json().Unmarshal(p, &m.Request) {
		return m.Error("消息格式错误")
	}
	if m.Request.GroupId == 0 {
		return m.Error("群ID不能为空")
	}
	
	// 检查用户是否在群聊中
	var member group_models.GroupMemberModel
	err := m.svcCtx.DB.Preload("GroupModel").Take(&member, "group_id = ? and user_id = ?", m.Request.GroupId, m.Req.UserId).Error
	if err != nil {
		return m.Error("用户不在群")
	}
	m.Member = member
	
	if m.isBan() != nil || m.isMessage() != nil {
		return m.Error(m.Err.Error())
	}
	
	// 获取用户信息
	userRpc, err := m.svcCtx.UserRpc.User.UserInfo(m.ctx, &user_rpc.IdList{Id: []uint64{m.Req.UserId}})
	if err != nil {
		return m.Error(err.Error())
	}
	
	UserOnlineMapWebsocket[m.Req.UserId] = &UserInfoWebsocket{
		UserInfo: mtype.UserInfo{
			UserId:     m.Req.UserId,
			Username:   userRpc.Info.Username,
			Avatar: userRpc.Info.Avatar,
		},
		ConnMap: map[string]*websocket.Conn{
			m.Conn.RemoteAddr().String(): m.Conn,
		},
	}
	return nil
}

func GroupChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		
		var m = Message{ctx: r.Context(), svcCtx: svcCtx, Conn: src.Client().Websocket(w, r), Req: req}
		
		defer func() { m.Conn.Close() }()
		for {
			// 用户断开聊天
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
