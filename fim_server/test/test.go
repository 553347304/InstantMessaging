package main

import (
	"fim_server/models"
	"fim_server/utils/stores/conv"
	"fmt"
)

type GroupModel struct {
	models.Model
	ValidInfo []models.ValidInfo `json:"valid_info"` // 验证问题
}

func main() {
	r := []GroupModel{
		{ValidInfo: []models.ValidInfo{
			{Issue: []string{"issue"},
				Answer: []string{"answer"}}},
		},
	}
	
	fmt.Println(conv.Struct(r).StructSliceMap("id","valid_info"))
	
}
