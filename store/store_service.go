package store

// Store Service Setup

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define the struct wrapper around the Redis client
type storeService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeSvc = &storeService{}
	ctx      = context.Background()
)

const CacheDuration = 6 * time.Hour

// Initialize the store service and return a store pointer
func InitializeStore() *storeService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeSvc.redisClient = redisClient
	return storeSvc
}

// Storage API design and Implementation
func SaveUrlMapping(shortUrl, originalUrl string, userId string) {
	err := storeSvc.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Error saving URL mapping: %v", err))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeSvc.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Error retrieving URL mapping: %v", err))
	}
	return result
}
