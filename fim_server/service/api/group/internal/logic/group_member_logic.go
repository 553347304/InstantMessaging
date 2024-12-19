package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/src/sqls"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method/method_list"
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
	UserId         uint   `gorm:"column:user_id"`
	Role           int8   `gorm:"column:role"`
	CreatedAt      string `gorm:"column:created_at"`
	MemberName     string `gorm:"column:member_name"`
	NewMessageDate string `gorm:"column:new_message_date"`
}

func (l *GroupMemberLogic) GroupMember(req *types.GroupMemberRequest) (resp *types.GroupMemberResponse, err error) {
	// todo: add your logic here and delete this line

	if !method_list.InRegex([]string{"", "role", "created_at"}, req.Sort) {
		return nil, logs.Error("不支持的排序模式")
	}

	member := sqls.GetListGroup(group_models.GroupMemberModel{GroupId: req.Id}, sqls.Mysql{
		DB: l.svcCtx.DB.Select("group_id,user_id,role,created_at,member_name,"+
			"(select group_message_models.created_at "+
			"from group_message_models where group_member_models.group_id = ? "+
			"and group_message_models.send_user_id = user_id limit 1) as new_message_date",
			1),
		PageInfo: src.PageInfo{Page: req.Page, Limit: req.Limit, Sort: req.Sort},
	}, &[]mysqlSel{})
	var userIdList []uint32
	for _, data := range member.List {
		userIdList = append(userIdList, uint32(data.UserId))
	}

	// 关于降级
	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{UserIdList: userIdList})
	if err != nil {
		logs.Error(err)
	}

	var userInfoMap = map[uint]mtype.UserInfo{}
	for u, info := range userListResponse.UserInfo {
		userInfoMap[uint(u)] = mtype.UserInfo{
			ID:     uint(u),
			Name:   info.Name,
			Avatar: info.Avatar,
		}
	}

	userListResponse.UserInfo = map[uint32]*user_rpc.UserInfo{}

	var userOnlineMap = map[uint]bool{}
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(context.Background(), &user_rpc.UserOnlineListRequest{})
	if err != nil {
		logs.Error(err)
	}
	for _, u := range userOnlineResponse.UserIdList {
		userOnlineMap[uint(u)] = true
	}

	resp = new(types.GroupMemberResponse)
	for _, data := range member.List {
		resp.List = append(resp.List, types.GroupMemberInfo{
			UserId:         data.UserId,
			Name:           userInfoMap[data.UserId].Name,
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
