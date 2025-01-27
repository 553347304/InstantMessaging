package open_api_qq

type LoginConfig struct {
	Code     string
	AppID    string
	AppKey   string
	Redirect string
}

type EmailConfig struct {
	Code        string `json:"code"`         // 授权码
	SendUser    string `json:"send_user"`    // 发件人
	ReceiveUser string `json:"receive_user"` // 接收人
	Title       string `json:"title"`        // 标题
	Content     string `json:"content"`      // 内容
}