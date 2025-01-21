package main

import (
	"fim_server/models"
	"fim_server/utils/src"
)

func main() {
	var aaa models.Test
	src.Mysql(src.MysqlServer[aaa]{
		Model: models.Test{},
	}).GetList()
	
}
