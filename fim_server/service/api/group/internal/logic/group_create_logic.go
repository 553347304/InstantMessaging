package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fim_server/utils/stores/times"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateRequest) (resp *types.GroupCreateResponse, err error) {
	// todo: add your logic here and delete this line

	var groupModel = group_models.GroupModel{
		Leader:      req.UserId,
		IsSearch:    false,
		Valid: 			2,
		Size:        50,
		Sign:        fmt.Sprintf("本群创建于%s  群主很聪明,什么都没有留下", times.Now()),
	}

	var groupUserList = []uint{req.UserId}
	switch req.Mode {
	case 1:
		if req.Name == "" || req.Size >= 1000 {
			return nil, logs.Error("群名不能为空")
		}
		groupModel.Name = req.Name
		groupModel.Size = req.Size
		groupModel.IsSearch = req.IsSearch
	case 2:
		if len(req.UserIdList) == 0 {
			return nil, logs.Error("没有选择的好友")
		}

		var userIdList = []uint32{uint32(req.UserId)} // 先把自己放进去
		for _, u := range req.UserIdList {
			userIdList = append(userIdList, uint32(u))
			groupUserList = append(groupUserList, u)
		}

		userListResponse, err2 := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
			UserIdList: userIdList,
		})
		if err2 != nil {
			return nil, logs.Error(err2)
		}

		var nameList []string
		for _, info := range userListResponse.UserInfo {
			if len(strings.Join(nameList, ".")) >= 29 {
				break
			}
			nameList = append(nameList, info.Name)
		}

		groupModel.Name = strings.Join(nameList, "、") + "的群聊"

		userFriendList, err1 := l.svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
			UserId: uint32(req.UserId),
		})
		if err1 != nil {
			return nil, logs.Error(err1)
		}
		var friendList []uint
		for _, i2 := range userFriendList.FriendList {
			friendList = append(friendList, uint(i2.UserId))
		}
		slice := method.List(req.UserIdList).Difference(friendList)
		if len(slice) != 0 {
			return nil, logs.Error("列表中有人不是好友")
		}

	default:
		return nil, logs.Error("不支持的模式")
	}

	groupModel.Avatar = string([]rune(groupModel.Name)[0])
	err = l.svcCtx.DB.Create(&groupModel).Error

	if err != nil {
		return nil, logs.Error("群创建失败")
	}
	var members []group_models.GroupMemberModel
	for i, u := range groupUserList {
		memBerModel := group_models.GroupMemberModel{
			GroupId: groupModel.ID,
			UserId:  u,
			Role:    3,
		}
		if i == 0 {
			memBerModel.Role = 1
		}
		members = append(members, memBerModel)
	}
	err = l.svcCtx.DB.Create(&members).Error

	if err != nil {
		return nil, logs.Error("群成员添加失败")
	}


	return
}
