package flags

import (
	"fim_server/models/chat_models"
	"fim_server/models/file_models"
	"fim_server/models/group_models"
	"fim_server/models/user_models"
	"fim_server/utils/src"
)

func MigrationTable() error {
	// global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	return src.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&user_models.UserModel{},
			&user_models.FriendModel{},
			&user_models.FriendAuthModel{},
			&user_models.UserConfigModel{},

			&chat_models.ChatModel{},
			&chat_models.TopUserModel{},
			&chat_models.UserChatDeleteModels{},

			&group_models.GroupModel{},
			&group_models.GroupMemberModel{},
			&group_models.GroupMessageModel{},
			&group_models.GroupAuthModel{},

			&file_models.FileModel{},
		)
}
