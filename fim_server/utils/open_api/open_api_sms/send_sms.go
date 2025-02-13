package open_api_sms

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/https"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/valid"
	"fmt"
	"time"
)

type Response struct {
	Returnstatus string `json:"returnstatus"`
	Code         string `json:"code"`
	TaskID       []struct {
	} `json:"taskID"`
}

func Send(c SmsConfig) error {
	password := valid.MD5().Hash(c.Password)
	timestamp := time.Now().UnixMilli()

	param := https.Form(map[string]any{
		"action":    "sendtemplate",
		"username":  c.Username,
		"password":  password,
		"token":     c.Token,
		"timestamp": fmt.Sprint(timestamp),
	})
	data := https.Form(map[string]any{
		"templateid": c.TemplateID,
		"param":      fmt.Sprintf("%s|%s", c.ReceiveTel, c.Content),
		"rece":       "json",
		"sign":       valid.MD5().Hash(param),
	}) + "&" + param
	var scan Response
	r := https.Post("http://www.lokapi.cn/smsUTF8.aspx", nil, []byte(data))
	conv.Json().Unmarshal(r.Body, &scan)

	if scan.Returnstatus != "success" {
		return logs.Error(string(r.Body))
	}
	return nil
}
