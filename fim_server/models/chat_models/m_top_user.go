package chat_models

import "fim_server/models"

// TopUserModel 置顶用户表
type TopUserModel struct {
	models.Model
	UserId    uint `json:"user_id"`
	TopUserId uint `json:"top_user_id"`
}
