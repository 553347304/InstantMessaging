package chat_models

// UserChatDeleteModels // 用户删除聊天记录表
type UserChatDeleteModels struct {
	UserId uint `json:"user_id"`
	ChatId uint `json:"chat_id"`
}
