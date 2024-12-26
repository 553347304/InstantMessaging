package chat_models

import (
	"fim_server/models"
	"fim_server/models/mtype"
	"fmt"
)

// ChatModel 对话表
type ChatModel struct {
	models.Model
	SendUserId     uint                 `json:"send_user_id"`                   // 发送人
	ReceiveUserId  uint                 `json:"receive_user_id"`                // 接收人
	MessageType    mtype.MessageType    `json:"message_type"`                   // 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回 9 回复 10 引用
	MessagePreview string               `gorm:"size:64" json:"message_preview"` // 消息预览
	Message        mtype.Message        `json:"message"`                        // 消息内容
	SystemMessage  *mtype.SystemMessage `json:"system_message"`                 // 系统消息
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
	return chat.Message.Preview()
}
