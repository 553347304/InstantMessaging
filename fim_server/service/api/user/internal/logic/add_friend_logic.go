package logic

import (
	"context"
	"fim_server/models"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/conv"
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

	is, err1 := l.svcCtx.UserRpc.Curtail.IsCurtail(l.ctx, &user_rpc.ID{Id: uint32(req.UserId)})
	if err1 != nil || !is.CurtailAddUser.Is {
		return nil, conv.Type(is.CurtailAddUser.Error).Error()
	}

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
	var validModel = user_models.FriendValidModel{
		SendUserId:    req.UserId,
		ReceiveUserId: req.FriendId,
		SendStatus:    1,
		ValidMessage:  req.ValidMessage,
		ValidInfo:     conv.Struct(models.ValidInfo{}).Type(req.ValidInfo),
	}

	logs.Info(req.FriendId)
	logs.Info(userConfig.Valid)

	switch userConfig.Valid {
	case 0:
		return nil, logs.Error("不允许任何人添加")
	case 1:
		validModel.ReceiveStatus = 1 // 允许任何人添加
		break
	case 2:
		validModel.ReceiveStatus = 0 // 需要验证
		return
	case 3:
		validModel.ReceiveStatus = 0 // 需要回答问题
		return
	case 4:
		// 需要正确回答问题
		if !userConfig.ValidInfo.Valid(req.ValidInfo.Answer) {
			return nil, logs.Error("答案错误")
		}
		validModel.ReceiveStatus = 1 // 直接加好友
	}

	err = l.svcCtx.DB.Create(&validModel).Error

	// 加好友
	if validModel.ReceiveStatus != 1 {
		return
	}
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
