package flags

import (
	"fim_server/config"
	"fim_server/models/chat_models"
	"fim_server/models/file_models"
	"fim_server/models/group_models"
	"fim_server/models/log_models"
	"fim_server/models/setting_models"
	"fim_server/models/user_models"
)

func MigrationTable() error {
	// global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	return config.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&user_models.UserModel{},
			&user_models.UserConfigModel{},

			&user_models.FriendModel{},
			&user_models.FriendValidModel{},

			&chat_models.ChatModel{},

			&group_models.GroupModel{},
			&group_models.GroupMemberModel{},
			&group_models.GroupMessageModel{},
			&group_models.GroupValidModel{},

			&file_models.FileModel{},

			&log_models.LogModel{},

			&setting_models.ConfigModel{},

			 // &models.Test{},
		)
}
