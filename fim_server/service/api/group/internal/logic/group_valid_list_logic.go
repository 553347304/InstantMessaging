package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupValidListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupValidListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupValidListLogic {
	return &GroupValidListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupValidListLogic) GroupValidList(req *types.GroupValidListRequest) (resp *types.GroupValidListResponse, err error) {
	// todo: add your logic here and delete this line
	
	// 是不是群主/管理员
	var groupIdList []uint
	l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
		Where("user_id = ? and (role = 1 or role = 2)", req.UserId).
		Select("group_id").Scan(&groupIdList)
	
	if len(groupIdList) == 0 {
		return nil, logs.Error("用户不在群内")
	}
	
	groups := src.Mysql(src.ServiceMysql[group_models.GroupValidModel]{
		DB:      l.svcCtx.DB.Where("group_id in ? or user_id = ?", groupIdList, req.UserId),
		Preload: []string{"GroupModel"},
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}).GetList()
	
	// 用户列表
	var UserIDList []uint64
	for _, group := range groups.List {
		UserIDList = append(UserIDList, group.UserId)
	}
	userResponse, err1 := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: UserIDList})
	if err1 != nil {
		return nil, logs.Error("用户服务错误")
	}
	
	resp = new(types.GroupValidListResponse)
	resp.Total = groups.Total
	

	
	for _, group := range groups.List {
		
		var validInfo types.ValidInfo
		method.Struct().To(group.ValidInfo, &validInfo)
		
		resp.List = append(resp.List, types.GroupValidInfo{
			ID:         group.ID,
			UserId:     group.UserId,
			GroupId:    group.GroupId,
			Name:       group.GroupModel.Name,
			Status:     group.Status,
			Valid:      group.Valid,
			ValidInfo:  validInfo,
			Type:       group.Type,
			CreatedAt:  group.CreatedAt.String(),
			Username:   userResponse.InfoList[group.UserId].Username,
			UserAvatar: userResponse.InfoList[group.UserId].Avatar,
		})
	}
	return
}
