package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RedisPingCheck(addr string) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
