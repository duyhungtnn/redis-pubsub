package redisWrapper

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

//PubSub atom
type PubSub struct {
	client *redis.Client
}

//Service PubSub globall ival
var Service *PubSub

//InitPubSub Service before use , call after connect to redis successful
func InitPubSub() {
	Service = &PubSub{RedisClient()}
}

//PublishString to redis pubsub
func (ps *PubSub) PublishString(channel, message string) *redis.IntCmd {
	return ps.client.Publish(channel, message)
}

//Publish interface{} to redis pubsub
func (ps *PubSub) Publish(channel string, message interface{}) *redis.IntCmd {
	// TODO reflect if interface{} type is string, Publish as-is
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	messageString := string(jsonBytes)
	return ps.client.Publish(channel, messageString)
}
