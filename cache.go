package redisWrapper

import (
	"time"

	"github.com/go-redis/cache"
	"github.com/go-redis/cache/lrucache"
	"github.com/go-redis/redis"
	"gopkg.in/vmihailenco/msgpack.v2"
)

var ring *redis.Ring
var codec *cache.Codec

//InitCacheWithOptions atom
func InitCacheWithOptions(redisRingMap map[string]string, password string, db int, uselocalCache bool, cacheExpiration time.Duration, cacheMaxLen int) {
	/*
		 map[string]string{
				"server1": ":6379",
				"server2": ":6380",
			}
	*/

	ring = redis.NewRing(&redis.RingOptions{
		Addrs:    redisRingMap,
		Password: password,
		DB:       db,
	})

	codec = &cache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	if uselocalCache {
		codec.LocalCache = lrucache.New(cacheExpiration, cacheMaxLen)
	}
}

//InitCacheWithMapAddrs with ring of redis server
func InitCacheWithMapAddrs(redisRingMap map[string]string, password string, db int) {
	InitCacheWithOptions(redisRingMap, password, db, false, 0, 0)
}

//DestroyCache before close server
func DestroyCache() {
	codec = nil
	ring.Close()
	ring = nil
}

//SaveObjectToCache method to save object to cache
func SaveObjectToCache(obj interface{}, key string, expiration time.Duration) error {
	return codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: expiration,
	})
}

//RemoveObjectFromCache method to remove
func RemoveObjectFromCache(key string) error {
	return codec.Delete(key)
}

//GetObjectFromCache method to get object from cache
func GetObjectFromCache(key string, obj interface{}) error {
	return codec.Get(key, &obj)
}
