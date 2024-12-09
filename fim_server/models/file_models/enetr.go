package file_models

import (
	"fim_server/models"
	"github.com/google/uuid"
)

type FileModel struct {
	models.Model
	Uid    uuid.UUID `json:"uid"`     // 文件唯一ID  /api/file/{uid}
	UserId uint      `json:"user_id"` // 用户id
	Name   string    `json:"name"`    // 文件名称
	Size   int64     `json:"size"`    // 文件大小
	Path   string    `json:"path"`    // 文件路径
	Hash   string    `json:"hash"`    // 文件哈希
}

func (file *FileModel) WebPath() string {
	return "/api/file/" + file.Uid.String()
}
