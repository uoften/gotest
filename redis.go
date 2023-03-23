package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	options := &redis.Options{
		Addr:        "127.0.0.1:6379",
		Dialer:      nil,
		OnConnect:   nil,
		Password:    "",
		DB:          0,
		DialTimeout: time.Duration(60) * time.Millisecond,
		PoolSize:    1024,
	}
	connect := redis.NewClient(options)
	pong := connect.Ping()
	if pong.String() != "ping: PONG" {
		fmt.Println("连接Redis ping失败")
	}
	client := connect

	//hash-------------------------------------------
	hashKey := "userkey_1"
	//set hash 适合存储结构
	client.HSet(hashKey, "name", "叶子")
	client.HSet(hashKey, "age", 18)

	//get hash
	hashGet, _ := client.HGet(hashKey, "name").Result()
	fmt.Println("HGet name", hashGet)


	//获取所有hash 返回map
	hashGetAll, _ := client.HGetAll(hashKey).Result()
	fmt.Println("HGetAll", hashGetAll)

	client.HDel(hashKey,"name")

	//获取所有hash 返回map
	hashGetAll, _ = client.HGetAll(hashKey).Result()
	fmt.Println("HGetAll", hashGetAll)

}
