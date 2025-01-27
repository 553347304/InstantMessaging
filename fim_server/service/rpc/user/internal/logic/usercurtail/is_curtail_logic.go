package usercurtaillogic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	"fmt"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type IsCurtailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsCurtailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsCurtailLogic {
	return &IsCurtailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsCurtailLogic) IsCurtail(in *user_rpc.ID) (*user_rpc.CurtailResponse, error) {
	// todo: add your logic here and delete this line

	resp := new(user_rpc.CurtailResponse)

	var user user_models.UserConfigModel
	err := l.svcCtx.DB.Take(&user, "user_id = ?", in.Id).Error
	if err != nil {
		return nil, logs.Error("用户不存在" + fmt.Sprint(in.Id))
	}
	
	if !user.CurtailChat {
		resp.CurtailChat = "当前用户被限制聊天"
	}
	if !user.CurtailAddUser {
		resp.CurtailAddUser = "当前用户被限制加好友"
	}
	if !user.CurtailCreateGroup {
		resp.CurtailCreateGroup = "当前用户被限制创建群聊"
	}
	if !user.CurtailAddGroup {
		resp.CurtailAddGroup = "当前用户被限制加群"
	}
	
	return resp, nil
}
