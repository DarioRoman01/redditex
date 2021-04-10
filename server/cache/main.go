package cache

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to read env")
	}
}

// client for sessions
func Client() *RedisStore {
	store, err := NewRedisStore(32, "tcp", os.Getenv("REDIS_ADDRS"), os.Getenv("REDIS_PWD"), []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Fatal("unable to connect to redis: ", err)
	}

	return &store
}

// Client to interact with redis
func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRS"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	return rdb
}
