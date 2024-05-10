package redis

import (
	"time"

	"github.com/go-redis/redis"
)

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	client.Set("MemberID-1-1234", "1234", 0)
	client.Set("MemberID-1-1235", "1235", 0)
	client.Set("TTL-30", "30sec", 30*time.Second)
}
