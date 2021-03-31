package cache

import "log"

func Client() *RedisStore {
	store, err := NewRedisStore(32, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatal("unable to connect to redis")
	}

	return &store
}
