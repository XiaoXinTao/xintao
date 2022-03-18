package dal

import (
	"time"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic("failed to initial redis " + err.Error())
	}
}

const (
	ExpirationUser  = time.Minute * 5
	ExpirationToken = time.Hour * 24 * 7
)

const (
	// 登录
	keyUserToken     = "account_rpc2::user:%v:token"      // web_server:app:{app_id}:user:{sdk_open_id}:token
)
