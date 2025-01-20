// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Empty struct {
}

type LogListResponse struct {
	List  []Empty `json:"list"`
	Total int     `json:"total"`
}

type LogReadRequest struct {
	ID uint `path:"id"`
}

type LogRemoveRequest struct {
	IdList []uint `json:"id_list"`
}

type PageInfo struct {
	Key   string `form:"key,optional"`
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
}
