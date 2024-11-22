package jwts

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

// GenToken 加密 Token
func GenToken(payload PayLoad) (string, error) {
	claims := Claims{payload, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))), // 过期时间
		// Issuer:    issuer,                                                                 // 签发人
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key)) // 创建Token
}

// ParseToken 解析 Token
func ParseToken(tokenStr string) *Claims {
	if tokenStr == "" {
		log.Println("未携带token")
		return nil
	}

	token, _ := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(key), nil
	})

	if token != nil {
		// 验证token
		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			return claims
		}
	}
	log.Println("无效token")
	return nil
}
