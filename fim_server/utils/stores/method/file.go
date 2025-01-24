package method

import (
	"fim_server/utils/stores/logs"
	"os"
)

type fileServerInterface interface {
	Delete()
}
type fileServer struct{ Path string }

func File(path string) fileServerInterface { return &fileServer{Path: path} }

func (f *fileServer) Delete() {
	err := os.Remove(f.Path)
	if err != nil {
		logs.Error("error ->", err, "路径", f.Path)
	}
}
