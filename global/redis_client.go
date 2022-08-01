package global

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

var _rdb *redis.Client

var ctx = context.Background()

func init() {
	_rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10,
	})
	pong, _ := _rdb.Ping(ctx).Result()
	fmt.Println("Redis Client: " + pong)
}

func GetRedisClient() *redis.Client {
	return _rdb
}
