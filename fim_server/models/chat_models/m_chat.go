package chat_models

import (
	"fim_server/models"
	"fmt"
)

// ChatModel 对话表
type ChatModel struct {
	models.Model
	SendUserId     uint                  `json:"send_user_id"`                   // 发送人
	ReceiveUserId  uint                  `json:"receive_user_id"`                // 接收人
	MessageType    int8                  `json:"message_type"`                   // 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回 9 回复 10 引用
	MessagePreview string                `gorm:"size:64" json:"message_preview"` // 消息预览
	Message        models.Message        `json:"message"`                        // 消息内容
	SystemMessage  *models.SystemMessage `json:"system_message"`                 // 系统消息
}

func (chat ChatModel) MessagePreviewMethod() string {
	if chat.SystemMessage != nil {
		Map := map[int8]string{
			1: "涉黄",
			2: "涉恐",
			3: "涉政",
			4: "不正当言论",
		}
		return fmt.Sprintf("系统消息 - 该消息%s, 已被系统拦截", Map[chat.SystemMessage.Type])
	}

	messageMap := map[int8]string{
		1:  *chat.Message.Content,
		2:  "图片消息" + chat.Message.Image.Title,
		3:  "视频消息" + chat.Message.Image.Title,
		4:  "文件消息" + chat.Message.File.Title,
		5:  "语音消息",
		6:  "语音通话",
		7:  "视频通话",
		8:  "撤回消息" + chat.Message.Withdraw.Content,
		9:  "回复消息" + chat.Message.Reply.Content,
		10: "引用消息" + chat.Message.Quote.Content,
		11: "@消息" + chat.Message.At.Content,
	}
	return messageMap[chat.Message.Type]
}
