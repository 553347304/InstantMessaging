package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthStatusLogic {
	return &AuthStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthStatusLogic) AuthStatus(req *types.FriendAuthStatusRequest) (resp *types.FriendAuthStatusResponse, err error) {
	// todo: add your logic here and delete this line

	logs.Info("aa")

	var friendAuth user_models.FriendAuthModel
	err = l.svcCtx.DB.Take(&friendAuth, "id = ? and receive_user_id = ?", req.AuthId, req.UserId).Error
	if err != nil {
		return nil, logs.Error("验证记录不存在")
	}
	if friendAuth.Status != 0 {
		return nil, logs.Error("不可更改状态")
	}

	switch req.Status {
	case 1:
		// 同意
		friendAuth.ReceiveStatus = 1
		l.svcCtx.DB.Create(&user_models.FriendModel{
			SendUserId:    friendAuth.SendUserId,
			ReceiveUserId: friendAuth.ReceiveUserId,
		})
	case 2:
		// 拒绝
		friendAuth.ReceiveStatus = 2
	case 3:
		// 忽略
		friendAuth.ReceiveStatus = 3
	case 4: // 删除
		friendAuth.Status = 4
		l.svcCtx.DB.Delete(&friendAuth)
		return nil, nil
	}
	l.svcCtx.DB.Save(&friendAuth)
	return
}