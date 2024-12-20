package group_models

import (
	"fim_server/models"
	"fim_server/models/mtype"
)

// GroupMessageModel 群聊消息表
type GroupMessageModel struct {
	models.Model
	GroupId        uint                 `json:"groupId"` // 发送人
	GroupModel     GroupModel           `gorm:"foreignKey:GroupId" json:"-"`
	SendUserId     uint                 `json:"send_user_id"`
	MessageType    int8                 `json:"MessageType"`                    // 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回 9 回复 10 引用
	MessagePreview string               `gorm:"size:64" json:"message_preview"` // 消息预览
	Message        mtype.Message        `json:"message"`                        // 消息内容
	SystemMessage  *mtype.SystemMessage `json:"system_message"`                 // 系统消息
}
