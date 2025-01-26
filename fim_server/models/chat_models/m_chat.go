package chat_models

import (
	"fim_server/models"
	"fim_server/models/mgorm"
	"fim_server/models/mtype"
)

// ChatModel 对话表
type ChatModel struct {
	models.Model
	SendUserId    uint          `json:"send_user_id"`           // 发送人
	ReceiveUserId uint          `json:"receive_user_id"`        // 接收人
	Type          mtype.Int8    `json:"type"`                   // 消息类型
	Preview       string        `gorm:"size:64" json:"preview"` // 消息预览
	Message       mtype.Message `json:"message"`                // 消息内容
	DeleteUserID  mgorm.String  `json:"delete_user_id"`         // 用户删除的聊天记录
}
