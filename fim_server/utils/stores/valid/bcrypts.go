package valid

import (
	"fim_server/utils/stores/logs"
	"golang.org/x/crypto/bcrypt"
)

func (bcryptService) Hash(value string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	if err != nil {
		logs.Warn(err)
	}
	return string(hash)
}
func (bcryptService) Check(hash string, value string) bool {
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(value))
	if err != nil {
		logs.Warn(err)
		return false
	}
	return true
}
