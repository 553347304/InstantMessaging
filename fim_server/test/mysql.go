package main

import (
	"fim_server/config"
	"fim_server/utils/stores/logs"
)

type Content struct {
	Type    int8   `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Size struct {
	Type    int8   `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Size    string `json:"size"`
}

func main() {
	
	config.Init()
	
	var scan map[string]any
	
	// s1 := conv.Json().Marshal([]Size{
	// 	{Size: "coontnent"},
	// 	{Size: "coontnent"},
	// 	{Size: "coontnent"},
	// })
	// config.DB.Create(&models.Test{
	// 	Message: s1,
	// })
	
	
	
	config.DB.Find(&scan)
	
	
	
	logs.Info(scan)
	
	
}

func find(scan interface{}) {
	config.DB.Find(scan)
}
