package main

import (
	"fim_server/config/core"
)

func main() {

	core.Init()

	// src.DB.Create(&models.Test{
	// 	String: "66",
	// 	AuthQuestion: models.AuthQuestion{
	// 		Issue:  []string{"What is your name?", "What is your age?"},
	// 	},
	// })
	// var cr models.Test
	// src.DB.Take(&cr, 26)
	// logs.Info(cr)
	// logs.Info(cr.AuthQuestion.Answer)
	// logs.Info(cr.AuthQuestion)
}
