// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Empty struct {
}

type OpenLoginInfoResponse struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Href string `json:"href"` // 跳转地址
}

type SendEmailResponse struct {
	Code        string `json:"code"`
	SendUser    string `json:"send_user"`
	ReceiveUser string `json:"receive_user"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
