package infras

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("cache.host") + ":" + viper.GetString("cache.port"),
		Password: viper.GetString("cache.password"),
		DB:       0,
	})
}
