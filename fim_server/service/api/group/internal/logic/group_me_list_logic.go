package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMeListLogic {
	return &GroupMeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMeListLogic) GroupMeList(req *types.GroupMeListRequest) (resp *types.GroupMeListResponse, err error) {
	// todo: add your logic here and delete this line
	var groupIdList []uint
	// 我加入的群聊
	query := l.svcCtx.DB.Model(&group_models.GroupMemberModel{}).Where("user_id = ?", req.UserId)
	if req.Mode == 1 {
		query.Where("role = ?", 1) // 我创建的群聊
	}
	query.Select("group_id").Scan(&groupIdList)
	
	var pageInfo src.PageInfo
	method.Struct().To(req.PageInfo, &pageInfo)
	groups := src.Mysql(src.ServiceMysql[group_models.GroupModel]{
		DB:       l.svcCtx.DB.Where("id in ?", groupIdList),
		PageInfo: pageInfo,
	}).GetList()

	resp = new(types.GroupMeListResponse)
	for _, model := range groups.List {
		var role int32
		for _, memberModel := range model.MemberList {
			if memberModel.UserId == req.UserId {
				role = memberModel.Role
			}
		}
		resp.List = append(resp.List, types.GroupMeInfo{
			Id:          model.ID,
			Name:        model.Name,
			Avatar:      model.Avatar,
			MemberTatal: int64(len(model.MemberList)),
			Role:        role,
			Mode:        req.Mode,
		})
	}

	logs.Info(groups)

	return
}
