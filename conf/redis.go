package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

const defaultDuration = 30 * 24 * time.Hour

type RedisClient struct {
}

var client *redis.Client

func InitRedis() (*RedisClient, error) {
	client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       viper.GetInt("redis.DB"),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any) error {
	return client.Set(context.Background(), key, value, defaultDuration).Err()
}

func (rc *RedisClient) Get(key string) error {
	return client.Get(context.Background(), key).Err()
}

func (rc *RedisClient) Delete(key ...string) error {
	return client.Del(context.Background(), key...).Err()
}
func (rc *RedisClient) GetExpireDuration(key string) (time.Duration, error) {
	return client.TTL(context.Background(), key).Result()
}
