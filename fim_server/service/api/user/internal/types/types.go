// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddFriendRequest struct {
	UserId       uint64    `header:"User-ID"`
	FriendId     uint64    `json:"friend_id"`              // 好友ID
	ValidMessage string    `json:"valid_message,optional"` // 验证消息
	ValidInfo    ValidInfo `json:"valid_info,optional"`
}

type AddFriendResponse struct {
}

type Empty struct {
}

type FriendDeleteRequest struct {
	UserId   uint64 `header:"User-ID"`
	FriendId uint64 `json:"friend_id"` // 好友ID
}

type FriendDeleteResponse struct {
}

type FriendInfoRequest struct {
	UserId   uint64 `header:"User-ID"`
	Role     int32  `header:"Role"`
	FriendId uint64 `form:"friend_id"` // 好友ID
}

type FriendInfoResponse struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
	Notice   string `json:"notice"`    // 备注
	IsOnline bool   `json:"is_online"` // 是否在线
}

type FriendListRequest struct {
	PageInfo
	UserId uint64 `header:"User-ID"`
	Role   int32  `header:"Role"`
}

type FriendListResponse struct {
	List  []FriendInfoResponse `json:"list"`
	Total int64                `json:"total"`
}

type FriendNoticeUpdateRequest struct {
	UserId   uint64 `header:"User-ID"`
	FriendId uint64 `json:"friend_id"` // 好友ID
	Notice   string `json:"notice"`    // 备注
}

type FriendNoticeUpdateResponse struct {
}

type FriendValidInfo struct {
	UserId       uint64    `json:"user_id"`
	Username     string    `json:"username"`
	Avatar       string    `json:"avatar"`
	ValidMessage string    `json:"valid_message,optional"` // 验证消息
	ValidInfo    ValidInfo `json:"valid_info,optional"`
	Status       int32     `json:"status"`
	Auth         int32     `json:"auth"` // 好友验证
	Id           uint64    `json:"id"`
	Flag         string    `json:"flag"`       // send  rev
	CreatedAt    string    `json:"created_at"` // 验证时间
}

type FriendValidListRequest struct {
	PageInfo
	UserId uint64 `header:"User-ID"`
}

type FriendValidListResponse struct {
	List  []FriendValidInfo `json:"list"`
	Total int64             `json:"total"`
}

type Header struct {
	UserId uint64 `header:"User-ID"`
}

type PageInfo struct {
	Key   string `form:"key,optional"`
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
}

type ParamsPath struct {
	Id uint64 `path:"id"`
}

type RequestDelete struct {
	IdList []uint64 `json:"id_list"`
}

type ResponseList struct {
	Total int64   `json:"total"`
	List  []Empty `json:"list"`
}

type SearchInfo struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"`
}

type SearchRequest struct {
	PageInfo
	UserId uint64 `header:"User-ID"`
	Online bool   `form:"online,optional"`
}

type SearchResponse struct {
	List  []SearchInfo `json:"list"`
	Total int64        `json:"total"`
}

type User struct {
	UserId uint64 `header:"User-ID"`
}

type UserConfig struct {
	RecallMessage *string    `json:"recall_message,optional"` // 撤回消息内容
	FriendOnline  *bool      `json:"friend_online,optional"`  // 好友上线
	Sound         *bool      `json:"sound,optional"`          // 好友上线声音
	SecureLink    *bool      `json:"secure_link,optional"`    // 安全链接
	SavePassword  *bool      `json:"save_password,optional"`  // 保存密码
	SearchUser    *int32     `json:"search_user,optional"`    // 别人查找到你的方式
	Valid         *int32     `json:"valid,optional"`
	ValidInfo     *ValidInfo `json:"valid_info,optional"`
}

type UserCurtailRequest struct {
	UserId             uint64 `json:"user_id"`              // 限制的用户
	CurtailChat        bool   `json:"curtail_chat"`         // 限制聊天
	CurtailAddUser     bool   `json:"curtail_add_user"`     // 限制加人
	CurtailCreateGroup bool   `json:"curtail_create_group"` // 限制建群
	CurtailAddGroup    bool   `json:"curtail_add_group"`    // 限制加群
}

type UserInfo struct {
	Username *string `json:"username,optional"`
	Sign     *string `json:"sign,optional"`
	Avatar   *string `json:"avatar,optional"`
}

type UserInfoRequest struct {
	UserId uint64 `header:"User-ID"`
	Role   int32  `header:"Role"`
}

type UserInfoResponse struct {
	UserId        uint64     `json:"user_id"`
	Username      string     `json:"username"`
	Sign          string     `json:"sign"`
	Avatar        string     `json:"avatar"`
	RecallMessage *string    `json:"recall_message"` // 撤回消息内容
	FriendOnline  bool       `json:"friend_online"`  // 好友上线
	Sound         bool       `json:"sound"`          // 好友上线声音
	SecureLink    bool       `json:"secure_link"`    // 安全链接
	SavePassword  bool       `json:"save_password"`  // 保存密码
	SearchUser    int32      `json:"search_user"`    // 别人查找到你的方式
	Valid         *int32     `json:"valid,optional"`
	ValidInfo     *ValidInfo `json:"valid_info,optional"`
}

type UserListInfoResponse struct {
	ID                 uint64 `json:"id"`
	CreatedAt          string `json:"created_at"`
	Name               string `json:"name"`
	Avatar             string `json:"avatar"`
	IP                 string `json:"ip"`
	Addr               string `json:"addr"`
	IsOnline           bool   `json:"is_online"`
	SendMsgCount       int    `json:"send_msg_count"`       // 发送消息个数
	GroupAdminCount    int    `json:"group_admin_count"`    // 建群数量
	GroupCount         int    `json:"group_count"`          // 进群数量
	CurtailChat        bool   `json:"curtail_chat"`         // 限制聊天
	CurtailAddUser     bool   `json:"curtail_add_user"`     // 限制加人
	CurtailCreateGroup bool   `json:"curtail_create_group"` // 限制建群
	CurtailAddGroup    bool   `json:"curtail_add_group"`    // 限制加群
}

type UserListResponse struct {
	Total int64                  `json:"total"`
	List  []UserListInfoResponse `json:"list"`
}

type UserUpdateRequest struct {
	User
	UserInfo   *UserInfo   `json:"user_info,optional"`
	UserConfig *UserConfig `json:"user_config,optional"`
}

type UserUpdateResponse struct {
}

type ValidInfo struct {
	Issue  *string `json:"issue,optional"`
	Answer *string `json:"answer,optional"`
}

type ValidIssueRequest struct {
	UserId uint64 `header:"User-ID"`
	Id     uint64 `path:"id"` // 好友ID
}

type ValidIssueResponse struct {
	Valid     int32     `json:"valid"`
	ValidInfo ValidInfo `json:"valid_info"`
}

type ValidStatusRequest struct {
	UserId  uint64 `header:"User-ID"`
	ValidId uint64 `json:"valid_id"`
	Status  int32  `json:"status"`
}

type ValidStatusResponse struct {
	UserId  uint64 `header:"User-ID"`
	ValidId uint64 `json:"valid_id"`
}
