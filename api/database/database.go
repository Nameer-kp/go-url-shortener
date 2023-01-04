package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()
var rdb1, rdb0 *redis.Client = CreateClient(1), CreateClient(0)

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}

func GetRDBClient0() *redis.Client {
	return rdb0
}
func GetRDBClient1() *redis.Client {
	return rdb1
}
