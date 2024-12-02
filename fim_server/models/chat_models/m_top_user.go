package chat_models

// TopUserModel 置顶用户表
type TopUserModel struct {
	UserId    uint `json:"user_id"`
	TopUserId uint `json:"top_user_id"`
}
