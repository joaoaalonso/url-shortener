package configs

import "github.com/go-redis/redis/v8"

var rdb *redis.Client

// GetRedisConnection return a redis connection
func GetRedisConnection() *redis.Client {
	if rdb != nil {
		return rdb
	}

	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opt)
	return rdb
}
