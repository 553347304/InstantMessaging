package mtype

type UserInfo struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
