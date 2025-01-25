package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: []uint32{uint32(req.UserId)}})
	if err != nil {
		return nil, logs.Error(err.Error())
	}

	var userConfigModel user_models.UserConfigModel
	conv.Json().Unmarshal(userResponse.Info.UserConfigModel, &userConfigModel)

	resp = &types.UserInfoResponse{
		UserId:        uint(userResponse.Info.Id),
		Name:          userResponse.Info.Name,
		Sign:          userResponse.Info.Sign,
		Avatar:        userResponse.Info.Avatar,
		RecallMessage: userConfigModel.RecallMessage,
		FriendOnline:  userConfigModel.FriendOnline,
		Sound:         userConfigModel.Sound,
		SecureLink:    userConfigModel.SecureLink,
		SavePassword:  userConfigModel.SavePassword,
		SearchUser:    userConfigModel.SearchUser,
		Valid:         userConfigModel.Valid,
		ValidInfo:     conv.Struct(types.ValidInfo{}).Type(userConfigModel.ValidInfo),
	}
	return resp, nil
}
