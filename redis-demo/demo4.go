package main

import (
	"code.byted.org/kv/redis-v6"
	"fmt"
	"time"
)


func main() {
	var redisClient1 = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	var redisClient2 = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

}

func writeToken(redisClient *redis.Client, token string, value string) string {
	// 如果存在直接返回
	if redisClient.Exists(token).Val() == 1 {
		return redisClient.Get(token).Val()
	}

	// 如果不存在有并发问题
	// 1.获取锁
	if redisClient.SetNX(token, value, 10*time.Millisecond).Val() {

	} else {
		// 2.未获取成功, 表示已经存在token

		for times := 0; times < 3; times++{
			if redisClient.Exists(token).Val() == 1 {
				// 判断过期时间
				newValue := redisClient.Get(token).Val()
				return newValue
			}

			time.Sleep(5 * time.Millisecond)
		}
	}
}
