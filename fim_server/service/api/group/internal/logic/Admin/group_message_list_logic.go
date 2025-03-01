package Admin

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"fim_server/utils/stores/method"
	"time"
	
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
)

type GroupMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMessageListLogic {
	return &GroupMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HistoryResponse struct {
	ID         uint64          `json:"id"`
	UserId     uint64          `json:"user_id"`
	UserAvatar string        `json:"user_avatar"`
	CreatedAt  time.Time     `json:"created_at"`
	MemberId   uint64          `json:"member_id"`
	MemberName string        `json:"member_name"`
	IsMe       bool          `json:"is_me"`
	Message    mtype.Message `json:"message"`
}

func (l *GroupMessageListLogic) GroupMessageList(req *types.PageInfo) (resp *response.List[HistoryResponse], err error) {
	// todo: add your logic here and delete this line
	
	groupMessageList := src.Mysql(src.ServiceMysql[group_models.GroupMessageModel]{
		DB: l.svcCtx.DB,
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
		Preload: []string{"MemberModel"},
	}).GetList()
	
	var UserIDList []uint64
	for _, model := range groupMessageList.List {
		UserIDList = append(UserIDList, model.SendUserId)
	}
	
	UserIDList = method.List(UserIDList).Unique() // 去重
	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: UserIDList})
	if err != nil {
		return nil, err
	}
	
	var list = make([]HistoryResponse, 0)
	for _, info := range groupMessageList.List {
		list = append(list, HistoryResponse{
			ID:         info.ID,
			UserId:     info.SendUserId,
			UserAvatar: userResponse.InfoList[info.SendUserId].Avatar,
			CreatedAt:  info.CreatedAt,
			MemberId:   info.MemberId,
			MemberName: userResponse.InfoList[info.SendUserId].Username,
			Message:    info.Message,
		})
	}
	
	resp = new(response.List[HistoryResponse])
	resp.List = list
	resp.Total = groupMessageList.Total
	
	return
}
