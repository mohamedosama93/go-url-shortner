package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	client *redis.Client
}

var storeService = &StorageService{}
var ctx = context.Background()

const duration = 6 * time.Hour

func InitializeStore() *StorageService {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	pong,err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to redis %v" , err))
	}

	fmt.Printf("Connected successfully to redis message = %s", pong);
	storeService.client = client

	return storeService;
}

func SaveUrl(key string, url string, userId string) {
	err := storeService.client.Set(ctx, key, url, duration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to store {%s} with error %v", key, err))
	}
}

func getUrl(key string) string {
	res, err := storeService.client.Get(ctx, key).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve {%s} with error %v", key, err))
	}

	return res
}