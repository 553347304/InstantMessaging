package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	"strings"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateRequest) (resp *types.GroupCreateResponse, err error) {
	// todo: add your logic here and delete this line

	var groupModel = group_models.GroupModel{
		Leader:   req.UserId,
		IsSearch: false,
		Valid:    2,
		Size:     50,
		Sign:     fmt.Sprintf("本群创建于%s  群主很聪明,什么都没有留下", method.Time().Now),
	}

	is, _ := l.svcCtx.UserRpc.Curtail.IsCurtail(l.ctx, &user_rpc.ID{Id: req.UserId})
	if is != nil && is.CurtailCreateGroup != "" {
		return nil, conv.Type(is.CurtailCreateGroup).Error()
	}
	

	var groupUserList = []uint64{req.UserId}
	switch req.Mode {
	case 1:
		if req.Name == "" || req.Size >= 1000 {
			return nil, logs.Error("群名不能为空")
		}
		groupModel.Name = req.Name
		groupModel.Size = req.Size
		groupModel.IsSearch = req.IsSearch
	case 2:
		if len(req.UserIDList) == 0 {
			return nil, logs.Error("没有选择的好友")
		}

		var UserIDList = []uint64{req.UserId} // 先把自己放进去
		for _, u := range req.UserIDList {
			UserIDList = append(UserIDList, u)
			groupUserList = append(groupUserList, u)
		}
		
		UserIDList = method.List(UserIDList).Unique()
		userResponse, err2 := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: UserIDList})
		if err2 != nil {
			return nil, logs.Error(err2, UserIDList)
		}

		var nameList []string
		for _, info := range userResponse.InfoList {
			if len(strings.Join(nameList, ".")) >= 29 {
				break
			}
			nameList = append(nameList, info.Username)
		}

		groupModel.Name = strings.Join(nameList, "、") + "的群聊"

		userFriendList, err1 := l.svcCtx.UserRpc.Friend.FriendList(l.ctx, &user_rpc.ID{Id: req.UserId})
		if err1 != nil {
			return nil, logs.Error(err1)
		}
		var friendList []uint64
		for _, i2 := range userFriendList.FriendList {
			friendList = append(friendList, i2.Id)
		}
		slice := method.List(req.UserIDList).Difference(friendList)
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
