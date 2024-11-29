package jwts

import "net/http"

// Auth 登录认证
func Auth(r *http.Request) *Claims {
	token := r.Header.Get("token")
	return ParseToken(token)
}
