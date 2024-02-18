package rwredis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
	PoolSize     int
}

func NewRedis(conf *RedisConfig) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		DialTimeout:  conf.DialTimeout,
		PoolSize:     conf.PoolSize,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		messqge := fmt.Errorf("credis connect error: %v", err)
		panic(messqge)
	}
	return
}
