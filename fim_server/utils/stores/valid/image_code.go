package valid

import (
	"fim_server/utils/stores/logs"
	"github.com/mojocn/base64Captcha"
)

// go get github.com/mojocn/base64Captcha

type imageCodeResponse struct {
	ID     string `json:"id"`
	Answer string `json:"answer"`
	Base64 string `json:"base64"`
}

var store = base64Captcha.DefaultMemStore

func (imageCodeService) ImageView() imageCodeResponse {
	id, base64, answer, err := base64Captcha.NewCaptcha(&base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          4,
		Source:          "0123456789",
		BgColor:         nil,
		Fonts:           nil,
	}, store).Generate()
	if err != nil {
		logs.Fatal("图片验证码生成失败")
	}
	return imageCodeResponse{ID: id, Answer: answer, Base64: base64}
}
func (imageCodeService) Check(id string, value string) bool {
	return store.Verify(id, value, false)
}
