package valid

import (
	"crypto/md5"
	"encoding/hex"
)

func (md5Service) Hash(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	sum := m.Sum(nil)
	hash := hex.EncodeToString(sum)
	return hash
}
func (md5 md5Service) Check(old, new string) bool {
	return md5.Hash(old) == md5.Hash(new)
}
