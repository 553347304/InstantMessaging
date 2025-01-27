package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupSessionLogic {
	return &GroupSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Data struct {
	GroupId string `gorm:"column:group_id"`
	NewDate string `gorm:"column:new_date"`
	Preview string `gorm:"column:preview"`
}

func (l *GroupSessionLogic) GroupSession(req *types.GroupSessionRequest) (resp *types.GroupSessionResponse, err error) {
	// todo: add your logic here and delete this line
	response := src.Mysql(src.ServiceMysql[Data]{
		Model: group_models.GroupMessageModel{},
		DB: l.svcCtx.DB.Select("group_id,max(created_at),"+
			"(select preview from group_message_models as g "+
			"where g.group_id = g.group_id order by g.created_at desc limit 1)as new_date").
			Where("group_id in (?)", req.UserID).
			Group("group_id"),
	}).GetListGroup()

	resp = new(types.GroupSessionResponse)
	logs.Info(response)
	return
}
