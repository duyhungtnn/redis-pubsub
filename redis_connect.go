package redisWrapper

import (
	"encoding/json"
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
func ConnectToRedis(options ...int) error {
    indexDB := 0
    if len(options) >= 1 {
        indexDB = options[0]
    }
    
    redisAddress := os.Getenv("REDIS_ADDR")
    redisPW := os.Getenv("REDIS_PASSWORD")
    client = redis.NewClient(&redis.Options{
        Addr:     redisAddress,
        Password: redisPW,
        DB:       indexDB,
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
	fmt.Println("GetValuedForKey ", key, client, RedisClient())
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

//PublishMessage to pubsub channel
func PublishMessage(message, channel string) (int64, error) {
	return client.Publish(channel, message).Result()
}

//Publish interface{} to redis pubsub
func Publish(channel string, message interface{}) (int64, error) {
	// TODO reflect if interface{} type is string, Publish as-is
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	messageString := string(jsonBytes)
	return client.Publish(channel, messageString).Result()
}

//SubscriberToChannel atom
func SubscriberToChannel(channel string) (*redis.PubSub, error) {
	if len(channel) > 0 {
		return client.Subscribe(channel), nil
	}
	return client.Subscribe(), nil
}

//SubscriberToPChannel atom
func SubscriberToPChannel(pchannel string) (*redis.PubSub, error) {
	return client.PSubscribe(pchannel), nil
}
