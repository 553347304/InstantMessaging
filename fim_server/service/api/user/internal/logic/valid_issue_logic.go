package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidIssueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidIssueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidIssueLogic {
	return &ValidIssueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidIssueLogic) ValidIssue(req *types.ValidIssueRequest) (resp *types.ValidIssueResponse, err error) {
	// todo: add your logic here and delete this line

	var friend user_models.FriendModel
	if friend.IsFriend(l.svcCtx.DB, req.UserId, req.Id) {
		return nil, logs.Error("已经是好友了")
	}

	var userConfig user_models.UserConfigModel
	err = l.svcCtx.DB.Take(&userConfig, "user_id = ?", req.Id).Error
	if err != nil {
		return nil, logs.Error("用户不存在")
	}
	resp = new(types.ValidIssueResponse)
	resp.Valid = userConfig.Valid
	resp.ValidInfo = conv.Struct(types.ValidInfo{}).Type(userConfig.ValidInfo)
	return
}
