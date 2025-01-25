package usercurtaillogic

import (
	"context"
	"fim_server/models/user_models"
	"fmt"

	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"

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
		tip := "用户不存在" + fmt.Sprint(in.Id)
		resp.CurtailChat.Error = tip
		resp.CurtailAddUser.Error = tip
		resp.CurtailCreateGroup.Error = tip
		resp.CurtailAddGroup.Error = tip
		return resp, err
	}

	resp.CurtailChat.Is = user.CurtailChat
	resp.CurtailChat.Error = "当前用户被限制聊天"
	resp.CurtailAddUser.Is = user.CurtailAddUser
	resp.CurtailAddUser.Error = "当前用户被限制加好友"
	resp.CurtailCreateGroup.Is = user.CurtailCreateGroup
	resp.CurtailCreateGroup.Error = "当前用户被限制创建群聊"
	resp.CurtailAddGroup.Is = user.CurtailAddGroup
	resp.CurtailAddGroup.Error = "当前用户被限制加群"

	return resp, nil
}
