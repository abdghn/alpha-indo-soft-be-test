package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func ConnectRedis() *redis.Client {
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	host := viper.Get("REDIS_HOST").(string)
	password := viper.Get("REDIS_PASSWORD").(string)

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}
