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

	switch chat.Message.MessageType {
	case mtype.MessageTypeIText:
		return chat.Message.MessageText.Content
	case mtype.MessageTypeImage:
		return "图片消息" + chat.Message.MessageImage.Title
	case mtype.MessageTypeVideo:
		return "视频消息" + chat.Message.MessageImage.Title
	case mtype.MessageTypeFile:
		return "文件消息" + chat.Message.MessageFile.Title
	case mtype.MessageTypeVoice:
		return "语音消息"
	case mtype.MessageTypeVoiceCall:
		return "语音通话"
	case mtype.MessageTypeVideoCall:
		return "视频通话"
	case mtype.MessageTypeWithdraw:
		return "撤回消息" + chat.Message.MessageWithdraw.Content
	case mtype.MessageTypeReply:
		return "回复消息" + chat.Message.MessageReply.Content
	case mtype.MessageTypeQuote:
		return "引用消息" + chat.Message.MessageQuote.Content
	case mtype.MessageTypeAt:
		return "@消息" + chat.Message.MessageAt.Content
	default:
		return ""
	}
}
