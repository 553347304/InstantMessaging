package logic

import (
	"context"
	"fim_server/models"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAddLogic {
	return &GroupAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupAddLogic) GroupAdd(req *types.GroupAddRequest) (resp *types.GroupAddResponse, err error) {
	// todo: add your logic here and delete this line

	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupId, req.UserId).Error
	if err == nil {
		return nil, logs.Error("请勿重复加群")
	}

	var group group_models.GroupModel
	err = l.svcCtx.DB.Take(&group, "id = ?", req.GroupId).Error
	if err != nil {
		return nil, logs.Error("群不存在")
	}

	resp = new(types.GroupAddResponse)

	var verifyModel = group_models.GroupAuthModel{
		GroupId:    req.GroupId,
		UserId:     req.UserId,
		Verify:     group.Verify,
		VerifyInfo: models.VerifyInfo{Issue: req.VerifyInfo.Issue, Answer: req.VerifyInfo.Answer},
		Type:       1,
	}

	logs.Info(group.VerifyInfo)
	logs.Info(req.VerifyInfo.Answer)
	switch group.Verify {
	case 0:
		return nil, logs.Error("不允许任何人添加")
	case 1:
		verifyModel.Status = 1 // 直接加群
		l.svcCtx.DB.Create(&verifyModel)
	// 允许任何人添加
	case 2:
		verifyModel.Status = 2 // 需要验证
	case 3:
		verifyModel.Status = 2 // 需要验证
	case 4:
		if !group.VerifyInfo.Verify(req.VerifyInfo.Answer) {
			return nil, logs.Error("答案错误")
		}
		verifyModel.Status = 1 // 直接加群
	}

	err = l.svcCtx.DB.Create(&verifyModel).Error
	if err != nil { return nil, err }
	
	// 加群
	if verifyModel.Status != 1 {
		return
	}
	var groupMember = group_models.GroupMemberModel{
		GroupId: req.GroupId,
		UserId:  req.UserId,
		Role:    3,
	}
	l.svcCtx.DB.Create(&groupMember)

	return
}
