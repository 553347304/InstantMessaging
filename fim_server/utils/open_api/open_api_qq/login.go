package open_api_qq

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/https"
	"fim_server/utils/stores/logs"
	"net/url"
)

type qqResponse struct {
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	OpenID   string `json:"open_id"`
	Error    error
}
type qqOpenLoginResponse struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	IsLost          int    `json:"is_lost"`
	Nickname        string `json:"nickname"`
	Gender          string `json:"gender"`
	GenderType      int    `json:"gender_type"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Year            string `json:"year"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQq1    string `json:"figureurl_qq_1"`
	FigureurlQq2    string `json:"figureurl_qq_2"`
	FigureurlQq     string `json:"figureurl_qq"`
	IsYellowVip     string `json:"is_yellow_vip"`
	Vip             string `json:"vip"`
	YellowVipLevel  string `json:"yellow_vip_level"`
	Level           string `json:"level"`
	IsYellowYearVip string `json:"is_yellow_year_vip"`
}

func Login(qq LoginConfig) qqResponse {
	// 获取token和openID
	r := https.Get("https://graph.qq.com/oauth2.0/token", nil, map[string]any{
		"code":          qq.Code,
		"client_id":     qq.AppID,
		"client_secret": qq.AppKey,
		"redirect_uri":  qq.Redirect,
		"grant_type":    "authorization_code",
		"need_openid":   1,
	})
	if r.Error != nil {
		return qqResponse{Error: logs.Error(r.Error)}
	}
	v, _ := url.ParseQuery(string(r.Body))
	accessToken := v.Get("access_token")
	openID := v.Get("openid")
	
	// 获取用户信息
	info := https.Get("https://graph.qq.com/user/get_user_info", nil, map[string]any{
		"oauth_consumer_key": qq.AppID,
		"access_token":       accessToken,
		"openid":             openID,
	})
	if info.Error != nil {
		return qqResponse{Error: logs.Error(info.Error)}
	}
	
	var user qqOpenLoginResponse
	if !conv.Json().Unmarshal(info.Body, &user) {
		return qqResponse{Error: logs.Error("参数错误")}
	}
	if user.Ret != 0 {
		return qqResponse{Error: logs.Error(user.Msg)}
	}
	return qqResponse{
		Error:    nil,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Avatar:   user.FigureurlQq1,
		OpenID:   openID,
	}
}
