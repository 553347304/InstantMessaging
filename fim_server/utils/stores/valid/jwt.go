package valid

import (
	"fim_server/utils/stores/logs"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type claims struct {
	PayLoad interface{}
	jwt.RegisteredClaims
}

func (j jwtService) Hash(payload interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{payload, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.expires))), // 过期时间
		Issuer:    j.issuer,                                                                 // 签发人
	}})
	sign, err := token.SignedString([]byte(j.key)) // 创建Token
	if err != nil {
		logs.Warn("生成 token 失败", err.Error())
		return ""
	}
	return sign
}

func (j jwtService) Parse(token string) *claims {
	if token == "" {
		logs.Warn("token为空")
		return nil
	}

	parse, _ := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.key), nil
	})

	// 验证token
	if parse != nil {
		_claims, ok := parse.Claims.(*claims)
		if ok && parse.Valid {
			return _claims
		}
	}
	logs.Warn("无效token")
	return nil
}
