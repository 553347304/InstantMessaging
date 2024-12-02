package logic

import (
	"context"
	"encoding/json"
	"fim_server/models/user_models"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"

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
	var user user_models.UserModel
	err = json.Unmarshal(result.Data, &user)
	if err != nil {
		return nil, logs.Error("数据错误")
	}

	if user.UserConfigModel == nil {
		return nil, logs.Error("当前用户ID没有配置", user.ID)
	}

	resp = &types.UserInfoResponse{
		UserId:        user.ID,
		Name:          user.Name,
		Sign:          user.Sign,
		Avatar:        user.Avatar,
		RecallMessage: user.UserConfigModel.RecallMessage,
		FriendOnline:  user.UserConfigModel.FriendOnline,
		Sound:         user.UserConfigModel.Sound,
		SecureLink:    user.UserConfigModel.SecureLink,
		SavePassword:  user.UserConfigModel.SavePassword,
		SearchUser:    user.UserConfigModel.SearchUser,
		Auth:          user.UserConfigModel.Auth,
	}
	if user.UserConfigModel.AuthQuestion != nil {
		resp.AuthQuestion = &types.AuthQuestion{
			Problem1: user.UserConfigModel.AuthQuestion.Problem1,
			Problem2: user.UserConfigModel.AuthQuestion.Problem2,
			Problem3: user.UserConfigModel.AuthQuestion.Problem3,
			Answer1:  user.UserConfigModel.AuthQuestion.Answer1,
			Answer2:  user.UserConfigModel.AuthQuestion.Answer2,
			Answer3:  user.UserConfigModel.AuthQuestion.Answer3,
		}
	}
	return resp, nil
}
