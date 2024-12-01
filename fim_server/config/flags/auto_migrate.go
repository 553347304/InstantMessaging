package flags

import (
	"fim_server/models/chat_models"
	"fim_server/models/group_models"
	"fim_server/models/user_models"
	"fim_server/utils/service/src"
	"fmt"
)

func MigrationTable() {
	var err error
	// global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	err = src.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&user_models.UserModel{},
			&user_models.FriendModel{},
			&user_models.FriendAuthModel{},
			&user_models.UserConfigModel{},

			&chat_models.ChatModel{},

			&group_models.GroupModel{},
			&group_models.GroupMemberModel{},
			&group_models.GroupMessageModel{},
			&group_models.GroupAuthModel{},
		)
	if err != nil {
		fmt.Println("[生成数据库表结构失败]")
		return
	}
	fmt.Println("[生成数据库表结构成功]")
}
