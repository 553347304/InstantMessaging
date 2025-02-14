package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberLogic {
	return &GroupMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type mysqlSel struct {
	GroupId        uint64   `gorm:"column:group_id"`
	UserId         uint64   `gorm:"column:user_id"`
	Role           int32   `gorm:"column:role"`
	CreatedAt      string `gorm:"column:created_at"`
	MemberName     string `gorm:"column:member_name"`
	NewMessageDate string `gorm:"column:new_message_date"`
}

func (l *GroupMemberLogic) GroupMember(req *types.GroupMemberRequest) (resp *types.GroupMemberResponse, err error) {
	// todo: add your logic here and delete this line

	if !method.List([]string{"", "role", "created_at"}).InRegex(req.Sort) {
		return nil, logs.Error("不支持的排序模式")
	}
	member := src.Mysql(src.ServiceMysql[mysqlSel]{
		Model: group_models.GroupMemberModel{GroupId: req.Id},
		DB: l.svcCtx.DB.Select("group_id,user_id,role,created_at,member_name,"+
			"(select group_message_models.created_at "+
			"from group_message_models where group_member_models.group_id = ? "+
			"and group_message_models.send_user_id = user_id limit 1) as new_message_date",
			1),
		PageInfo: src.PageInfo{Page: req.Page, Limit: req.Limit, Sort: req.Sort},
	}).GetListGroup()

	var UserIdList []uint64
	for _, data := range member.List {
		UserIdList = append(UserIdList, data.UserId)
	}

	// 关于降级
	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: UserIdList})
	if err != nil {
		logs.Error(err)
	}

	var userInfoMap = map[uint64]mtype.UserInfo{}
	for u, info := range userResponse.InfoList {
		userInfoMap[u] = mtype.UserInfo{
			UserId:     u,
			Username:   info.Username,
			Avatar: info.Avatar,
		}
	}

	var userOnlineMap = map[uint64]bool{}
	userOnlineResponse, err := l.svcCtx.UserRpc.User.UserOnlineList(l.ctx, &user_rpc.Empty{})
	if err != nil {
		logs.Error(err)
	}
	for _, u := range userOnlineResponse.UserIdList {
		userOnlineMap[u] = true
	}

	resp = new(types.GroupMemberResponse)
	for _, data := range member.List {
		resp.List = append(resp.List, types.GroupMemberInfo{
			UserId:         data.UserId,
			Username:           userInfoMap[data.UserId].Username,
			Avatar:         userInfoMap[data.UserId].Avatar,
			InOnline:       userOnlineMap[data.UserId],
			Role:           data.Role,
			MemberName:     data.MemberName,
			CreatedAt:      data.CreatedAt,
			NewMessageDate: data.NewMessageDate,
		})
	}
	resp.Total = member.Total
	return
}
