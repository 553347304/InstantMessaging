package redis_service

import (
	"context"
	"fim_server/models/mtype"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/conv"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// GetUserInfo 获取用户信息
func GetUserInfo(rdb *redis.Client, UserRpc user_rpc.UserClient, userId uint) (userInfo mtype.UserInfo, err error) {
	key := fmt.Sprint("user_info_", userId)
	str, err1 := rdb.Get(key).Result()
	if err1 != nil {
		// 用户信息
		info, err5 := UserRpc.UserBaseInfo(context.Background(), &user_rpc.UserBaseInfoRequest{
			UserId: uint32(userId),
		})
		if err5 != nil {

			return userInfo, err5
		}
		userInfo.ID = userId
		userInfo.Avatar = info.Avatar
		userInfo.Name = info.Name

		data := conv.Marshal(userInfo)
		rdb.Set(key, string(data), time.Hour) // 1个小时过期
		return
	}
	conv.Unmarshal([]byte(str), &userInfo)
	return
}
