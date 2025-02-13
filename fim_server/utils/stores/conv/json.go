package conv

import (
	"fim_server/utils/stores/logs"
	"encoding/json"
)

type jsonServerInterface interface {
	Marshal(interface{}) []byte         // 结构体解析Json
	Unmarshal([]byte, interface{}) bool // json解析为结构体 [v:结构体指针]
}
type jsonServer struct{}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Json() jsonServerInterface { return &jsonServer{} }

func (*jsonServer) Marshal(v interface{}) []byte {
	marshal, err := json.Marshal(v)
	if err != nil {
		logs.Error("Marshal conversion error", err.Error())
		return nil
	}
	return marshal
}

func (b *jsonServer) Unmarshal(byte []byte, scan interface{}) bool {
	err := json.Unmarshal(byte, scan)
	if err != nil {
		logs.Error("Unmarshal conversion error", err.Error())
	}
	return err == nil
}
