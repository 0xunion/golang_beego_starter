package decorator

/*
	this package is used to decorate the curd functions of the model
	so that we can cache the data in the redis or memcache or other cache
	what we need to do is just specify which cache we want to use
	and implement the interface of the cache
*/

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"sync"
	"time"

	config "github.com/0xunion/exercise_back/src/const/conf"
	redis_key "github.com/0xunion/exercise_back/src/const/key"
	routine "github.com/0xunion/exercise_back/src/routine"
	memcache "github.com/0xunion/exercise_back/src/routine/cache"
	strings "github.com/0xunion/exercise_back/src/util/strings"
	redis "github.com/go-redis/redis/v8"
)

type Cache interface {
	// get the data from the cache
	Get(key string) (string, error)
	// set the data to the cache
	Set(key string, value string) error
	// delete the data from the cache
	Del(key string) error
}

type RedisCache struct {
	Cache
	Handle *redis.Client
	Alive  bool
}

func (r *RedisCache) Get(key string) (string, error) {
	if r.Alive {
		return r.Handle.Get(r.Handle.Context(), key).Result()
	}
	return "", errors.New("redis is not alive")
}

func (r *RedisCache) Set(key string, value string) error {
	if r.Alive {
		return r.Handle.Set(r.Handle.Context(), key, value, redisExpire).Err()
	}
	return errors.New("redis is not alive")
}

func (r *RedisCache) Del(key string) error {
	if r.Alive {
		return r.Handle.Del(r.Handle.Context(), key).Err()
	}
	return errors.New("redis is not alive")
}

var defaultCache Cache
var cacheTypeNames sync.Map
var redisExpire time.Duration

func init() {
	expire, err := strconv.Atoi(config.RedisExpire())
	if err != nil || expire <= 0 {
		redisExpire = time.Hour
	} else {
		redisExpire = time.Duration(expire) * time.Second
	}

	reconnectRedis := func() {
		routine.Info("[Cache] Start connect redis")
		defaultCache = &RedisCache{
			Handle: redis.NewClient(&redis.Options{
				Addr:     config.RedisHost() + ":" + config.RedisPort(),
				Password: config.RedisPass(),
				DB:       0,
			}),
			Alive: true,
		}
	}

	routine.Info("[Cache] Start init redis")
	reconnectRedis()
	// launch a goroutine to monitor redis connection
	routine.Go("redis_monitor", func() {
		routine.Info("[Cache] Start monitor redis connection")
		for {
			_, err := defaultCache.(*RedisCache).Handle.Ping(defaultCache.(*RedisCache).Handle.Context()).Result()
			if err != nil {
				routine.Error("[Cachce] redis connection failed, err: %v", err)
				defaultCache.(*RedisCache).Alive = false
				reconnectRedis()
			}
			time.Sleep(time.Second * 5)
		}
	})
}

func WithCacheGet[T any](fn func(key string) (T, error), key string) (T, error) {
	var data T
	// get type name of T
	typeName := reflect.TypeOf(data).Name()

	// fetch the data from the cache
	item, err := defaultCache.Get(strings.StringJoin(
		redis_key.REDIS_KEY_PREFIX, ":", typeName, ":", key,
	))
	if err == redis.Nil {
		return fn(key)
	}

	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(item), &data)
	if err == nil {
		// set the data to the cache
		defaultCache.Set(redis_key.REDIS_KEY_PREFIX+key, item)
	}

	return data, err
}

func WithMemCacheGet[T any](fn func(key string) (T, error), key string) (T, error) {
	var data T

	// get type name of T
	typeName := reflect.TypeOf(data).Name()

	key = strings.StringJoin(
		redis_key.MEMCACHE_KEY_PREFIX, ":", typeName, ":", key,
	)
	// fetch the data from the cache
	item, err := memcache.GetCache[T](key)
	if err == nil {
		return *item, nil
	}

	data, err = fn(key)
	if err == nil {
		// set the data to the cache
		memcache.SetCache(key, data)
	}

	return data, err
}
