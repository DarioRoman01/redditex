package cache

import (
	"log"

	"github.com/go-redis/redis/v8"
)

// this client only work with sessions
func Client() *RedisStore {
	store, err := NewRedisStore(32, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatal("unable to connect to redis")
	}

	return &store
}

// this client is to do more complex interactions with redis
func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
