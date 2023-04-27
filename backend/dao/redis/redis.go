package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/webforum/setting"
)

var (
	client *redis.Client
)
var ctx = context.Background()

func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println("redis init success!")
	return err

}

func Close() {
	_ = client.Close()
}
