package logic

import (
	"context"
	"encoding/json"
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
	result, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{
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
		Valid:         user.UserConfigModel.Valid,
		ValidInfo:     conv.Struct(types.ValidInfo{}).Type(user.UserConfigModel.ValidInfo),
	}
	return resp, nil
}
