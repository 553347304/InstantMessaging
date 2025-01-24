// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Empty struct {
}

type FileInfo struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"path"`
	Hash string `json:"hash"`
}

type FileListResponse struct {
	List  []FileInfo `json:"list"`
	Total int64      `json:"total"`
}

type FileRequest struct {
	UserId uint `header:"User-Id"`
}

type FileResponse struct {
	Src string `json:"src"`
}

type PageInfo struct {
	Key   string `form:"key,optional"`
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
}

type RequestDelete struct {
	IdList []uint `json:"id_list"`
}

type ShowRequest struct {
	Name string `path:"name"`
}

type ShowResponse struct {
	Url string `json:"url"`
}
