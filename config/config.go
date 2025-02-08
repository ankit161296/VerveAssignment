package config

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	"os"
	"sync"
)

// Global variables
var (
	Ctx           = context.Background()
	Deduplication *redis.Client
	RequestCount  sync.Map
	LogFile       *os.File
)

// Initialize dependencies
func Init() {
	// Use environment variables for Redis and Kafka (Dockerized)
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	// Initialize Redis
	Deduplication = redis.NewClient(&redis.Options{
		Addr: redisHost,
	})
}

// Cleanup resources
func Cleanup() {
	Deduplication.Close()
}
