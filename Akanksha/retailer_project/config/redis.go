package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
		DB:   0,                
	})

	// Test connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Println("Warning: Redis connection failed. Running without caching.")
		RedisClient = nil  
	} else {
		fmt.Println("Connected to Redis successfully!")
	}
}
