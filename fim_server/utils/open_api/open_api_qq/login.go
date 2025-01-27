package open_api_qq

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/method"
	"fmt"
	"net/url"
)

type qqResponse struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Avatar string `json:"avatar"`
	OpenID string `json:"open_id"`
	Error  error
}

func Login(qq LoginConfig) qqResponse {
	// 获取token和openID
	r := method.Http("https://graph.qq.com/oauth2.0/token").Get(map[string]any{
		"code":          qq.Code,
		"client_id":     qq.AppID,
		"client_secret": qq.AppKey,
		"redirect_uri":  qq.Redirect,
		"grant_type":    "authorization_code",
		"need_openid":   1,
	})
	if r.Error != nil {
		return qqResponse{Error: r.Error}
	}
	v, _ := url.ParseQuery(string(r.Body))
	accessToken := v.Get("access_token")
	openID := v.Get("openid")
	
	// 获取用户信息
	info := method.Http("https://graph.qq.com/user/get_user_info").Get(map[string]any{
		"oauth_consumer_key": qq.AppID,
		"access_token":       accessToken,
		"openid":             openID,
	})
	if info.Error != nil {
		return qqResponse{Error: info.Error}
	}
	
	var user = make(map[string]any)
	conv.Json().Unmarshal(info.Body, &user)
	return qqResponse{
		Error:  nil,
		Name:   fmt.Sprint(user["nickname"]),
		Gender: fmt.Sprint(user["gender"]),
		Avatar: fmt.Sprint(user["figureurl_qq_1"]),
		OpenID: openID,
	}
}
