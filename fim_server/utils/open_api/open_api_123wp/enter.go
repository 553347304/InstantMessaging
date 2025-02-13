package open_api_123wp

import (
	"mime/multipart"
)

type URL struct {
	Upload        string // 上传文件
	GetUploadUrl  string // 获取上传url
	UploadSucceed string // 上传完毕
	FileList      string // 文件列表
	GetURL        string // 获取直链
	Mkdir         string // 创建目录
}

type trayServiceInterface interface {
	Upload(*multipart.FileHeader, uint) error       // 上传文件   文件 | 文件夹ID
	List(uint, int, string) ([]fileResponse, error) // 文件列表   文件夹ID | 页数
	GetURL(uint) (string, error)                    // 获取直链   文件ID
	Mkdir(uint64, string) (uint64, error)           // 创建目录   父级ID | 文件名
}
type trayService struct {
	url          URL
	header       map[string]string
	root         uint64 // 父级目录ID
	preuploadID  string // 预上传ID
	presignedURL string // 上传地址
}

func Tray() trayServiceInterface {
	const base = "https://open-api.123pan.com"
	const root = 12218728
	return &trayService{
		url: URL{
			Upload:        base + "/upload/v1/file/create",
			GetUploadUrl:  base + "/upload/v1/file/get_upload_url",
			UploadSucceed: base + "/upload/v1/file/upload_complete",
			FileList:      base + "/api/v1/file/list",
			GetURL:        base + "/api/v1/direct-link/url",
			Mkdir:         base + "/upload/v1/file/mkdir",
		},
		header: map[string]string{
			"Platform":      "open_platform",
			"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk2Nzk4MjEsImlhdCI6MTczOTA3NTAyMSwiaWQiOjE4MjE1NjAyNDYsIm1haWwiOiIiLCJuaWNrbmFtZSI6IueZvemfs34iLCJ1c2VybmFtZSI6MTM5NDU0MjQxMzEsInYiOjB9._1PY_vXCADw5VhmHu5LhI9dxlhvnPe1P5m_RGlKAb4k",
		},
	}
}

type baseResponse[T interface{}] struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Data     T      `json:"data"`
	XTraceID string `json:"x-traceID"`
}
