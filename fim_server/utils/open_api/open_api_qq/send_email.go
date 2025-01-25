package open_api_qq

import (
	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	Code        string `json:"code"`         // 授权码
	SendUser    string `json:"send_user"`    // 发件人
	ReceiveUser string `json:"receive_user"` // 接收人
	Title       string `json:"title"`        // 标题
	Content     string `json:"content"`      // 内容
}

func SendEmail(e EmailConfig) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(e.SendUser, e.ReceiveUser))
	m.SetHeader("To", e.ReceiveUser)
	m.SetHeader("Subject", e.Title)
	m.SetBody("text/html", e.Content)
	d := gomail.NewDialer("smtp.qq.com", 465, e.SendUser, e.Code)
	err := d.DialAndSend(m)
	return err
}

