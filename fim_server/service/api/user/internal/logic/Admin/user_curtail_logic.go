package Admin

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
)

type UserCurtailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCurtailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCurtailLogic {
	return &UserCurtailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCurtailLogic) UserCurtail(req *types.UserCurtailRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	var user user_models.UserModel
	err = l.svcCtx.DB.Preload("UserConfigModel").Take(&user, req.UserID).Error
	if err != nil {
		return nil, logs.Error("用户不存在")
	}
	l.svcCtx.DB.Model(&user.UserConfigModel).Updates(map[string]any{
		"curtail_chat":         req.CurtailChat,
		"curtail_add_user":     req.CurtailAddUser,
		"curtail_create_group": req.CurtailCreateGroup,
		"curtail_add_group":    req.CurtailAddGroup,
	})
	return
}
