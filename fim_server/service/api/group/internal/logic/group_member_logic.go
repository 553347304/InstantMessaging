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
	GroupId        uint   `gorm:"column:group_id"`
	UserID         uint   `gorm:"column:user_id"`
	Role           int8   `gorm:"column:role"`
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

	var UserIDList []uint32
	for _, data := range member.List {
		UserIDList = append(UserIDList, uint32(data.UserID))
	}

	// 关于降级
	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: UserIDList})
	if err != nil {
		logs.Error(err)
	}

	var userInfoMap = map[uint]mtype.UserInfo{}
	for u, info := range userResponse.InfoList {
		userInfoMap[uint(u)] = mtype.UserInfo{
			ID:     uint(u),
			Name:   info.Name,
			Avatar: info.Avatar,
		}
	}

	var userOnlineMap = map[uint]bool{}
	userOnlineResponse, err := l.svcCtx.UserRpc.User.UserOnlineList(l.ctx, &user_rpc.Empty{})
	if err != nil {
		logs.Error(err)
	}
	for _, u := range userOnlineResponse.UserIDList {
		userOnlineMap[uint(u)] = true
	}

	resp = new(types.GroupMemberResponse)
	for _, data := range member.List {
		resp.List = append(resp.List, types.GroupMemberInfo{
			UserID:         data.UserID,
			Name:           userInfoMap[data.UserID].Name,
			Avatar:         userInfoMap[data.UserID].Avatar,
			InOnline:       userOnlineMap[data.UserID],
			Role:           data.Role,
			MemberName:     data.MemberName,
			CreatedAt:      data.CreatedAt,
			NewMessageDate: data.NewMessageDate,
		})
	}
	resp.Total = member.Total
	return
}
