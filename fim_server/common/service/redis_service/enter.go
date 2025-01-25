package redis_service

import (
	"context"
	"fim_server/models/mtype"
	"fim_server/service/rpc/user/client"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/conv"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func GetUserInfo(rdb *redis.Client, UserRpc client.UserRpc, userId uint) (userInfo mtype.UserInfo, err error) {
	key := fmt.Sprint("user_info_", userId)
	str, err1 := rdb.Get(key).Result()
	if err1 != nil {
		// 用户信息
		userResponse, err5 := UserRpc.User.UserInfo(context.Background(), &user_rpc.IdList{Id: []uint32{uint32(userId)}})
		if err5 != nil {
			return userInfo, err5
		}
		userInfo.ID = userId
		userInfo.Avatar = userResponse.Info.Avatar
		userInfo.Name = userResponse.Info.Name

		data := conv.Json().Marshal(userInfo)
		rdb.Set(key, string(data), time.Hour) // 1个小时过期
		return
	}
	conv.Json().Unmarshal([]byte(str), &userInfo)
	return
}
