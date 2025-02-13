package files

import (
	"mime/multipart"
)

type Config struct {
	Header  *multipart.FileHeader // 文件
	White   []string              // 白名单
	MaxSize *int64                // 限制大小
}

type uploadConfigResponse struct {
	Name string `json:"name"` // 文件名
	Ext  string `json:"ext"`  // 文件扩展名
	Size int64  `json:"size"` // 文件大小 单位/字节
	Md5  string `json:"md5"`
	Byte []byte `json:"byte"`
}