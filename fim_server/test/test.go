package main

import (
	"fim_server/config/core"
	"fim_server/models"
	"fim_server/utils/src"
)

func main() {
	core.Init()

	src.DB.Create(&models.Test{
		Json: []string{"Reading", "Coding", "Traveling"},
	})

}
