package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/src"
	"fmt"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line
	// Online: req.Online
	userConfigs := src.Mysql(src.ServiceMysql[user_models.UserConfigModel]{
		DB: l.svcCtx.DB.Joins("left join user_models um on um.id = user_config_models.user_id").
			Where(""+
				"(user_config_models.search_user <> 0 or user_config_models.search_user is not null) and "+
				"(user_config_models.search_user = 1 and um.id = ?) or "+
				"(user_config_models.search_user = 2 and (um.id = ? or um.name like ?))",
				req.Key, req.Key, fmt.Sprintf("%%%s%%", req.Key)),
		Preload: []string{"UserModel"},
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}).GetList()
	var friend user_models.FriendModel
	friends := friend.MeFriend(l.svcCtx.DB, req.UserID)
	userMap := map[uint]bool{}
	for _, model := range friends {
		if model.SendUserID == req.UserID {
			userMap[model.ReceiveUserID] = true
		} else {
			userMap[model.SendUserID] = true
		}
	}

	list := make([]types.SearchInfo, 0)
	for _, userConfig := range userConfigs.List {
		list = append(list, types.SearchInfo{
			UserID:   userConfig.UserID,
			Name:     userConfig.UserModel.Name,
			Sign:     userConfig.UserModel.Sign,
			Avatar:   userConfig.UserModel.Avatar,
			IsFriend: userMap[userConfig.UserID],
		})
	}
	return &types.SearchResponse{List: list, Total: userConfigs.Total}, nil
}
