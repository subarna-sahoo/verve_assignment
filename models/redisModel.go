package models

import (
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	// Redis client initialization
	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Adjust if necessary
	})
}

func CheckUniqueID(id string) (bool, error) {
	_, err := redisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		// ID is unique
		redisClient.Set(ctx, id, "1", time.Minute)
		return true, nil
	}
	return false, err
}
