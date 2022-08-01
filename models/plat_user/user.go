package plat_user

import (
	"context"
	"errors"
	"fmt"
	"rt-msg-carrier/global"
	"strconv"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type IdaasUser interface {
	PlatUser
	Get_Idaas_Token() (string, error)
}

type PlatUser interface {
	Get_Id() int
	Login() (bool, error)
	Is_Logged_In() (bool, error)
	Logout() (bool, error)
	IsAdmin() bool
}

type platuser struct {
	id           int
	is_active    bool
	is_anonymous bool
}

func (u *platuser) Get_Id() int {
	return 1
}

func (u *platuser) IsAdmin() bool {
	return false
}

func (u *platuser) redis_token() string {
	key := "go_redis_token:"
	key += strconv.Itoa(u.id)
	return key
}

func (u *platuser) Get_Idaas_Token() (string, error) {
	return "", nil
}

func (u *platuser) Login() (bool, error) {
	rdb := global.GetRedisClient()
	ctx := context.Background()
	err := rdb.Set(ctx, u.redis_token(), "user token in redis", time.Hour*2).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *platuser) Is_Logged_In() (bool, error) {
	rdb := global.GetRedisClient()
	ctx := context.Background()
	result, err := rdb.Get(ctx, u.redis_token()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		// 出其他错了
		return false, err
	}
	fmt.Println(result)
	return true, nil
}

func (u *platuser) Logout() (bool, error) {
	rdb := global.GetRedisClient()
	ctx := context.Background()
	result, err := rdb.Del(ctx, u.redis_token()).Result()
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return true, nil
}

func NewPlatUser(session_id int, args ...string) PlatUser {
	u := &platuser{
		id:           1,
		is_active:    true,
		is_anonymous: false,
	}
	u.id = session_id
	return u
}

type ThirdPartyUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

func (u *ThirdPartyUser) Get_Id() int {
	return 0
}

func (u *ThirdPartyUser) IsAdmin() bool {
	return true
}

func (u *ThirdPartyUser) Login() (bool, error) {
	return true, nil
}

func (u *ThirdPartyUser) Logout() (bool, error) {
	return true, nil
}

func (u *ThirdPartyUser) Is_Logged_In() (bool, error) {
	return true, nil
}
