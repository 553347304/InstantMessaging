package Admin

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/_sys"
	"fim_server/utils/stores/logs"
	"gorm.io/gorm"
	
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
)

type GroupListRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListRemoveLogic {
	return &GroupListRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListRemoveLogic) GroupListRemove(req *types.RequestDelete) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	
	var groupList []group_models.GroupModel
	l.svcCtx.DB.Preload("MemberList").Preload("GroupMessageModel").Find(&groupList, "id in ?", req.IdList)
	
	for _, model := range groupList {
		var validList []group_models.GroupValidModel
		err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
			return _sys.Try(
				func() error {
					return tx.Find(&validList, "group_id = ?", model.ID).Delete(&validList).Error
				},
				func() error {
					if len(model.GroupMessageModel) > 0 {
						return tx.Delete(&model.GroupMessageModel).Error
					}
					return nil
				},
				func() error {
					if len(model.MemberList) > 0 {
						return tx.Delete(&model.MemberList).Error
					}
					return nil
				},
				func() error {
					return tx.Delete(&model).Error
				},
			)
		})
		if err != nil {
			logs.Error(err)
			continue
		}
		logs.InfoF("删除群聊id %d, 群名称 %s", model.ID, model.Name)
		logs.Info("删除群验证总数", len(validList))
		logs.Info("删除群消息总数", len(model.GroupMessageModel))
		logs.Info("删除群用户总数", len(model.MemberList))
	}
	
	return
}
