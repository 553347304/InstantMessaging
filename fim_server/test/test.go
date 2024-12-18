package main

import (
	"encoding/json"
	"fim_server/utils/stores/logs"
	"fmt"
)
type Example struct {
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

func main() {
	example := Example{
		Name:  "",
		Value: 1,
	}

	jsonData, err := json.Marshal(example)
	if err != nil {
		fmt.Println(err)
		return
	}
	logs.Info(string(jsonData))
}
