package md5s

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(data []byte) string {
	h := md5.New()
	h.Write(data)
	sum := h.Sum(nil)
	hash := hex.EncodeToString(sum)
	return hash
}

// Check 验证密码  hash之后的密码  输入的密码
func Check(old, new []byte) bool {
	return Hash(old) == Hash(new)
}
