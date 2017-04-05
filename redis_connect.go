package redisWrapper

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

//RedisClient return global redis client
func RedisClient() *redis.Client {
	return client
}

//ConnectToRedis need call before use redis
func ConnectToRedis() error {
	redisAddress := os.Getenv("REDIS_ADDR")
	redisPW := os.Getenv("REDIS_PASSWORD")
	client = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPW,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return err
}

//DisconnectToRedis call when server close
func DisconnectToRedis() error {
	return client.Close()
}

//GetValuedForKey atom
func GetValuedForKey(key string) (interface{}, error) {
	fmt.Println("GetValuedForKey ", key)
	return client.Get(key).Result()
}

//SetValuedForKey atom
func SetValuedForKey(value, key string) error {
	fmt.Println("SetValuedForKey ", key)
	return client.Set(key, value, 0).Err()
}

//AddValueToList atom
func AddValueToList(value, key string) error {
	fmt.Println("AddValueToList ", key)
	return client.LPush(key, value).Err()
}

//LengthOfList automatically
func LengthOfList(key string) (int64, error) {
	return client.LLen(key).Result()
}
