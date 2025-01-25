package Admin

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/rpc/chat/chat_rpc"
	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"github.com/zeromicro/go-zero/core/logx"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserOnlineList() map[uint]bool {
	var userOnlineMap = map[uint]bool{}
	userOnlineResponse, err1 := l.svcCtx.UserRpc.User.UserOnlineList(l.ctx, &user_rpc.Empty{})
	if err1 == nil {
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		logx.Error(err1)
	}
	return userOnlineMap
}
func (l *UserListLogic) UserList(req *types.RequestPageInfo) (resp *response.List[types.UserListInfoResponse], err error) {
	// todo: add your logic here and delete this line

	result := src.Mysql(src.ServiceMysql[user_models.UserModel]{
		DB:      l.svcCtx.DB,
		Preload: []string{"UserConfigModel"},
		Where:   "name like ? or ip like ?",
		PageInfo: src.PageInfo{
			Key:   req.Key,
			Limit: req.Limit,
			Page:  req.Page,
		},
	}).GetList()
	resp = new(response.List[types.UserListInfoResponse])
	var userIdList []uint32
	for _, model := range result.List {
		userIdList = append(userIdList, uint32(model.ID))
	}

	userOnlineMap := l.UserOnlineList()

	userGroupCreateTotal, _ := l.svcCtx.GroupRpc.UserGroupSearch(l.ctx, &group_rpc.UserGroupSearchRequest{UserIdList: userIdList, Mode: 1})
	userGroupAddTotal, _ := l.svcCtx.GroupRpc.UserGroupSearch(l.ctx, &group_rpc.UserGroupSearchRequest{UserIdList: userIdList, Mode: 2})
	chatResponse, _ := l.svcCtx.ChatRpc.UserListChatTotal(l.ctx, &chat_rpc.UserListChatTotalRequest{UserIdList: userIdList})

	for _, model := range result.List {
		info := types.UserListInfoResponse{
			ID:                 model.ID,
			CreatedAt:          model.CreatedAt.String(),
			Name:               model.Name,
			Avatar:             model.Avatar,
			IP:                 model.IP,
			Addr:               model.Addr,
			IsOnline:           userOnlineMap[model.ID],
			GroupAdminCount:    int(userGroupCreateTotal.Result[uint32(model.ID)]),
			GroupCount:         int(userGroupAddTotal.Result[uint32(model.ID)]),
			CurtailChat:        model.UserConfigModel.CurtailChat,
			CurtailAddUser:     model.UserConfigModel.CurtailAddUser,
			CurtailCreateGroup: model.UserConfigModel.CurtailCreateGroup,
			CurtailAddGroup:    model.UserConfigModel.CurtailAddGroup,
		}
		if chatResponse.Result[uint32(model.ID)] != nil {
			info.SendMsgCount = int(chatResponse.Result[uint32(model.ID)].SendMessageTotal)
		}

		resp.List = append(resp.List, info)
	}

	return
}
