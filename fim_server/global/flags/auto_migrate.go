package flags

import (
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_group/group_models"
	"fim_server/fim_user/user_models"
	"fim_server/global"
	"fmt"
)

func MigrationTable() {
	var err error
	// global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
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
