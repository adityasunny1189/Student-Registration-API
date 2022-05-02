package redisdb

import (
	"context"
	"encoding/json"
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

func SetData(key string, val []byte) {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func GetData(key string) (models.Student, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("error getting data from redis db", err)
		return models.Student{}, err
	}
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		return models.Student{}, err
	} else if err != nil {
		fmt.Println(err)
		return models.Student{}, err
	} else {
		var student models.Student
		json.Unmarshal([]byte(val), &student)
		return student, nil
	}
}
