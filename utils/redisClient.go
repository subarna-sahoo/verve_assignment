package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Global Redis client
var RedisClient *redis.Client

// InitRedis initializes the Redis client once
func InitRedis() {
	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",           // Redis container address (from Docker Compose)
		Password: "your_secure_password", // Password set in docker-compose
		DB:       0,                      // Default DB
	})

	// Ping to test connection with a context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5-second timeout
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		// Log a detailed error message
		fmt.Println("❌ Failed to connect to Redis:", err)
	} else {
		// Successful connection
		fmt.Println("✅ Redis Connected!")
	}
}
