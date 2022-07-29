package redis

import (
	"bluebell/setting"
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的 client 变量
var client *redis.Client

// Init 初始化连接
func Init(cfg *setting.RedisConfig) error {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
