package logic

import (
	"context"
	"fim_server/models"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/utils/stores/logs"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendRequest) (resp *types.AddFriendResponse, err error) {
	// todo: add your logic here and delete this line

	var friend user_models.FriendModel
	if friend.IsFriend(l.svcCtx.DB, req.UserId, req.FriendId) {
		return nil, logs.Error("已经是好友了")
	}

	var userConfig user_models.UserConfigModel
	err = l.svcCtx.DB.Take(&userConfig, "user_id = ?", req.FriendId).Error
	if err != nil {
		return nil, logs.Error("用户不存在")
	}
	resp = new(types.AddFriendResponse)


	// 创建验证消息
	var verifyModel = user_models.FriendAuthModel{
		SendUserId:    req.UserId,
		ReceiveUserId: req.FriendId,
		SendStatus:    1,
		VerifyMessage: req.VerifyMessage,
		VerifyInfo: models.VerifyInfo{
			Issue:  req.VerifyInfo.Issue,
			Answer: req.VerifyInfo.Answer,
		},
	}

	logs.Info(req.FriendId)
	logs.Info(userConfig.Verify)

	switch userConfig.Verify {
	case 0:
		return nil, logs.Error("不允许任何人添加")
	case 1:
		verifyModel.ReceiveStatus = 1 // 允许任何人添加
		break
	case 2:
		verifyModel.ReceiveStatus = 0 // 需要验证
		err = l.svcCtx.DB.Create(&verifyModel).Error
		return
	case 3:
		verifyModel.ReceiveStatus = 0 // 需要回答问题
		err = l.svcCtx.DB.Create(&verifyModel).Error
		return
	case 4:
		// 需要正确回答问题
		if !userConfig.VerifyInfo.Verify(req.VerifyInfo.Answer) {
			return nil, logs.Error("答案错误")
		}
		verifyModel.ReceiveStatus = 1 // 直接加好友
	}

	// 加好友
	var userFriend = user_models.FriendModel{
		SendUserId:    req.UserId,
		ReceiveUserId: req.FriendId,
	}
	err = l.svcCtx.DB.Create(&userFriend).Error
	if err != nil {
		return nil, logs.Error("添加好友失败")
	}

	return
}
