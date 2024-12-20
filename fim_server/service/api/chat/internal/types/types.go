// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ChatDeleteRequest struct {
	UserId uint   `header:"User-Id"`
	IdList []uint `json:"id_list"`
}

type ChatDeleteResponse struct {
}

type ChatHistoryRequest struct {
	UserId   uint `header:"User-Id"`
	Page     int  `form:"page,optional"`
	Limit    int  `form:"limit,optional"`
	FriendId uint `form:"friend_id"`
}

type ChatHistoryResponse struct {
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

type UserTopRequest struct {
	UserId   uint `header:"User-Id"`
	FriendId uint `json:"friend_id"`
}

type UserTopResponse struct {
}
