package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAuthAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupAuthAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAuthAddLogic {
	return &GroupAuthAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupAuthAddLogic) GroupAuthAdd(req *types.GroupAuthAddRequest) (resp *types.GroupAuthAddResponse, err error) {
	// todo: add your logic here and delete this line


	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupId, req.UserId).Error
	if err != nil {
		return nil, logs.Error("请勿重复加群")
	}

	var group group_models.GroupModel
	err = l.svcCtx.DB.Take(&group, "id = ?", req.GroupId).Error
	if err != nil {
		return nil, logs.Error("群不存在")
	}

	resp = new(types.GroupAuthAddResponse)
	resp.Verify = group.Verify
	switch group.Verify {
	case 0:
		return nil, logs.Error("不允许任何人添加")
	case 1:
	// 允许任何人添加
	case 2:
	// 需要验证
	case 3, 4:
		// 需要回答问题   需要正确回答问题
		// if group.AuthQuestion != nil {
		// 	resp.AuthQuestion = types.AuthQuestion{
		// 		Problem1: group.AuthQuestion.Problem1,
		// 		Problem2: group.AuthQuestion.Problem2,
		// 		Problem3: group.AuthQuestion.Problem3,
		// 	}
		// }
	}

	return
}
