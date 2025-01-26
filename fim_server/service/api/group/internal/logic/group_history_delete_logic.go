package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupHistoryDeleteLogic {
	return &GroupHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupHistoryDeleteLogic) GroupHistoryDelete(req *types.GroupHistoryDeleteRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.Id, req.UserId).Error
	if err != nil {
		return nil, logs.Error("用户不是群成员", err.Error())
	}

	for _, id := range req.IdList {
		var deleteUserId group_models.GroupMessageModel
		is := l.svcCtx.DB.Take(&deleteUserId, "id = ? and delete_user_id not like ?",
			id, fmt.Sprintf("%%\"%d\"%%", req.UserId)).Error
		// 用户删除列表中没找到就添加
		if is == nil {
			deleteUserId.DeleteUserID = append(deleteUserId.DeleteUserID, fmt.Sprint(req.UserId))
			deleteUserId.DeleteUserID = method.List(deleteUserId.DeleteUserID).Sort(true)
			l.svcCtx.DB.Updates(&deleteUserId)
		}
		logs.Info("删除消息ID", id, is == nil)
	}

	return
}
