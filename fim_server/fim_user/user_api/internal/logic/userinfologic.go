package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fim_server/fim_user/user_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils/stores/logs"

	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"

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
	result, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: uint32(req.UserId),
	})
	if err != nil {
		return nil, logs.Error(err.Error())
	}
	var user user_models.User
	err = json.Unmarshal(result.Data, &user)
	if err != nil {
		return nil, errors.New("数据错误")
	}
	if user.UserConfig == nil {
		return nil, logs.Error("当前用户ID没有配置", user.ID)
	}

	resp = &types.UserInfoResponse{
		UserId:        user.ID,
		Name:          user.Name,
		Sign:          user.Sign,
		Avatar:        user.Avatar,
		RecallMessage: user.UserConfig.RecallMessage,
		FriendOnline:  user.UserConfig.FriendOnline,
		Sound:         user.UserConfig.Sound,
		SecureLink:    user.UserConfig.SecureLink,
		SavePassword:  user.UserConfig.SavePassword,
		SearchUser:    user.UserConfig.SearchUser,
		Auth:          user.UserConfig.Auth,
	}
	if user.UserConfig.AuthQuestion != nil {
		resp.AuthQuestion = &types.AuthQuestion{
			Problem1: user.UserConfig.AuthQuestion.Problem1,
			Problem2: user.UserConfig.AuthQuestion.Problem2,
			Problem3: user.UserConfig.AuthQuestion.Problem3,
			Answer1:  user.UserConfig.AuthQuestion.Answer1,
			Answer2:  user.UserConfig.AuthQuestion.Answer2,
			Answer3:  user.UserConfig.AuthQuestion.Answer3,
		}
	}
	return resp, nil
}
