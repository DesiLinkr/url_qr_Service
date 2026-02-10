package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

var RedisClientCounter *redis.Client

func init() {

	RedisClient = redis.NewClient(&redis.Options{

		Addr:     getEnv("REDIS_URL", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnvAsInt("REDIS_DB", 0),
	})

	RedisClientCounter = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_COUNTER_URL", "localhost:6379"),
		Password: getEnv("REDIS_COUNTER_PASSWORD", ""),
		DB:       getEnvAsInt("REDIS_DB", 0),
	})
	_, Countererr := RedisClientCounter.Ping(context.Background()).Result()
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("❌ Redis connection failed:", err)
	}

	if Countererr != nil {
		log.Fatal("❌ Redis Counter connection failed:", err)
	}
	log.Println("✅ Redis connected")
}
