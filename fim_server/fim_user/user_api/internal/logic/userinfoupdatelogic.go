package logic

import (
	"context"
	"fim_server/common/models"
	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"
	"fim_server/fim_user/user_models"
	"fim_server/utils/stores/converts"
	"fim_server/utils/stores/logs"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserUpdateRequest) (resp *types.UserUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	userMap := converts.StructMap(*req, "user")
	logs.Info(userMap)
	if len(userMap) != 0 {
		var user user_models.User
		err = l.svcCtx.DB.Take(&user, req.UserId).Error
		if err != nil {
			return nil, logs.Error("用户不存在", req.UserId)
		}
		err = l.svcCtx.DB.Model(&user).Updates(userMap).Error
		if err != nil {
			return nil, logs.Error("用户信息更新失败", err)
		}
	}

	userConfigMaps := converts.StructMap(*req, "json")
	logs.Info(userConfigMaps)
	if len(userConfigMaps) != 0 {
		delete(userConfigMaps, "name")
		delete(userConfigMaps, "sign")
		delete(userConfigMaps, "avatar")
		delete(userConfigMaps, "auth_question")
		var userConfig user_models.UserConfig
		err = l.svcCtx.DB.Take(&userConfig, "user_id = ?", req.UserId).Error
		if err != nil {
			return nil, logs.Error("用户配置不存在")
		}

		err = l.svcCtx.DB.Model(&userConfig).Updates(&user_models.UserConfig{AuthQuestion: (*models.AuthQuestion)(req.AuthQuestion)}).Error
		if err != nil {
			return nil, logs.Error("用户配置验证问题更新失败", err)
		}
		err = l.svcCtx.DB.Model(&userConfig).Updates(userConfigMaps).Error
		if err != nil {
			return nil, logs.Error("用户配置更新失败", err)
		}
	}

	return
}
