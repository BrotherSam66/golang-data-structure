package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisDb *redis.Client

// 初始化
func initRedisDb() error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := RedisDb.Ping().Result()
	return err
}

func main() {
	fmt.Println("")
	err := initRedisDb()
	if err != nil {
		fmt.Println("初始化redis出错：", err)
		return
	}
	fmt.Println("初始化redis成功")

}
