package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/utils/src"
	"fim_server/utils/src/sqls"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyListLogic {
	return &VerifyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyListLogic) VerifyList(req *types.FriendVerifyListRequest) (resp *types.FriendVerifyListResponse, err error) {
	// todo: add your logic here and delete this line

	fvs := sqls.GetList(user_models.FriendAuthModel{}, sqls.Mysql{
		DB:      l.svcCtx.DB.Where("send_user_id = ? or receive_user_id = ?", req.UserId, req.UserId),
		Preload: []string{"ReceiveUserModel.UserConfigModel", "SendUserModel.UserConfigModel"},
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})

	var list []types.FriendVerifyInfo
	for _, fv := range fvs.List {
		info := types.FriendVerifyInfo{
			VerifyMessage: fv.VerifyMessage,
			VerifyInfo:    types.VerifyInfo{Issue: fv.VerifyInfo.Issue, Answer: fv.VerifyInfo.Answer},
			Status:        fv.Status,
			Id:            fv.ID,
		}

		if fv.SendUserId == req.UserId {
			// 发起方
			info.UserId = fv.SendUserId
			info.Name = fv.SendUserModel.Name
			info.Avatar = fv.SendUserModel.Avatar
			info.Auth = fv.SendUserModel.UserConfigModel.Verify
			info.Status = fv.SendStatus
			info.Flag = "send"
		}
		if fv.ReceiveUserId == req.UserId {
			// 接收方
			info.UserId = fv.ReceiveUserId
			info.Name = fv.ReceiveUserModel.Name
			info.Avatar = fv.ReceiveUserModel.Avatar
			info.Auth = fv.ReceiveUserModel.UserConfigModel.Verify
			info.Status = fv.ReceiveStatus
			info.Flag = "receive"
		}

		list = append(list, info)
	}

	return &types.FriendVerifyListResponse{List: list, Total: fvs.Total}, nil
}
