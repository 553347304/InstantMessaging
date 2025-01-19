package group_models

import (
	"fim_server/models"
	"fim_server/utils/src"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"time"
)

// GroupMemberModel 群成员表
type GroupMemberModel struct {
	models.Model
	GroupId    uint       `json:"groupId"` // 群ID
	GroupModel GroupModel `gorm:"foreignKey:GroupId" json:"-"`
	UserId     uint       `json:"user_id"`                    // 用户ID
	MemberName string     `gorm:"size:32" json:"member_name"` // 群名称
	Role       int8       `json:"role"`                       // 1 群主 2 管理员 3 普通成员
	BanTime    *int       `json:"ban_time"`                   // 禁言时间 单位分钟
}

func (g GroupMemberModel) GetBanTime(db *gorm.DB, rdb *redis.Client) *int {
	
	if g.BanTime == nil {
		return nil
	}
	key := fmt.Sprintf("ban_time__%d", g.ID)
	// 是否过期
	t := src.Redis(rdb).IsExpiration(key)
	if t < 0 {
		rdb.Del(key)
		db.Model(&g).Update("ban_time", nil)
		return nil
	}
	res := int(t / time.Minute)
	return &res
}
