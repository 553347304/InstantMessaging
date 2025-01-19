package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"
	"fmt"
	"time"
	
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupBanUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupBanUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupBanUpdateLogic {
	return &GroupBanUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupBanUpdateLogic) GroupBanUpdate(req *types.GroupBanUpdateRequest) (resp *types.GroupBanUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	
	var member group_models.GroupMemberModel
	l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupId, req.UserId)
	if !(member.Role == 1 || member.Role == 2) {
		return nil, logs.Error("权限不足")
	}
	
	var member1 group_models.GroupMemberModel
	l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.GroupId, req.MemberId)
	if member.Role > member1.Role {
		return nil, logs.Error("权限不足")
	}
	
	l.svcCtx.DB.Model(&member1).Update("ban_time", req.BanTime)
	// 禁言时间
	key := fmt.Sprintf("ban_time__%d", member1.ID)
	if req.BanTime != 0 {
		l.svcCtx.Redis.Set(key, "1", time.Duration(req.BanTime)*time.Minute)
	}
	
	return
}
