package mtype

import (
	"database/sql/driver"
	"encoding/json"
	"fim_server/utils/stores/method"
	"time"
)

type Int8 int8

var MessageType = struct {
	Null      Int8
	Image     Int8
	Text      Int8
	Video     Int8
	File      Int8
	Voice     Int8
	VoiceCall Int8
	VideoCall Int8
	Withdraw  Int8
	Reply     Int8
	At        Int8
	Tip       Int8

	IsWithdraw Int8
}{
	Null:      0,  // 未知消息
	Text:      1,  // 文本消息
	Image:     2,  // 图片消息
	Video:     3,  // 视频消息
	File:      4,  // 文件消息
	Voice:     5,  // 语音消息
	VoiceCall: 6,  // 语言通话
	VideoCall: 7,  // 视频通话
	Withdraw:  8,  // 撤回消息
	Reply:     9,  // 回复消息
	At:        11, // @用户的消息 群聊才有
	Tip:       12, // 提示消息   不入库

	IsWithdraw: 51, // 已经被撤回的消息
}

type MessageArray []Message

func (s MessageArray) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}
func (s *MessageArray) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, s)
	}
	*s = make(MessageArray, 0)
	return nil
}

// Message 消息表
type Message struct {
	Type    Int8   `json:"type"`
	Content string `json:"content"`

	State     string `json:"state,omitempty"`      // 消息状态
	Size      int64  `json:"size,omitempty"`       // 图片大小 | 文件大小
	MessageId uint   `json:"message_id,omitempty"` // 撤回消息ID | 回复消息ID
}

func (m Message) GetPreview() string {
	maps := map[Int8]string{
		MessageType.Null:      "[未知消息]",
		MessageType.Text:      method.String(m.Content).Slice(4),
		MessageType.Image:     "[图片]",
		MessageType.Video:     "[视频]",
		MessageType.File:      "[文件]",
		MessageType.Voice:     "[语音]",
		MessageType.VoiceCall: "[语音通话]",
		MessageType.VideoCall: "[视频通话]",
		MessageType.Withdraw:  m.Content,
		MessageType.Reply:     method.String(m.Content).Slice(4),
		MessageType.At:        "提示消息",
		MessageType.Tip:       "提示消息",
	}
	return maps[m.Type]
}

// type Message struct {

// MessageType      MessageType       `json:"message_type"`                // 消息类型 MessageType
// MessageText      *MessageText      `json:"message_text,omitempty"`      //
// MessageImage     *MessageImage     `json:"message_image,omitempty"`     //
// MessageVideo     *MessageVideo     `json:"message_video,omitempty"`     //
// MessageFile      *MessageFile      `json:"message_file,omitempty"`      //
// MessageVoice     *MessageVoice     `json:"message_voice,omitempty"`     //
// MessageVoiceCall *MessageVoiceCall `json:"message_voiceCall,omitempty"` //
// MessageVideoCall *MessageVideoCall `json:"message_videoCall,omitempty"` //
// MessageWithdraw  *MessageWithdraw  `json:"message_withdraw,omitempty"`  //
// MessageReply     *MessageReply     `json:"message_reply,omitempty"`     //
// MessageQuote     *MessageQuote     `json:"message_quote,omitempty"`     //
// MessageAt        *MessageAt        `json:"message_at,omitempty"`        //
// MessageTip       *MessageTip       `json:"message_tip,omitempty"`       //
// }

//
// func (m *Message) Scan(value interface{}) error {
// 	err := json.Unmarshal(value.([]byte), m)
// 	if err != nil {
// 		return err
// 	}
// 	// 如果是撤回消息 就不返回
// 	if m.MessageType == MessageTypeWithdraw {
// 		m.MessageWithdraw = nil
// 	}
// 	return nil
// }
// func (m Message) Value() (driver.Value, error) { return json.Marshal(m) }

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
