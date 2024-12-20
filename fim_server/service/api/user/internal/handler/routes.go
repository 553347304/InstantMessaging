// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"fim_server/service/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(src *rest.Server, serverCtx *svc.ServiceContext) {
	src.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/add_friend",
				Handler: AddFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/auth",
				Handler: userAuthHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/friend",
				Handler: FriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/friend",
				Handler: FriendNoticeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/user/friend",
				Handler: FriendDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/friend_info",
				Handler: FriendInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/search",
				Handler: searchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/user_info",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/user_info",
				Handler: UserInfoUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/verify",
				Handler: VerifyListHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/verify_status",
				Handler: VerifyStatusHandler(serverCtx),
			},
		},
	)
}
