package handler

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"

	"fim_server/service/server/response"
)

type UserInfoWebsocket struct {
	UserInfo mtype.UserInfo
	ConnMap  map[string]*websocket.Conn
}

var UserOnlineMapWebsocket = make(map[uint]*UserInfoWebsocket)

type ChatRequest struct {
	GroupId uint               `json:"group_id"`
	Message mtype.MessageArray `json:"message"`
}
type ChatResponse struct {
	UserId    uint               `json:"user_id"`
	Name      string             `json:"name"`
	Avatar    string             `json:"avatar"`
	Message   mtype.MessageArray `json:"message"`
	Id        uint               `json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	IsMe      bool               `json:"is_me"`
}
type sendMessage struct {
	SvcCtx  *svc.ServiceContext
	Conn    *websocket.Conn
	Req     types.GroupChatRequest
	Request ChatRequest
	Member  group_models.GroupMemberModel
	Err     error
}

func (s *sendMessage) Error(err string) error {
	s.Err = conv.Type(err).Error()
	return s.Err
}
func (s *sendMessage) InsertDatabase() uint {
	groupModel := group_models.GroupMessageModel{
		GroupId:    s.Request.GroupId,
		SendUserId: s.Req.UserId,
		Message:    s.Request.Message,
		MemberId:   s.Member.ID,
	}
	groupModel.Preview = groupModel.PreviewMethod()
	err := s.SvcCtx.DB.Create(&groupModel).Error
	if err != nil {
		s.TipError("数据库插入失败: " + err.Error())
		return 0
	}
	return groupModel.ID
}
func (s *sendMessage) TipError(message string) {
	resp := ChatResponse{
		Message: mtype.MessageArray{
			{Type: mtype.MessageType.Tip, State: "error", Content: message},
		},
		CreatedAt: time.Now(),
	}
	s.Conn.WriteMessage(websocket.TextMessage, conv.Json().Marshal(resp))
}
func (s *sendMessage) GroupOnlineUser(messageId uint) {

	// 用户在线列表
	var userOnlineIdList []uint
	for u, _ := range UserOnlineMapWebsocket {
		userOnlineIdList = append(userOnlineIdList, u)
	}

	// 群成员在线列表
	var groupMemberOnlineIdList []uint
	s.SvcCtx.DB.Model(&group_models.GroupMemberModel{}).
		Where("group_id = ? and user_id in ?", s.Request.GroupId, userOnlineIdList).
		Select("user_id").Scan(&groupMemberOnlineIdList)

	info, _ := UserOnlineMapWebsocket[s.Req.UserId]
	var chatResponse = ChatResponse{
		UserId:    s.Req.UserId,
		Name:      info.UserInfo.Name,
		Avatar:    info.UserInfo.Avatar,
		Message:   s.Request.Message,
		Id:        messageId,
		CreatedAt: time.Now(),
	}

	for _, u := range groupMemberOnlineIdList {
		wsUserInfo, ok := UserOnlineMapWebsocket[u]
		if !ok {
			continue
		}
		chatResponse.IsMe = wsUserInfo.UserInfo.ID == s.Req.UserId

		for _, w2 := range wsUserInfo.ConnMap {
			w2.WriteMessage(websocket.TextMessage, conv.Json().Marshal(chatResponse))
		}
	}
}
func (s *sendMessage) IsMessage(member group_models.GroupMemberModel) error {
	for _, m := range s.Request.Message {
		// 撤回消息
		if m.Type == mtype.MessageType.Withdraw {
			if m.MessageId == 0 {
				return s.Error("撤回消息id为空")
			}
			var groupMessage group_models.GroupMessageModel
			err := s.SvcCtx.DB.Take(&groupMessage, m.MessageId).Error
			if err != nil {
				return s.Error("原消息不存在")
			}
			if groupMessage.Type == mtype.MessageType.IsWithdraw {
				return s.Error("消息已经被撤回了")
			}

			// 管理员和群主撤回
			if member.Role == 1 || member.Role == 2 {
				var messageUserRole int8 = 3
				s.SvcCtx.DB.Model(group_models.GroupMemberModel{}).
					Where("group_id = ? and user_id = ?", s.Request.GroupId, groupMessage.SendUserId).
					Select("role").Scan(&messageUserRole)
				if messageUserRole == 1 || (messageUserRole == 2 && groupMessage.SendUserId != s.Req.UserId) {
					return s.Error("管理员只能撤回自己或普通用户的消息")
				}
			}

			// 自己撤回
			if s.Req.UserId == groupMessage.SendUserId {
				now := time.Now()
				if now.Sub(groupMessage.CreatedAt) > 2*time.Minute {
					return s.Error("撤回消息时间超过两分钟")
				}
			}

			// 撤回消息
			s.SvcCtx.DB.Model(&groupMessage).Update("type", mtype.MessageType.IsWithdraw)
			s.Request.Message[0].Content = "你撤回了一条消息"
		}
		// 回复消息
		if m.Type == mtype.MessageType.Reply {
			if m.MessageId == 0 {
				return s.Error("回复消息ID不能为空")
			}
			var groupMessage group_models.GroupMessageModel
			err1 := s.SvcCtx.DB.Take(&groupMessage, m.MessageId).Error
			if err1 != nil {
				return s.Error("消息不存在")
			}
			if groupMessage.Type == mtype.MessageType.IsWithdraw {
				return s.Error("消息已经被撤回了")
			}

		}
	}

	return nil
}
func (s *sendMessage) IsBan() error {
	if s.Member.GroupModel.IsBan {
		return s.Error("当前群正在全员禁言中")
	}
	if s.Member.BanTime != nil {
		return s.Error("当前用户备禁言中")
	}
	return nil
}
func (s *sendMessage) Init(p []byte) error {
	// 获取发送的消息
	if !conv.Json().Unmarshal(p, &s.Request) {
		return s.Error("消息格式错误")
	}

	// 检查用户是否在群聊中
	var member group_models.GroupMemberModel
	err := s.SvcCtx.DB.Preload("GroupModel").Take(&member, "group_id = ? and user_id = ?", s.Request.GroupId, s.Req.UserId).Error
	if err != nil {
		return s.Error("用户不是群成员")
	}
	s.Member = member

	if s.IsBan() != nil || s.IsMessage(member) != nil {
		return s.Err
	}

	// 获取用户信息
	userResponse, err := s.SvcCtx.UserRpc.User.UserInfo(context.Background(), &user_rpc.IdList{Id: []uint32{uint32(s.Req.UserId)}})
	if err != nil {
		return s.Error("获取用户信息失败" + err.Error())
	}
	UserOnlineMapWebsocket[s.Req.UserId] = &UserInfoWebsocket{
		UserInfo: mtype.UserInfo{
			ID:     s.Req.UserId,
			Name:   userResponse.Info.Name,
			Avatar: userResponse.Info.Avatar,
		},
		ConnMap: map[string]*websocket.Conn{
			s.Conn.RemoteAddr().String(): s.Conn,
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

		conn := src.Client().Websocket(w, r)
		if conn == nil {
			response.Response(r, w, nil, logs.Error("websocket连接失败"))
			return
		}
		var SendMessage = sendMessage{SvcCtx: svcCtx, Conn: conn, Req: req}

		defer func() { conn.Close() }()
		for {
			// 用户断开聊天
			_, p, err := conn.ReadMessage()
			if err != nil {
				return
			}

			err = SendMessage.Init(p)
			if err != nil {
				SendMessage.TipError(err.Error())
				continue
			}
			messageId := SendMessage.InsertDatabase()
			SendMessage.GroupOnlineUser(messageId)
		}

	}
}
