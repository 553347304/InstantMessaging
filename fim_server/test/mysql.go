package main

import (
	"fim_server/config/core"
	"fim_server/models"
	"fim_server/utils/src"
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

	core.Init()
	
	src.DB.Create(&models.Test{
		Arr: []models.Content{
			{Title: "13", Content: "你好你好", Size: 100},
			{Title: "22"},
			{Title: "33"},
		},
	})
	var cr models.Test
	src.DB.Take(&cr, 0)
	logs.Info(cr.Arr[0].Title)
	
}
