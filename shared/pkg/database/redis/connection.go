package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	rdb      = make(map[string]*redis.Client)
	connList []string
	err      error
)

type RedisOption struct {
	ConnName string
	Username string
	Password string
	Host     string
	Db       int
	Port     int
	Role     int
}

type redisOption interface {
	Init() error
}

// New - redis constructor
func New(role, port, db int, host, pwd, connname string) redisOption {
	return &RedisOption{
		ConnName: connname,
		Password: pwd,
		Host:     host,
		Db:       db,
		Port:     port,
		Role:     role,
	}
}

// Init - redis init
func (r *RedisOption) Init() error {
	dsn := fmt.Sprintf("%s:%d", r.Host, r.Port)
	rdb[r.ConnName] = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: r.Password, // no password set
		DB:       r.Db,       // use default DB
	})

	_, err = rdb[r.ConnName].Ping(rdb[r.ConnName].Context()).Result()
	if err != nil {
		errMsg := fmt.Sprintf("Redis ERROR : error to create %s connection", r.ConnName)
		return errors.New(errMsg)
	}

	connList = append(connList, r.ConnName)

	return nil
}

// GetConnection - return db connection
func GetConnection(connname string) *redis.Client {
	return rdb[connname]
}

// GetConnectionList - return mysql connection names
func GetConnectionList() []string {
	return connList
}
