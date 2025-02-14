package mtype

import (
	"database/sql/driver"
	"encoding/json"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"time"
)

type Int8 int8

var MessageType = struct {
	Null       Int8
	Image      Int8
	Text       Int8
	Video      Int8
	File       Int8
	Voice      Int8
	VoiceCall  Int8
	VideoCall  Int8
	Withdraw   Int8
	Reply      Int8
	At         Int8
	Tip        Int8
	IsWithdraw Int8
	Error      Int8
}{
	Null:       0,   // 未知消息
	Text:       1,   // 文本消息
	Image:      2,   // 图片消息
	Video:      3,   // 视频消息
	File:       4,   // 文件消息
	Voice:      5,   // 语音消息
	VoiceCall:  6,   // 语言通话
	VideoCall:  7,   // 视频通话
	Withdraw:   8,   // 撤回消息
	Reply:      9,   // 回复消息
	At:         11,  // @用户的消息 群聊才有
	Tip:        30,  // 系统提示
	IsWithdraw: 51,  // 已经被撤回的消息
	Error:      101, // 错误提示
}

type Message struct {
	MessageText      *MessageText      `json:"message_text,omitempty"`       // 文本消息
	MessageImage     *MessageImage     `json:"message_image,omitempty"`      // 图片消息
	MessageFile      *MessageFile      `json:"message_file,omitempty"`       // 文件消息
	MessageVideo     *MessageVideo     `json:"message_video,omitempty"`      // 视频消息
	MessageVoice     *MessageVoice     `json:"message_voice,omitempty"`      // 语音消息
	MessageWithdraw  *MessageWithdraw  `json:"message_withdraw,omitempty"`   // 撤回消息
	MessageReply     *MessageReply     `json:"message_reply,omitempty"`      // 回复消息
	MessageTip       *MessageTip       `json:"message_tip,omitempty"`        // 提示消息
	MessageError     *MessageError     `json:"message_error,omitempty"`      // 系统错误
	MessageVideoCall *MessageVideoCall `json:"message_video_call,omitempty"` // 视频通话
	MessageVoiceCall *MessageVoiceCall `json:"message_voice_call,omitempty"` // 语音通话
	MessageAt        *MessageAt        `json:"message_at,omitempty"`         // @用户的消息 群聊才有
}

func (v Message) Value() (driver.Value, error)  { return json.Marshal(v) }
func (v *Message) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), v) }

type MessageText struct {
	Content string `json:"content"`
}
type MessageImage struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type MessageFile struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` // 文件大小
	Ext   string `json:"ext"`  // 文件类型 word
}
type MessageVideo struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` // 时长 单位秒
}
type MessageVoice struct {
	Src  string `json:"src"`
	Time int    `json:"time"` // 时长 单位秒
}
type MessageWithdraw struct {
	MessageID uint   `json:"message_id"` // 撤回消息ID
	Content   string `json:"content"`    // 撤回的提示词
}
type MessageReply struct {
	MessageID     uint      `json:"message_id"` // 消息id
	UserId        uint      `json:"user_id"`
	Name          string    `json:"name"`
	Content       string    `json:"content"`        // 回复的文本消息，目前只能限制回复文本
	OriginMessage time.Time `json:"origin_message"` // 原消息时间
}
type MessageTip struct {
	Content string `json:"content"`
}
type MessageError struct {
	Status  string `json:"status"` // 提示状态
	Content string `json:"content"`
}
type MessageVideoCall struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
	Flag      int8      `json:"flag"`      // 标识，标识客户端弹框的模式
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	Data      any       `json:"data"`
}
type MessageVoiceCall struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}
type MessageAt struct {
	UserId  uint   `json:"user_id"`
	Content string `json:"content"` // 回复的文本消息
}

func (m *Message) GetPreview(t Int8) string {
	
	switch t {
	case MessageType.Null:
		return "[未知消息]"
	case MessageType.Text:
		return method.String(m.MessageText.Content).Slice(4)
	case MessageType.Image:
		return "[图片]"
	case MessageType.Video:
		return "[视频]"
	case MessageType.File:
		return "[文件]"
	case MessageType.Voice:
		return "[语音]"
	case MessageType.VoiceCall:
		return "[语音通话]"
	case MessageType.VideoCall:
		return "[视频通话]"
	case MessageType.Withdraw:
		return m.MessageWithdraw.Content
	case MessageType.Reply:
		return method.String(m.MessageText.Content).Slice(4)
	case MessageType.At:
		return "[@AT消息]"
	case MessageType.Tip:
		return "[系统提示]"
	case MessageType.Error:
		return "[系统错误]"
	}
	logs.Info(t)
	return "[类型错误]"
}
