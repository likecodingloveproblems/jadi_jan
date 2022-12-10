package models

import (
	"github.com/go-redis/redis/v8"
)

func NewRedis() store {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return store{db: rdb}
}

type store struct {
	db *redis.Client
}

var redisStore = NewRedis()
