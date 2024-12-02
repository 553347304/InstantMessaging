package main

import (
	"fim_server/config/core"
	"fim_server/models/chat_models"
	"fim_server/utils/src"
	"fim_server/utils/src/sqls"
	"fmt"
)

func main() {
	core.Init()

	type Data struct {
		SU         uint   `gorm:"column:s_u"`
		RU         uint   `gorm:"column:r_u"`
		MaxDate    string `gorm:"column:max_date"`
		MaxPreview string `gorm:"column:max_preview"`
	}

	var chatList []Data
	sqls.GetListGroup(chat_models.ChatModel{}, &chatList, sqls.Mysql{
		DB: src.DB.
			Select("least(send_user_id, receive_user_id) as s_u",
				"greatest(send_user_id, receive_user_id) as r_u",
				" max(created_at)   as max_date",
				"max(message_preview) as max_preview").
			Where("send_user_id = ? or receive_user_id = ?", 1, 1).
			Group("least(send_user_id, receive_user_id)").
			Group("greatest(send_user_id, receive_user_id)"),
		PageInfo: src.PageInfo{
			Sort:  "max_date desc",
			Page:  1,
			Limit: 10,
		},
	})

	fmt.Println(chatList)
}
