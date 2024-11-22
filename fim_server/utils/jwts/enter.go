package jwts

import (
	"github.com/golang-jwt/jwt/v4"
)

type PayLoad struct {
	UserId uint   `json:"userId"` // 用户id
	Name   string `json:"name"`   // 昵称
	Role   int8   `json:"role"`   // 用户权限
}

type Claims struct {
	PayLoad
	jwt.RegisteredClaims
}

// const issuer = "baiyin" // 签发人
const key = "key"  // 秘钥
const expires = 48 // 过期时间
