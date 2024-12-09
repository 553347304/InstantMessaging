package files

import (
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

const (
	White = "white"
	Black = "black"
)

type File struct {
	MaxSize  float64 // MB
	WhiteEXT []string
	BlackEXT []string
}

type FileResponse struct {
	Name  string
	Ext   string
	Size  int64
	Byte  []byte
	Error string
}

// inRoster 判断文件是否在黑白名单中
func inRoster(f File, ext string) string {
	if len(f.WhiteEXT) != 0 {
		if !method.InList(f.WhiteEXT, ext) {
			return fmt.Sprintf("不支持此文件类型 %s/%s", ext, f.WhiteEXT)
		}

	}
	if len(f.BlackEXT) != 0 {
		if method.InList(f.BlackEXT, ext) {
			return fmt.Sprintf("不支持此文件类型 %s/%s", ext, f.BlackEXT)
		}
	}
	return ""
}

// FormFile 上传文件 form-data
func (f File) FormFile(file multipart.File, fileHeader *multipart.FileHeader, err error) FileResponse {
	if err != nil {
		return FileResponse{Error: "FormFile: " + err.Error()}
	}

	fileName := fileHeader.Filename
	fileExt := strings.ToLower(path.Ext(fileName)) // 文件小写扩展名
	sizeMB := float64(fileHeader.Size) / float64(1024*1024)

	// 限制文件大小
	if f.MaxSize != 0 {
		if sizeMB > f.MaxSize {
			return FileResponse{Error: fmt.Sprintf("文件超出限制大小 %0.2f/%0.2fMB", sizeMB, f.MaxSize)}
		}
	}

	// 文件类型黑白名单
	in := inRoster(f, fileExt)
	if in != "" {
		return FileResponse{Error: in}
	}

	// 读取文件
	byteData, err := io.ReadAll(file)
	if err != nil {
		return FileResponse{Error: "io.ReadAll: " + err.Error()}
	}

	return FileResponse{
		Name:  fileName,
		Ext:   fileExt,
		Size:  fileHeader.Size,
		Byte:  byteData,
		Error: "",
	}
}

func ReadFile(filePath string) []byte {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		logs.Error("ReadFile: " + err.Error())
		return nil
	}
	return readFile
}

func IsFileExist(filePath string) bool {
	f, err := os.Open(filePath)
	defer f.Close()
	return err == nil
}
