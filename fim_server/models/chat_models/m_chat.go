package chat_models

import (
	"fim_server/models"
	"fim_server/models/mtype"
)

// ChatModel 对话表
type ChatModel struct {
	models.Model
	SendUserId    uint                 `json:"send_user_id"`           // 发送人
	ReceiveUserId uint                 `json:"receive_user_id"`        // 接收人
	Type          mtype.Int8           `json:"type"`                   // 消息类型 0:成功|1:被撤回|2:删除
	Preview       string               `gorm:"size:64" json:"preview"` // 消息预览
	Message       mtype.MessageArray   `json:"message"`                // 消息内容
	SystemMessage *mtype.SystemMessage `json:"system_message"`         // 系统消息
}

func (chat ChatModel) PreviewMethod() string {
	var preview string
	for _, m := range chat.Message {
		preview += m.GetPreview()
	}
	return preview
}
