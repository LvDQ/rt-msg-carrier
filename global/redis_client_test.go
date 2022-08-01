package global

import (
	"context"
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	rdb := GetRedisClient()
	ctx := context.Background()
	result, _ := rdb.Ping(ctx).Result()
	fmt.Println("test redis connection: " + result)
	t.Log("complete")
}
