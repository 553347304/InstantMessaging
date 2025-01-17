package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/conv"
	
	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type UserListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListInfoLogic {
	return &UserListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListInfoLogic) UserListInfo(in *user_rpc.UserListInfoRequest) (*user_rpc.UserListInfoResponse, error) {
	// todo: add your logic here and delete this line
	
	err := conv.Type("用户不存在").Error()
	
	var userList []user_models.UserModel
	l.svcCtx.DB.Find(&userList, in.UserIdList)
	
	resp := new(user_rpc.UserListInfoResponse)
	
	resp.UserInfo = make(map[uint32]*user_rpc.UserInfo)
	for _, user := range userList {
		err = nil
		resp.UserInfo[uint32(user.ID)] = &user_rpc.UserInfo{
			Name:   user.Name,
			Avatar: user.Avatar,
		}
	}
	
	return resp, err
}
