// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddFriendRequest struct {
	UserId        uint       `header:"User-Id"`
	FriendId      uint       `json:"friend_id"`               // 好友ID
	VerifyMessage string     `json:"verify_message,optional"` // 验证消息
	VerifyInfo    VerifyInfo `json:"verify_info,optional"`
}

type AddFriendResponse struct {
}

type FriendDeleteRequest struct {
	UserId   uint `header:"User-Id"`
	FriendId uint `json:"friend_id"` // 好友ID
}

type FriendDeleteResponse struct {
}

type FriendInfoRequest struct {
	UserId   uint `header:"User-Id"`
	Role     int8 `header:"Role"`
	FriendId uint `form:"friend_id"` // 好友ID
}

type FriendInfoResponse struct {
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
	Notice   string `json:"notice"`    // 备注
	IsOnline bool   `json:"is_online"` // 是否在线
}

type FriendListRequest struct {
	UserId uint `header:"User-Id"`
	Role   int8 `header:"Role"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type FriendListResponse struct {
	List  []FriendInfoResponse `json:"list"`
	Total int64                `json:"total"`
}

type FriendNoticeUpdateRequest struct {
	UserId   uint   `header:"User-Id"`
	FriendId uint   `json:"friend_id"` // 好友ID
	Notice   string `json:"notice"`    // 备注
}

type FriendNoticeUpdateResponse struct {
}

type FriendVerifyInfo struct {
	UserId        uint       `json:"user_id"`
	Name          string     `json:"name"`
	Avatar        string     `json:"avatar"`
	VerifyMessage string     `json:"verify_message,optional"` // 验证消息
	VerifyInfo    VerifyInfo `json:"verify_info,optional"`
	Status        int8       `json:"status"`
	Auth          int8       `json:"auth"` // 好友验证
	Id            uint       `json:"id"`
	Flag          string     `json:"flag"` // send  rev
}

type FriendVerifyListRequest struct {
	UserId uint `header:"User-Id"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type FriendVerifyListResponse struct {
	List  []FriendVerifyInfo `json:"list"`
	Total int64              `json:"total"`
}

type SearchInfo struct {
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"`
}

type SearchRequest struct {
	UserId uint   `header:"User-Id"`
	Key    string `form:"key"`
	Online bool   `form:"online,optional"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
}

type SearchResponse struct {
	List  []SearchInfo `json:"list"`
	Total int64        `json:"total"`
}

type UserAuthRequest struct {
	UserId   uint `header:"User-Id"`
	FriendId uint `json:"friend_id"` // 好友ID
}

type UserAuthResponse struct {
	Verify     int8       `json:"verify，optional"`
	VerifyInfo VerifyInfo `json:"verify_info,optional"`
}

type UserInfoRequest struct {
	UserId uint `header:"User-Id"`
	Role   int8 `header:"Role"`
}

type UserInfoResponse struct {
	UserId        uint       `json:"user_id"`
	Name          string     `json:"name"`
	Sign          string     `json:"sign"`
	Avatar        string     `json:"avatar"`
	RecallMessage *string    `json:"recall_message"` // 撤回消息内容
	FriendOnline  bool       `json:"friend_online"`  // 好友上线
	Sound         bool       `json:"sound"`          // 好友上线声音
	SecureLink    bool       `json:"secure_link"`    // 安全链接
	SavePassword  bool       `json:"save_password"`  // 保存密码
	SearchUser    int8       `json:"search_user"`    // 别人查找到你的方式
	Verify        int8       `json:"verify，optional"`
	VerifyInfo    VerifyInfo `json:"verify_info,optional"`
}

type UserUpdateRequest struct {
	UserId        uint       `json:"user_id"`
	Name          *string    `json:"name,optional"`
	Sign          *string    `json:"sign,optional"`
	Avatar        *string    `json:"avatar,optional"`
	RecallMessage *string    `json:"recall_message,optional"` // 撤回消息内容
	FriendOnline  *bool      `json:"friend_online,optional"`  // 好友上线
	Sound         *bool      `json:"sound,optional"`          // 好友上线声音
	SecureLink    *bool      `json:"secure_link,optional"`    // 安全链接
	SavePassword  *bool      `json:"save_password,optional"`  // 保存密码
	SearchUser    *int8      `json:"search_user,optional"`    // 别人查找到你的方式
	Verify        int8       `json:"verify，optional"`
	VerifyInfo    VerifyInfo `json:"verify_info,optional"`
}

type UserUpdateResponse struct {
}

type VerifyInfo struct {
	Issue  []string `json:"issue,optional"`
	Answer []string `json:"answer,optional"`
}

type VerifyStatusRequest struct {
	UserId   uint `header:"User-Id"`
	VerifyId uint `json:"verify_id"`
	Status   int8 `json:"status"`
}

type VerifyStatusResponse struct {
	UserId   uint `header:"User-Id"`
	VerifyId uint `json:"verify_id"`
}
