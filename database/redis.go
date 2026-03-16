package database

import (
	"context"
	"sofia-backend/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis(s *config.Config) *redis.Client {

	ctx := context.TODO()

	// Handle Redis configuration for both username/password and no authentication
	var redisOptions *redis.Options
	if s.Redis.Username == "" && s.Redis.Password == "" {
		redisOptions = &redis.Options{
			Addr: s.Redis.Host,
		}
	} else if s.Redis.Username != "" && s.Redis.Password != "" {
		redisOptions = &redis.Options{
			Addr:     s.Redis.Host,
			Username: s.Redis.Username,
			Password: s.Redis.Password,
		}
	} else {
		panic("Redis configuration must include both username and password or neither.")

	}

	RedisClient := redis.NewClient(redisOptions)
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return RedisClient
}
