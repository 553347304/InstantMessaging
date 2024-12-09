package mtype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type MessageType int8

const (
	MessageTypeIText MessageType = iota + 1
	MessageTypeImage
	MessageTypeVideo
	MessageTypeFile
	MessageTypeVoice
	MessageTypeVoiceCall
	MessageTypeVideoCall
	MessageTypeWithdraw
	MessageTypeReply
	MessageTypeQuote
	MessageTypeAt
	MessageTypeTip
)

// Message 消息表
type Message struct {
	MessageType      MessageType       `json:"message_type"`                // 消息类型 MessageType
	MessageText      *MessageText      `json:"message_text,omitempty"`      // 文本消息
	MessageImage     *MessageImage     `json:"message_image,omitempty"`     // 图片消息
	MessageVideo     *MessageVideo     `json:"message_video,omitempty"`     // 视频消息
	MessageFile      *MessageFile      `json:"message_file,omitempty"`      // 文件消息
	MessageVoice     *MessageVoice     `json:"message_voice,omitempty"`     // 语音消息
	MessageVoiceCall *MessageVoiceCall `json:"message_voiceCall,omitempty"` // 语言通话
	MessageVideoCall *MessageVideoCall `json:"message_videoCall,omitempty"` // 视频通话
	MessageWithdraw  *MessageWithdraw  `json:"message_withdraw,omitempty"`  // 撤回消息
	MessageReply     *MessageReply     `json:"message_reply,omitempty"`     // 回复消息
	MessageQuote     *MessageQuote     `json:"message_quote,omitempty"`     // 引用消息
	MessageAt        *MessageAt        `json:"message_at,omitempty"`        // @用户的消息 群聊才有
	MessageTip       *MessageTip       `json:"message_tip,omitempty"`       // 提示消息   不入库
}

func (c *Message) Scan(value interface{}) error {
	err := json.Unmarshal(value.([]byte), c)
	if err != nil {
		return err
	}
	// 如果是撤回消息 就不返回
	if c.MessageType == MessageTypeWithdraw {
		c.MessageWithdraw = nil
	}
	return nil
}
func (c Message) Value() (driver.Value, error) { return json.Marshal(c) }

type MessageText struct {
	Content string `json:"content"`
}
type MessageImage struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type MessageVideo struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` // 时长 单位秒
}
type MessageFile struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` // 文件大小
	Ext   string `json:"ext"`  // 文件类型 word
}
type MessageVoice struct {
	Src  string `json:"src"`
	Time int    `json:"time"` // 时长 单位秒
}
type MessageVoiceCall struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}
type MessageVideoCall struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}

type MessageWithdraw struct {
	MessageId     uint     `json:"message_id"` // 撤回消息ID
	Content       string   `json:"content"`    // 撤回的提示词
	MessageOrigin *Message `json:"-"`          // 原消息
}
type MessageReply struct {
	MessageId     uint      `json:"message_id"` // 消息id
	UserId        uint      `json:"user_id"`
	Name          string    `json:"name"`
	Content       string    `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Message       *Message  `json:"message"`
	OriginMessage time.Time `json:"origin_message"` // 原消息时间
}
type MessageQuote struct {
	MessageId     uint      `json:"message_id"` // 消息id
	UserId        uint      `json:"user_id"`
	Name          string    `json:"name"`
	Content       string    `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Message       *Message  `json:"message"`
	OriginMessage time.Time `json:"origin_message"` // 原消息时间
}

type MessageAt struct {
	UserId  uint     `json:"user_id"`
	Content string   `json:"content"` // 回复的文本消息
	Message *Message `json:"message"`
}

type MessageTip struct {
	Status  string `json:"status"` // 提示状态
	Content string `json:"content"`
}
