package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	fmt.Println("redis addr:", viper.GetString("redis.addr"))

	_, err := client.Ping().Result()
	if err != nil {
		log.Warnf("[redis] redis ping err: %+v", err)
		return nil, err
	}

	return client, err
}
