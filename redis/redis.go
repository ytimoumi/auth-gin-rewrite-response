package redis

import (
	"auth/env"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// Interface for DB API
type RedisServicer interface {
	RedisSetAccessToken(accessToken string, value interface{}) error
	RedisGetAccessToken(accessToken string) (string, error)
	RedisDeleteAccessToken(accessToken string) error
}

type RedisService struct {
	r *redis.Client
}

func NewDefaultRedisServicer() RedisServicer {
	client := redis.NewClient(&redis.Options{
		Addr: env.GetEnv("REDIS_URL", "localhost:6379"),
	})
	return &RedisService{
		r: client,
	}
}

func NewRedisServicer(r *redis.Client) RedisServicer {
	return &RedisService{
		r: r,
	}
}

// RedisSetAccessToken function that set new uuid to redis
func (r *RedisService) RedisSetAccessToken(accessToken string, value interface{}) error {
	bytesSession, err := json.Marshal(value)
	if err != nil {
		return err
	}
	status := r.r.Set("access_token:"+accessToken, string(bytesSession), 4*time.Hour) // expiration time is : 4 hour
	err = status.Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisService) RedisGetAccessToken(accessToken string) (string, error) {
	return r.r.Get("access_token:" + accessToken).Result()
}

func (r *RedisService) RedisDeleteAccessToken(accessToken string) error {
	return r.r.Del("access_token:" + accessToken).Err()
}
