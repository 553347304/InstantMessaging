// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type GroupAddRequest struct {
	UserId     uint       `header:"User-Id"`
	GroupId    uint       `json:"group_id"`
	VerifyInfo VerifyInfo `json:"verify_info,optional"`
}

type GroupAddResponse struct {
}

type GroupAuthAddRequest struct {
	UserId  uint `header:"User-Id"`
	GroupId uint `json:"group_id"`
}

type GroupAuthAddResponse struct {
	Verify     int8       `json:"verify，optional"`
	VerifyInfo VerifyInfo `json:"verify_info,optional"`
}

type GroupCreateRequest struct {
	UserId     uint   `header:"User-Id"`
	Mode       int8   `json:"mode,optional"` // 模式 1 直接创建   2 创建模式
	Name       string `json:"name,optional"`
	IsSearch   bool   `json:"is_search,optional"`
	Size       int    `json:"size,optional"`
	UserIdList []uint `json:"user_id_list,optional"`
}

type GroupCreateResponse struct {
}

type GroupDeleteRequest struct {
	UserId uint `header:"User-Id"`
	Id     int8 `path:"id"`
}

type GroupDeleteResponse struct {
}

type GroupFriendsInfo struct {
	UserId    uint   `json:"user_id"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	IsInGroup bool   `json:"is_in_group"`
}

type GroupFriendsListRequest struct {
	UserId uint `header:"User-Id"`
	Id     uint `form:"id"`
}

type GroupFriendsListResponse struct {
	Total int64              `json:"total"`
	List  []GroupFriendsInfo `json:"list"`
}

type GroupInfoRequest struct {
	UserId uint `header:"User-Id"`
	Id     int8 `path:"id"`
}

type GroupInfoResponse struct {
	GroupId          uint       `json:"group_id"`
	Name             string     `json:"name"`
	Sign             string     `json:"sign"`
	Avatar           string     `json:"avatar"`
	MemberCount      int        `json:"member_count"`
	MemberOnlinCount int        `json:"member_onlin_count"`
	Leader           UserInfo   `json:"leader"` // 群主
	AdminList        []UserInfo `json:"admin_list"`
	Role             int8       `json:"role"` // 角色   1 群主 2 群管理员 3 群成员
}

type GroupMemberAddRequest struct {
	UserId       uint   `header:"User-Id"`
	Id           uint   `json:"id"`
	MemberIdList []uint `json:"member_id_list"`
}

type GroupMemberAddResponse struct {
}

type GroupMemberDeleteRequest struct {
	UserId   uint `header:"User-Id"`
	Id       uint `form:"id"`
	MemberId uint `form:"member_id"`
}

type GroupMemberDeleteResponse struct {
}

type GroupMemberInfo struct {
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	InOnline       bool   `json:"in_online"`
	Role           int8   `json:"role"`
	MemberName     string `json:"member_name"`
	CreatedAt      string `json:"created_at"`
	NewMessageDate string `json:"new_message_date"`
}

type GroupMemberNameRequest struct {
	UserId   uint   `header:"User-Id"`
	Id       uint   `json:"id"`
	MemberId uint   `json:"member_id"`
	Name     string `json:"name"`
}

type GroupMemberNameResponse struct {
}

type GroupMemberRequest struct {
	UserId uint   `header:"User-Id"`
	Id     uint   `form:"id"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
	Sort   string `form:"sort,optional"`
}

type GroupMemberResponse struct {
	Total int64             `json:"total"`
	List  []GroupMemberInfo `json:"list"`
}

type GroupMemberRoleRequest struct {
	UserId   uint `header:"User-Id"`
	Id       uint `json:"id"`
	MemberId uint `json:"member_id"`
	Role     int8 `json:"role"`
}

type GroupMemberRoleResponse struct {
}

type GroupSearchInfo struct {
	GroupId         uint   `json:"group_id"`
	Name            string `json:"name"`
	Sign            string `json:"sign"`
	Avatar          string `json:"avatar"`
	IsInGroup       bool   `json:"is_in_group"`
	UserCount       int    `json:"user_count"`
	UserOnlineCount int    `json:"user_online_count"`
}

type GroupSearchListRequest struct {
	UserId uint   `header:"User-Id"`
	Id     string `form:"id"`
	Key    string `form:"key"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
}

type GroupSearchListResponse struct {
	Total int64             `json:"total"`
	List  []GroupSearchInfo `json:"list"`
}

type GroupUpdateRequest struct {
	UserId             uint       `header:"User-Id"`
	Id                 int8       `json:"id"`
	Name               string     `json:"name,optional" conf:"name"`                                 // 群名
	Avatar             string     `json:"avatar,optional" conf:"avatar"`                             // 群头像
	Sign               string     `json:"sign,optional" conf:"sign"`                                 // 群简介
	IsSearch           *bool      `json:"is_search,optional" conf:"is_search"`                       // is搜索
	IsInvite           *bool      `json:"is_invite,optional" conf:"is_invite"`                       // is邀请
	IsTemporarySession *bool      `json:"is_temporary_session,optional" conf:"is_temporary_session"` // is临时会话
	IsForbiddenSpeech  *bool      `json:"is_forbidden_speech,optional" conf:"is_forbidden_speech"`   // is禁言
	Verify             int8       `json:"verify，optional"`
	VerifyInfo         VerifyInfo `json:"verify_info,optional"`
}

type GroupUpdateResponse struct {
}

type UserInfo struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type VerifyInfo struct {
	Issue  []string `json:"issue,optional"`
	Answer []string `json:"answer,optional"`
}
