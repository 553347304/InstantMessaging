package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type ValidListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidListLogic {
	return &ValidListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidListLogic) ValidList(req *types.FriendValidListRequest) (resp *types.FriendValidListResponse, err error) {
	// todo: add your logic here and delete this line
	
	fvs := src.Mysql(src.ServiceMysql[user_models.FriendValidModel]{
		DB:      l.svcCtx.DB.Where("send_user_id = ? or receive_user_id = ?", req.UserId, req.UserId),
		Preload: []string{"ReceiveUserModel.UserConfigModel", "SendUserModel.UserConfigModel"},
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}).GetList()
	
	var list []types.FriendValidInfo
	for _, fv := range fvs.List {
		info := types.FriendValidInfo{
			ValidMessage: fv.ValidMessage,
			ValidInfo:    conv.Struct(types.ValidInfo{}).Type(fv.ValidInfo),
			Status:       fv.Status,
			Id:           fv.ID,
			CreatedAt:    fv.CreatedAt.String(),
		}

		if fv.SendUserId == req.UserId {
			// 发起方
			info.UserId = fv.SendUserId
			info.Name = fv.SendUserModel.Name
			info.Avatar = fv.SendUserModel.Avatar
			info.Auth = fv.SendUserModel.UserConfigModel.Valid
			info.Status = fv.SendStatus
			info.Flag = "send"
		}
		if fv.ReceiveUserId == req.UserId {
			// 接收方
			info.UserId = fv.ReceiveUserId
			info.Name = fv.ReceiveUserModel.Name
			info.Avatar = fv.ReceiveUserModel.Avatar
			info.Auth = fv.ReceiveUserModel.UserConfigModel.Valid
			info.Status = fv.ReceiveStatus
			info.Flag = "receive"
		}

		list = append(list, info)
	}

	return &types.FriendValidListResponse{List: list, Total: fvs.Total}, nil
}
