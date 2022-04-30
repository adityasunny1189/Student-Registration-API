package redisdb

import (
	"context"
	"fmt"

	"github.com/adityasunny1189/Student-Registration-API/models"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func InitializeRedisDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetData(key string, val models.Student) {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func GetData(key string) (models.Student, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("error connecting to redis db", err)
		return models.Student{}, err
	}
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		return models.Student{}, err
	} else if err != nil {
		fmt.Println(err)
		return models.Student{}, err
	} else {
		fmt.Println("key", val)
		return models.Student{}, nil
	}
}
