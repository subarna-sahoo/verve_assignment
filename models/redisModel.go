package models

import (
	"context"
	"fmt"
	"time"
	"verve_assignment/utils"

	"github.com/redis/go-redis/v9"
)

// AddRequestIDToRedis stores the request ID in a Redis sorted set with the current timestamp
func StoreRequestID(reqID string) error {
	ctx := context.Background()
	timestamp := time.Now().Unix() // Current Unix timestamp

	// Use ZAdd with redis.Z from v9 package
	_, err := utils.RedisClient.ZAdd(ctx, "req:ids", redis.Z{
		Score:  float64(timestamp),
		Member: reqID,
	}).Result()

	if err != nil {
		return fmt.Errorf("failed to store request ID in Redis: %w", err)
	}

	return nil
}

// GetUniqueRequestIDsFromLastMinute fetches all unique request IDs received in the last minute
func GetUniqueRequestIDsFromLastMinute() ([]string, error) {
	ctx := context.Background()
	timestamp := time.Now().Unix() // Current Unix timestamp
	lastMinute := timestamp - 60   // One minute ago

	// Use ZRangeByScore to fetch all request IDs in the last minute
	result, err := utils.RedisClient.ZRangeByScore(ctx, "req:ids", &redis.ZRangeBy{
		Min: fmt.Sprintf("%d", lastMinute),
		Max: fmt.Sprintf("%d", timestamp),
	}).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch request IDs: %w", err)
	}

	return result, nil
}
