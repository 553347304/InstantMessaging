package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/stores/method"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupSearchLogic {
	return &GroupSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupSearchLogic) GroupSearch(req *types.GroupSearchListRequest) (resp *types.GroupSearchListResponse, err error) {
	// todo: add your logic here and delete this line

	// IsSearch 是否搜索
	groupListResponse := src.Mysql(src.ServiceMysql[group_models.GroupModel]{
		DB:      l.svcCtx.DB.Where("is_search = 1 and (id = ? or name like ?)", req.Id, "%"+req.Key+"%"),
		Preload: []string{"MemberList"},
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}).GetList()

	// 用户在线总数
	var userOnlineIdList []uint
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(l.ctx, &user_rpc.UserOnlineListRequest{})
	if err == nil {
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineIdList = append(userOnlineIdList, uint(u))
		}
	}

	resp = new(types.GroupSearchListResponse)
	for _, group := range groupListResponse.List {

		// 用户在线总数
		var isInGroup bool // 是否在群
		var groupMemberIdList []uint
		for _, model := range group.MemberList {
			groupMemberIdList = append(groupMemberIdList, model.UserId)
			if model.UserId == req.UserId {
				isInGroup = true
			}
		}
		total := method.List(groupMemberIdList).Intersect(userOnlineIdList)

		resp.List = append(resp.List, types.GroupSearchInfo{
			GroupId:         group.ID,
			Name:            group.Name,
			Sign:            group.Sign,
			Avatar:          group.Avatar,
			UserCount:       len(group.MemberList),
			UserOnlineCount: len(total),
			IsInGroup:       isInGroup,
		})
	}
	resp.Total = groupListResponse.Total
	return
}
