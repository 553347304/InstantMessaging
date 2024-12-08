package files

import (
	"fim_server/utils/stores/logs"
	"os"
	"path/filepath"
)

// type writeFileResponse struct {
// 	Path  string
// 	Hash  string
// 	Error error
// }
//
// type Write struct {
// }
//
// func (w Write) File(filePath string, byte []byte) writeFileResponse {
// 	nowHash := md5s.Hash(byte) // 新文件哈希
// 	if IsFileExist(filePath) {
// 		// 检查文件内容是否相同
// 		oldHash := md5s.Hash(ReadFile(filePath)) // 旧文件哈希
// 		if oldHash == nowHash {
// 			ext := path.Ext(filePath)
// 			newPath := strings.Replace(filePath, ext, randoms.String(5)+ext, -1) // 替换新名字
// 			logs.Info("原始名字: ", filePath, "随机名字: ", newPath)
// 			return w.File(newPath, byte)
// 		}
// 	}
// 	err := os.WriteFile(filePath, byte, 0666)
// 	if err != nil {
// 		return writeFileResponse{Error: logs.Error("WriteFile: " + err.Error())}
// 	}
// 	return writeFileResponse{
// 		Path:  "/" + filePath,
// 		Hash:  nowHash,
// 		Error: nil,
// 	}
// }

// MkdirAll 递归创建文件夹
func MkdirAll(filePath string) error {
	dir := filepath.Dir(filePath) // 获取文件所在的目录
	// 递归创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logs.Info(dir)
		if os.MkdirAll(dir, 0755) != nil {
			return err
		}
	}
	return nil
}

func WriteFile(filePath string, data []byte) error {

	err := MkdirAll(filePath)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
