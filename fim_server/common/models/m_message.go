package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Message 消息表
type Message struct {
	Type      int8       `json:"type"`      // 消息类型 MessageType
	Content   *string    `json:"content"`   // 为1的时候使用
	Image     *Image     `json:"image"`     // 图片消息
	Video     *Video     `json:"video"`     // 视频消息
	File      *File      `json:"file"`      // 文件消息
	Voice     *Voice     `json:"voice"`     // 语音消息
	VoiceCall *VoiceCall `json:"voiceCall"` // 语言通话
	VideoCall *VideoCall `json:"videoCall"` // 视频通话
	Withdraw  *Withdraw  `json:"withdraw"`  // 撤回消息
	Reply     *Reply     `json:"reply"`     // 回复消息
	Quote     *Quote     `json:"quote"`     // 引用消息
	At        *At        `json:"at"`        // @用户的消息 群聊才有
}

// Scan 取出来的时候的数据
func (c *Message) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 入库的数据
func (c Message) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type Image struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type Video struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` // 时长 单位秒
}
type File struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` // 文件大小
	Type  string `json:"type"` // 文件类型 word
}
type Voice struct {
	Src  string `json:"src"`
	Time int    `json:"time"` // 时长 单位秒
}
type VoiceCall struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}
type VideoCall struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}

type Withdraw struct {
	Content   string   `json:"content"` // 撤回的提示词
	OriginMsg *Message `json:"-"`       // 原消息
}
type Reply struct {
	MsgID   uint     `json:"msgID"`   // 消息id
	Content string   `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg     *Message `json:"msg"`
}
type Quote struct {
	MsgID   uint     `json:"msgID"`   // 消息id
	Content string   `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg     *Message `json:"msg"`
}

type At struct {
	UserID  uint     `json:"userID"`
	Content string   `json:"content"` // 回复的文本消息
	Msg     *Message `json:"msg"`
}
