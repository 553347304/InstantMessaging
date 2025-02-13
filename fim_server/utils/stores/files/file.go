package files

import (
	"fim_server/utils/stores/logs"
	"os"
	"os/exec"
	"strings"
)

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

// RootDir 获取根目录
func RootDir() string {
	var dir, _ = os.Getwd()
	return strings.ReplaceAll(dir, "\\", "/")
}

// RootDirName 获取当前根目录名
func RootDirName() string {
	cmd := exec.Command("go", "list", "-m")
	output, _ := cmd.CombinedOutput()
	return strings.TrimSpace(string(output))
}
