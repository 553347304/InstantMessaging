package files

import (
	"errors"
	"fim_server/utils/encryption_and_decryptio/md5s"
	"fim_server/utils/stores/algorithms"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/randoms"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type File struct {
	R        *http.Request
	Key      string
	MaxSize  *float64 // MB
	WhiteEXT *[]string
}

type FileResponse struct {
	Name  string
	Size  float64
	Byte  []byte
	Error error
}

// FormFile 上传文件 form-data
func FormFile(f File) FileResponse {
	file, fileHead, err := f.R.FormFile(f.Key)
	if err != nil {
		return FileResponse{Error: err}
	}

	// 限制文件大小
	sizeMB := float64(fileHead.Size) / float64(1024*1024)
	if f.MaxSize != nil {
		if sizeMB > *f.MaxSize {
			return FileResponse{Error: errors.New(fmt.Sprintf("文件超出限制大小 %0.2f/%0.2fMB", sizeMB, *f.MaxSize))}
		}
	}

	// 限制文件类型
	if f.WhiteEXT != nil {
		fileName := fileHead.Filename
		fileExt := path.Ext(fileName)
		lowerStr := strings.ToLower(fileExt)
		is := algorithms.InList(*f.WhiteEXT, lowerStr)
		if !is {
			return FileResponse{Error: errors.New(fmt.Sprintf("不支持此文件类型 %s/%s", lowerStr, *f.WhiteEXT))}
		}
	}
	byteData, err := io.ReadAll(file)
	if err != nil {
		return FileResponse{Error: errors.New("io.ReadAll: " + err.Error())}
	}
	return FileResponse{
		Name:  fileHead.Filename,
		Size:  sizeMB,
		Byte:  byteData,
		Error: nil,
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

func WriteFile(filePath string, byte []byte) string {
	if IsFileExist(filePath) {
		if md5s.Check(byte, ReadFile(filePath)) {
			ext := path.Ext(filePath)
			filePath = strings.Replace(filePath, ext, randoms.String(1)+ext, -1)
			logs.Info("图片已存在 随机名字", filePath)
			return WriteFile(filePath, byte)
		}
	}
	err := os.WriteFile(filePath, byte, 0666)
	if err != nil {
		logs.Error("WriteFile: " + err.Error())
		return ""
	}
	return "/" + filePath
}
