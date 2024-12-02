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
	var authModel = user_models.FriendAuthModel{
		SendUserId:    req.UserId,
		ReceiveUserId: req.FriendId,
		AuthMessage:   req.AuthMessage,
	}
	logs.Info(req.FriendId)
	logs.Info(userConfig.Auth)
	switch userConfig.Auth {
	case 0:
		return nil, logs.Error("不允许任何人添加")
	case 1:
		// 允许任何人添加
		authModel.ReceiveStatus = 1
		break
	case 2:
	// 需要验证
	case 3:
		// 需要回答问题

		if req.AuthQuestion != nil {
			authModel.AuthQuestion = &models.AuthQuestion{
				Problem1: req.AuthQuestion.Problem1,
				Problem2: req.AuthQuestion.Problem2,
				Problem3: req.AuthQuestion.Problem3,
				Answer1:  req.AuthQuestion.Answer1,
				Answer2:  req.AuthQuestion.Answer2,
				Answer3:  req.AuthQuestion.Answer3,
			}
		}
	case 4:
		// 需要正确回答问题
		if req.AuthQuestion != nil && userConfig.AuthQuestion != nil {
			// 判断问题是否回答正确
			// 考虑到一个问题，两个问题，三个问题的情况
			var count int
			if userConfig.AuthQuestion.Answer1 != nil && req.AuthQuestion.Answer1 != nil {
				if *userConfig.AuthQuestion.Answer1 == *req.AuthQuestion.Answer1 {
					count += 1
				}
			}
			if userConfig.AuthQuestion.Answer2 != nil && req.AuthQuestion.Answer2 != nil {
				if *userConfig.AuthQuestion.Answer2 == *req.AuthQuestion.Answer2 {
					count += 1
				}
			}
			if userConfig.AuthQuestion.Answer3 != nil && req.AuthQuestion.Answer3 != nil {
				if *userConfig.AuthQuestion.Answer3 == *req.AuthQuestion.Answer3 {
					count += 1
				}
			}
			if count != userConfig.ProblemCount() {
				return nil, logs.Error("答案错误")
			}
			// 直接加好友
			authModel.ReceiveStatus = 1
			authModel.AuthQuestion = userConfig.AuthQuestion

		}
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
