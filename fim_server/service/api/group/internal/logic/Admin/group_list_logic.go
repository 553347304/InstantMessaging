package Admin

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
)

type GroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) UserIdList(groupResponseList []group_models.GroupModel) []uint32 {
	var userIdList []uint
	for _, model := range groupResponseList {
		for _, memberModel := range model.MemberList {
			userIdList = append(userIdList, memberModel.UserId)
		}
	}
	newUserIdList := method.List(userIdList).Unique()
	return conv.Slice(newUserIdList).Uint32()
}

func (l *GroupListLogic) UserOnlineMap() map[uint]bool {
	var userOnlineMap = map[uint]bool{}
	userOnlineResponse, err := l.svcCtx.UserRpc.User.UserOnlineList(l.ctx, &user_rpc.Empty{})
	if err == nil {
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		logs.Info(err)
	}
	return userOnlineMap
}

func (l *GroupListLogic) GroupList(req *types.PageInfo) (resp *types.GroupListResponse, err error) {
	// todo: add your logic here and delete this line
	
	groupResponse := src.Mysql(src.ServiceMysql[group_models.GroupModel]{
		DB:       l.svcCtx.DB,
		Where:    "name like ?",
		Preload:  []string{"MemberList", "GroupMessageModel"},
		PageInfo: src.PageInfo{Page: req.Page, Limit: req.Limit, Key: req.Key, Sort: "created_at desc"},
	}).GetList()
	
	userIdList := l.UserIdList(groupResponse.List)
	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: userIdList})
	if err != nil {
		return nil, err
	}
	userOnlineMap := l.UserOnlineMap()
	
	// logs.Info(userResponse)
	
	resp = new(types.GroupListResponse)
	
	for _, g := range groupResponse.List {
		info := types.GroupListInfoResponse{
			ID:           g.ID,
			CreatedAt:    g.CreatedAt.String(),
			Name:         g.Name,
			Sign:         g.Sign,
			Avatar:       g.Avatar,
			MemberTotal:  len(g.MemberList),
			MessageTotal: len(g.GroupMessageModel),
			Leader: types.UserInfo{
				UserId: g.Leader,
				Avatar: userResponse.InfoList[uint32(g.Leader)].Avatar,
				Name:   userResponse.InfoList[uint32(g.Leader)].Name,
			},
		}
		
		var adminList []types.UserInfo
		
		for _, m := range g.MemberList {
			_, ok := userOnlineMap[m.UserId]
			if ok {
				info.MemberOnlineTotal++
			}
			if m.Role == 2 {
				adminList = append(adminList, types.UserInfo{
					UserId: m.UserId,
					Avatar: userResponse.InfoList[uint32(m.UserId)].Avatar,
					Name:   userResponse.InfoList[uint32(m.UserId)].Name,
				})
			}
		}
		info.AdminList = adminList
		
		resp.List = append(resp.List, info)
	}
	
	return
}
