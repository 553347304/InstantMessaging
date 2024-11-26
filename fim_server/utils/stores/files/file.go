package files

import (
	"os"
	"os/exec"
	"strings"
)

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
