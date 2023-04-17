package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       viper.GetInt("redis.DB"),
	})
	client.Ping(context.Background())
	return nil, nil
}
