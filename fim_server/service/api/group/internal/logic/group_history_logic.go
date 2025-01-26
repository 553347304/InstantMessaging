package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	"time"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
)

type GroupHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupHistoryLogic {
	return &GroupHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HistoryResponse struct {
	ID         uint               `json:"id"`
	UserId     uint               `json:"user_id"`
	UserAvatar string             `json:"user_avatar"`
	CreatedAt  time.Time          `json:"created_at"`
	MemberId   uint               `json:"member_id"`
	MemberName string             `json:"member_name"`
	IsMe       bool               `json:"is_me"`
	Message    mtype.Message `json:"message"`
}

// MemberList 查全部群成员
func (l *GroupHistoryLogic) MemberList(id uint) map[uint]group_models.GroupMemberModel {
	var memberList []group_models.GroupMemberModel
	var memberMap = map[uint]group_models.GroupMemberModel{}
	l.svcCtx.DB.Find(&memberList, "group_id = ?", id)
	for _, model := range memberList {
		memberMap[model.UserId] = model
	}
	return memberMap
}
func (l *GroupHistoryLogic) GroupHistory(req *types.GroupHistoryRequest) (resp *response.List[HistoryResponse], err error) {
	// todo: add your logic here and delete this line

	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.Id, req.UserId).Error
	if err != nil {
		return nil, logs.Error("用户不是群成员", err.Error())
	}

	groupMessageList := src.Mysql(src.ServiceMysql[group_models.GroupMessageModel]{
		DB: l.svcCtx.DB.Where("group_id = ? and delete_user_id not like ?",
			req.Id, fmt.Sprintf("%%\"%d\"%%", req.UserId)), // 查询删除的用户ID  "id"
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
		Preload: []string{"MemberModel"},
	}).GetList()

	var userIdList []uint32
	for _, model := range groupMessageList.List {
		userIdList = append(userIdList, uint32(model.SendUserId))
	}

	userIdList = method.List(userIdList).Unique() // 去重
	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: userIdList})
	if err != nil {
		return nil, err
	}

	memberMap := l.MemberList(req.Id) // 查成员列表
	var list = make([]HistoryResponse, 0)
	for _, info := range groupMessageList.List {
		// 群备注名称

		memberName := memberMap[info.SendUserId].MemberName
		if memberName == "" {
			memberName = userResponse.InfoList[uint32(info.SendUserId)].Name
		}
		list = append(list, HistoryResponse{
			ID:         info.ID,
			UserId:     info.SendUserId,
			UserAvatar: userResponse.InfoList[uint32(info.SendUserId)].Avatar,
			CreatedAt:  info.CreatedAt,
			MemberId:   info.MemberId,
			MemberName: memberName,
			IsMe:       info.ID == req.UserId,
			Message:    info.Message,
		})
	}

	resp = new(response.List[HistoryResponse])
	resp.List = list
	resp.Total = groupMessageList.Total
	return
}
