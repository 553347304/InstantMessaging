package group_models

import (
	"fim_server/models"
	"fim_server/models/mgorm"
	"fim_server/models/mtype"
	"fmt"
)

// GroupMessageModel 群聊消息表
type GroupMessageModel struct {
	models.Model
	GroupId        uint                 `json:"groupId"` // 发送人
	GroupModel     GroupModel           `gorm:"foreignKey:GroupId" json:"-"`
	SendUserId     uint                 `json:"send_user_id"`
	MessageType    mtype.MessageType    `json:"MessageType"`                    // 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回 9 回复 10 引用
	MessagePreview string               `gorm:"size:64" json:"message_preview"` // 消息预览
	Message        mtype.Message        `json:"message"`                        // 消息内容
	SystemMessage  *mtype.SystemMessage `json:"system_message"`                 // 系统消息
	DeleteUserId   mgorm.String         `json:"delete_user_id"`                 // 用户删除的聊天记录
}

func (chat GroupMessageModel) MessagePreviewMethod() string {
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
