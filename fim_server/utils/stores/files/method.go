package files

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/valid"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Merge(index int, filePath string) error {
	// 保存新文件
	newFile, err := os.Create(filePath)
	if err != nil {
		return logs.Error(err)
	}
	defer newFile.Close()
	
	// 合并分片文件
	for i := 0; i < index; i++ {
		openPath := fmt.Sprintf("%s_%d", filePath, i)
		openFile, err := os.Open(openPath)
		if err != nil {
			return logs.Error(err)
		}
		_, err = io.Copy(newFile, openFile)
		
		openFile.Close()
		Delete(openPath)
		
		if err != nil {
			return logs.Error(err)
		}
	}
	
	logs.InfoF("合并分片:%d", index)
	return nil
}

func Mkdir(filePath string) {
	dir := filepath.Dir(filePath) // 获取文件所在的目录
	// 递归创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if os.MkdirAll(dir, 0755) != nil {
			logs.Warn("文件夹创建失败", err.Error())
		}
	}
}

func Delete(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		logs.Warn("删除路径失败" + err.Error())
	}
}

func Write(data []byte, dir string) error {
	Mkdir(dir)
	err := os.WriteFile(dir, data, 0666)
	if err != nil {
		return conv.Type("文件存入失败" + err.Error()).Error()
	}
	
	return nil
}

func Read(file multipart.File, size int64) ([]byte, error) {
	// 分片读
	var progress int64
	buf := make([]byte, 0)
	for {
		readBuf := make([]byte, 1024)
		n, err := file.Read(readBuf)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return nil, err
		}
		
		buf = append(buf, readBuf[:n]...)
		
		progress += int64(n)
		logs.Progress(progress, size, "#") // 打印进度
	}
	
	return buf, nil
}

func Upload(c Config) (*uploadConfigResponse, error) {
	name := c.Header.Filename                           // 文件名
	size := c.Header.Size                               // 文件大小
	ext := strings.ToLower(path.Ext(c.Header.Filename)) // 文件小写扩展名
	
	// 限制文件大小
	if c.MaxSize != nil && size > *c.MaxSize {
		return nil, conv.Type(fmt.Sprintf("%s 超出限制大小 %d/%d", name, size, *c.MaxSize)).Error()
	}
	
	// 白名单
	for i := 0; i < len(c.White); i++ {
		if c.White[i] != ext {
			return nil, conv.Type(fmt.Sprintf("%s  类型不支持%s/%s", name, ext, c.White)).Error()
		}
	}
	
	// 读文件
	file, _ := c.Header.Open()
	data, err := Read(file, size)
	if err != nil {
		return nil, err
	}
	
	return &uploadConfigResponse{
		Name: name,
		Ext:  ext,
		Size: size,
		Md5:  valid.MD5().Hash(string(data)),
		Byte: data,
	}, nil
}
