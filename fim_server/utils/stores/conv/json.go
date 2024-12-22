package conv

import (
	"encoding/json"
	"fim_server/utils/stores/logs"
)

func Unmarshal(data []byte, v any) bool {
	err := json.Unmarshal(data, &v)
	if err != nil {
		logs.Error("Unmarshal conversion error", err.Error())
	}
	return err == nil
}
func Marshal(v any) []byte {
	marshal, err := json.Marshal(v)
	if err != nil {
		logs.Error("Marshal conversion error", err.Error())
		return nil
	}
	return marshal
}
