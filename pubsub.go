package pubsub

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

func init() {
	var client *redis.Client
	client = redis.NewTCPClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
	Service = &PubSub{client}
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
