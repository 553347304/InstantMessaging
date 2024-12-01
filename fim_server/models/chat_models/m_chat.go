package chat_models

import (
	models2 "fim_server/models"
)

// ChatModel 对话表
type ChatModel struct {
	models2.Model
	SendUserId     uint                   `json:"send_user_id"`                   // 发送人
	ReceiveUserId  uint                   `json:"receive_user_id"`                // 接收人
	MessageType    int8                   `json:"message_type"`                   // 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回 9 回复 10 引用
	MessagePreview string                 `gorm:"size:64" json:"message_preview"` // 消息预览
	Message        models2.Message        `json:"message"`                        // 消息内容
	SystemMessage  *models2.SystemMessage `json:"system_message"`                 // 系统消息
}
