package open_api_qq

import (
	"gopkg.in/gomail.v2"
)



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

