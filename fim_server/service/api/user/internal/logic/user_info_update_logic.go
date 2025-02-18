package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	
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
	
	if req.UserInfo != nil {
		userInfoMap := method.Struct().ToMap(req.UserInfo)
		var user user_models.UserModel
		err = l.svcCtx.DB.Take(&user, req.UserId).Error
		if err != nil {
			return nil, logs.Error("用户不存在", req.UserId)
		}
		err = l.svcCtx.DB.Model(&user).Updates(userInfoMap).Error
		if err != nil {
			return nil, logs.Error("用户信息更新失败", err)
		}
	}
	
	if req.UserConfig != nil {
		userConfigMap := method.Struct().ToMap(req.UserConfig)
		userConfigMap["valid_info"] = conv.Json().Marshal(req.UserConfig.ValidInfo)
		if req.UserConfig.ValidInfo.Issue == nil || req.UserConfig.ValidInfo.Answer == nil {
			delete(userConfigMap, "valid_info")
		}
		var userConfig user_models.UserConfigModel
		err = l.svcCtx.DB.Take(&userConfig, "user_id = ?", req.UserId).Error
		if err != nil {
			return nil, logs.Error("用户配置不存在", req.UserId)
		}
		err = l.svcCtx.DB.Model(&userConfig).Updates(userConfigMap).Error
		if err != nil {
			return nil, logs.Error("用户配置更新失败", err)
		}
	}
	
	return
}
