package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoLogic {
	return &GroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupInfoLogic) GroupInfo(req *types.GroupInfoRequest) (resp *types.GroupInfoResponse, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.GroupRpc.IsInGroupMember(context.Background(), &group_rpc.IsInGroupMemberRequest{
		UserId:  uint32(req.UserId),
		GroupId: uint32(req.Id),
	})
	if err != nil {
		return nil, err
	}

	var groupModel group_models.GroupModel
	err = l.svcCtx.DB.Preload("MemberList").Take(&groupModel, req.Id).Error
	if err != nil {
		return nil, logs.Error("群不存在")
	}
	resp = &types.GroupInfoResponse{
		GroupId:     groupModel.ID,
		Name:        groupModel.Name,
		Sign:        groupModel.Sign,
		MemberCount: len(groupModel.MemberList),
		Avatar:      groupModel.Avatar,
	}

	// 查用户列表信息
	var userIdList []uint32
	var userAllIdList []uint32
	for _, model := range groupModel.MemberList {
		if model.Role == 1 || model.Role == 2 {
			userIdList = append(userIdList, uint32(model.UserId))
		}
		userAllIdList = append(userAllIdList, uint32(model.UserId))
	}

	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIdList,
	})
	if err != nil {
		return nil, err
	}
	var leader types.UserInfo
	var adminList = make([]types.UserInfo, 0)
	for _, model := range groupModel.MemberList {
		if model.Role == 3 {
			continue
		}
		userInfo := types.UserInfo{
			UserId: model.UserId,
			Avatar: userListResponse.UserInfo[uint32(model.UserId)].Avatar,
			Name:   userListResponse.UserInfo[uint32(model.UserId)].Name,
		}
		if model.Role == 1 {
			leader = userInfo
			continue
		}
		if model.Role == 2 {
			adminList = append(adminList, userInfo)
		}
	}

	resp.Leader = leader
	resp.AdminList = adminList

	// 用户在线总数
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(context.Background(), &user_rpc.UserOnlineListRequest{})
	if err == nil {
		slice := method.ListIntersect(userOnlineResponse.UserIdList, userAllIdList)
		resp.MemberOnlinCount = len(slice)
	}

	return
}
