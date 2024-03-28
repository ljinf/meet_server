package cache

import (
	"encoding/json"
	"fmt"
	"github.com/ljinf/meet_server/internal/model"
	"time"
)

const (
	Expired             = 1800
	UserInfoCachePrefix = "cache:user:info:"
)

type AccountCache interface {
	SetUserInfoCache(info *model.UserInfo) error
	GetUserInfoCache(userId int64) (*model.UserInfo, error)
	DelUserInfoCache(userId int64) error
}

type accountCache struct {
	*Cache
}

func NewAccountCache(c *Cache) AccountCache {
	return &accountCache{
		Cache: c,
	}
}

func (a *accountCache) SetUserInfoCache(info *model.UserInfo) error {
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v%v", UserInfoCachePrefix, info.UserId)
	return a.rdb.Set(key, string(v), time.Duration(Expired)).Err()
}

func (a *accountCache) GetUserInfoCache(userId int64) (*model.UserInfo, error) {
	key := fmt.Sprintf("%v%v", UserInfoCachePrefix, userId)
	result, err := a.rdb.Get(key).Result()
	if err != nil {
		return nil, err
	}
	user := model.UserInfo{}
	err = json.Unmarshal([]byte(result), &user)
	return &user, err
}

func (a *accountCache) DelUserInfoCache(userId int64) error {
	key := fmt.Sprintf("%v%v", UserInfoCachePrefix, userId)
	return a.rdb.Del(key).Err()
}
