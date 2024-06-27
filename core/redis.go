package core

import (
	"context"
	"time"

	"github.com/backunderstar/zew/global"
	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr(),
		Password: redisConfig.Password,
		DB:       0,
		PoolSize: redisConfig.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Log.Errorf("redis连接失败%s", redisConfig.Addr())
		return nil
	}
	global.Log.Info("Redis初始化连接成功")
	return rdb
}
