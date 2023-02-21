package redis

import (
	"bluebell/setting"
	"fmt"
	"github.com/go-redis/redis"
)

// Client 声明一个全局的 Client 变量
var Client *redis.Client

// Init 初始化连接
func Init(cfg *setting.RedisConfig) error {
	Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = Client.Close()
}
