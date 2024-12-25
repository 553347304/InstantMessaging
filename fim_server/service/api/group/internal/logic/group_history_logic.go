package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/mtype"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"fim_server/utils/src/sqls"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	"time"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupHistoryLogic {
	return &GroupHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HistoryResponse struct {
	ID          uint              `json:"id"`
	UserId      uint              `json:"user_id"`
	UserName    string            `json:"user_name"`
	UserAvatar  string            `json:"user_avatar"`
	MessageType mtype.MessageType `json:"message_type"`
	Message     mtype.Message     `json:"message"`
	CreatedAt   time.Time         `json:"created_at"`
}

func (l *GroupHistoryLogic) GroupHistory(req *types.GroupHistoryRequest) (resp *response.List[HistoryResponse], err error) {
	// todo: add your logic here and delete this line

	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.Id, req.UserId).Error
	if err != nil {
		return nil, logs.Error("用户不是群成员", err.Error())
	}

	groupMessageList := sqls.GetList(group_models.GroupMessageModel{},
		sqls.Mysql{
			DB: l.svcCtx.DB.Where("group_id = ? and delete_user_id not like ?",
				req.Id, fmt.Sprintf("%%\"%d\"%%", req.UserId)),	// 查询删除的用户ID  "id"
			PageInfo: src.PageInfo{
				Page:  req.Page,
				Limit: req.Limit,
			},
		})

	var userIdList []uint32
	for _, model := range groupMessageList.List {
		userIdList = append(userIdList, uint32(model.SendUserId))
	}

	userIdList = method.List(userIdList).Unique() // 去重
	userInfoList, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{UserIdList: userIdList})
	if err != nil {
		return nil, err
	}
	var list = make([]HistoryResponse, 0)
	for _, info := range groupMessageList.List {
		list = append(list, HistoryResponse{
			ID:          info.ID,
			UserId:      info.SendUserId,
			UserName:    userInfoList.UserInfo[uint32(req.UserId)].Name,
			UserAvatar:  userInfoList.UserInfo[uint32(req.UserId)].Avatar,
			MessageType: info.MessageType,
			Message:     info.Message,
			CreatedAt:   info.CreatedAt,
		})
	}

	resp = new(response.List[HistoryResponse])
	resp.List = list
	resp.Total = groupMessageList.Total
	return
}
