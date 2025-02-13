package open_api_sms

// 乐讯通官网 http://yun.loktong.com/

type SmsConfig struct {
	Username   string `json:"username"`    // 用户名
	Password   string `json:"password"`    // 密码
	Token      string `json:"token"`       // 产品总览页面对应产品的Token
	TemplateID string `json:"template_id"` // 模板管理报备的模板ID
	ReceiveTel string `json:"receive_tel"` // 接收人手机号
	Content    string `json:"content"`     // 内容
}
