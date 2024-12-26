package handler

import (
	"context"
	"fim_server/config/core"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
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

var UserOnlineMapWebsocket map[uint]*UserInfoWebsocket

type ChatRequest struct {
	GroupId uint          `json:"group_id"`
	Message mtype.Message `json:"message"`
}
type ChatResponse struct {
	UserId      uint              `json:"user_id"`
	Name        string            `json:"name"`
	Avatar      string            `json:"avatar"`
	Message     mtype.Message     `json:"message"`
	Id          uint              `json:"id"`
	MessageType mtype.MessageType `json:"message_type"`
	CreatedAt   time.Time         `json:"created_at"`
	IsMe        bool              `json:"is_me"`
}

type sendMessage struct {
	SvcCtx  *svc.ServiceContext
	Conn    *websocket.Conn
	Req     types.GroupChatRequest
	Request ChatRequest
}

func (s sendMessage) InsertDatabase() uint {
	if s.Request.Message.MessageType == mtype.MessageTypeWithdraw {
		logs.Info("撤回消息不入库")
		return 0
	}
	
	groupModel := group_models.GroupMessageModel{
		GroupId:     s.Request.GroupId,
		SendUserId:  s.Req.UserId,
		MessageType: s.Request.Message.MessageType,
		Message:     s.Request.Message,
	}
	groupModel.MessagePreview = groupModel.MessagePreviewMethod()
	err := s.SvcCtx.DB.Create(&groupModel).Error
	if err != nil {
		s.TipError("数据库插入失败: " + err.Error())
		return 0
	}
	return groupModel.ID
}
func (s sendMessage) TipError(message string) {
	resp := ChatResponse{
		Message: mtype.Message{
			MessageType: mtype.MessageTypeTip,
			MessageTip: &mtype.MessageTip{
				Status:  "error",
				Content: message,
			},
		},
		CreatedAt: time.Now(),
	}
	s.Conn.WriteMessage(websocket.TextMessage, conv.Marshal(resp))
}
func (s sendMessage) GroupOnlineUser(messageId uint) {
	
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
		UserId:      s.Req.UserId,
		Name:        info.UserInfo.Name,
		Avatar:      info.UserInfo.Avatar,
		Message:     s.Request.Message,
		Id:          messageId,
		MessageType: s.Request.Message.MessageType,
		CreatedAt:   time.Now(),
	}
	
	for _, u := range groupMemberOnlineIdList {
		wsUserInfo, ok := UserOnlineMapWebsocket[u]
		if !ok {
			continue
		}
		chatResponse.IsMe = wsUserInfo.UserInfo.ID == s.Req.UserId
		
		for _, w2 := range wsUserInfo.ConnMap {
			w2.WriteMessage(websocket.TextMessage, conv.Marshal(chatResponse))
		}
	}
}
func (s sendMessage) Init(p []byte) error {
	// 获取发送的消息
	if !conv.Unmarshal(p, &s.Request) {
		return conv.Type("消息格式错误").Error()
	}
	
	// 检查用户是否在群聊中
	var member group_models.GroupMemberModel
	err := s.SvcCtx.DB.Take(&member, "group_id = ? and user_id = ?", s.Request.GroupId, s.Req.UserId).Error
	if err != nil {
		return conv.Type("用户不是群成员").Error()
	}
	
	// 获取用户信息
	baseInfoResponse, err := s.SvcCtx.UserRpc.UserBaseInfo(context.Background(), &user_rpc.UserBaseInfoRequest{UserId: uint32(s.Req.UserId)})
	if err != nil {
		return conv.Type("获取用户信息失败" + err.Error()).Error()
	}
	userInfo := mtype.UserInfo{
		ID:     s.Req.UserId,
		Name:   baseInfoResponse.Name,
		Avatar: baseInfoResponse.Avatar,
	}
	UserOnlineMapWebsocket[s.Req.UserId] = &UserInfoWebsocket{
		UserInfo: userInfo,
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
		
		conn := core.Websocket(w, r)
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
			
			errMessage := SendMessage.Init(p)
			if errMessage != nil {
				SendMessage.TipError(errMessage.Error())
				continue
			}
			messageId := SendMessage.InsertDatabase()
			SendMessage.GroupOnlineUser(messageId)
		}
		
	}
}
