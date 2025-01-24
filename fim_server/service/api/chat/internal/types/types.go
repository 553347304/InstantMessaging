// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ChatDeleteRequest struct {
	UserId uint   `header:"User-Id"`
	IdList []uint `json:"id_list"`
}

type ChatDeleteResponse struct {
}

type ChatHistoryAdminRequest struct {
	SendUserID    uint   `form:"send_user_id"`
	ReceiveUserID uint   `form:"receive_user_id"`
	Key           string `form:"key,optional"`
	Page          int    `form:"page,optional"`
	Limit         int    `form:"limit,optional"`
}

type ChatHistoryRequest struct {
	UserId   uint `header:"User-Id"`
	Page     int  `form:"page,optional"`
	Limit    int  `form:"limit,optional"`
	FriendId uint `form:"friend_id"`
}

type ChatRequest struct {
	UserId uint `header:"User-Id"`
}

type ChatResponse struct {
}

type ChatSession struct {
	UserId         uint   `header:"User-Id"`
	Avatar         string `json:"avatar"`
	Name           string `json:"name"`
	CreatedAt      string `json:"created_at"`
	MessagePreview string `json:"message_preview"`
	IsTop          bool   `json:"is_top"`
}

type ChatSessionAdminRequest struct {
	ReceiveUserID uint `form:"receive_user_id"`
}

type ChatSessionRequest struct {
	UserId uint   `header:"User-Id"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
	Key    string `form:"key,optional"`
}

type ChatSessionResponse struct {
	List  []ChatSession `json:"list"`
	Total int64         `json:"total"`
}

type Empty struct {
}

type PageInfo struct {
	Key   string `form:"key,optional"`
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
}

type RequestDelete struct {
	IdList []uint `json:"id_list"`
}

type UserInfo struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UserInfoListResponse struct {
	List  []UserInfo `json:"list"`
	Total int64      `json:"total"`
}

type UserTopRequest struct {
	UserId   uint `header:"User-Id"`
	FriendId uint `json:"friend_id"`
}

type UserTopResponse struct {
}
