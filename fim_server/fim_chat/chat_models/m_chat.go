package chat_models

import "fim_server/common/models"

// Chat 对话表
type Chat struct {
	models.Model
	SendUserId     uint                  `json:"sendUserId"`                    // 发送人
	ReceiveUserId  uint                  `json:"receiveUserId"`                 // 接收人
	MessageType    int8                  `json:"messageType"`                   // 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回 9 回复 10 引用
	MessagePreview string                `gorm:"size:64" json:"messagePreview"` // 消息预览
	Message        models.Message        `json:"message"`                       // 消息内容
	SystemMessage  *models.SystemMessage `json:"systemMessage"`                 // 系统消息
}
