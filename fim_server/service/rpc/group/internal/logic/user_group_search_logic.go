package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGroupSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGroupSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGroupSearchLogic {
	return &UserGroupSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type Data struct {
	UserId uint32 `gorm:"column:user_id"`
	Count  uint32 `gorm:"column:count"`
}

func (l *UserGroupSearchLogic) UserGroupCreateTotal(in *group_rpc.UserGroupSearchRequest, scan *[]Data) {
	l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
		Where("user_id in ?", in.UserIDList).
		Group("user_id").
		Select("user_id", "count(id) as count").Scan(&scan)
}
func (l *UserGroupSearchLogic) UserGroupAddTotal(in *group_rpc.UserGroupSearchRequest, scan *[]Data) {
	l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
		Where("user_id in ? and role = ?", in.UserIDList, 1).
		Group("user_id").
		Select("user_id", "count(id) as count").Scan(&scan)
}

func (l *UserGroupSearchLogic) UserGroupSearch(in *group_rpc.UserGroupSearchRequest) (resp *group_rpc.UserGroupSearchResponse, err error) {
	// todo: add your logic here and delete this line
	logs.Info("----------")
	var data []Data
	if in.Mode == 1 {
		l.UserGroupCreateTotal(in, &data)
	}
	if in.Mode == 2 {
		l.UserGroupAddTotal(in, &data)
	}
	var groupUserMap = map[uint32]uint32{}
	for _, u2 := range data {
		groupUserMap[u2.UserId] = u2.Count
	}
	resp = new(group_rpc.UserGroupSearchResponse)
	resp.Result = map[uint32]int32{}
	for _, uid := range in.UserIDList {
		resp.Result[uid] = int32(groupUserMap[uid])
	}

	logs.Info(data)
	return resp, nil
}
