package redis

import "github.com/redis/go-redis/v9"

var Client *redis.Client
var Nil = redis.Nil

func New() {
	Client = redis.NewClient(&redis.Options{})
}
