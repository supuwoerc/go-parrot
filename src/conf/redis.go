package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var client *redis.Client

const duration = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	client = rdb
	return &RedisClient{}, nil
}

func (c *RedisClient) Set(key string, value any) error {
	return client.Set(context.Background(), key, value, duration).Err()
}

func (c *RedisClient) Get(key string) (any, error) {
	return client.Get(context.Background(), key).Result()
}

func (c *RedisClient) Delete(keys ...string) error {
	return client.Del(context.Background(), keys...).Err()
}
