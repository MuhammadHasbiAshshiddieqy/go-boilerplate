package redis

import (
	"errors"
	"fmt"
	_config "microservice/shared/config"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var err error

// Init - redis init
func Init() error {
	conf := _config.GetConfig()
	dsn := fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return errors.New("Redis Connection Error")
	}

	return nil
}

// RedisManager - return db connection
func RedisManager() *redis.Client {
	return rdb
}
