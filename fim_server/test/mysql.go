package main

import (
	"fim_server/config/core"
	"fim_server/models/group_models"
	"fim_server/utils/src"
)

func main() {

	core.Init()
	var aa group_models.GroupMessageModel
	src.DB.Take(&aa, 6)
	src.DB.Model(&aa).Update("delete_user_id", "[]")

	// src.DB.Create(&models.Test{
	// 	String: "66",
	// 	Arr: []string{"What is your name?", "What is your age?"},
	// })
	// var cr models.Test
	// src.DB.Take(&cr, 10)
	// logs.Info(cr)

}
