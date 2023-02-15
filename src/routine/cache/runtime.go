package cache

import (
	"encoding/json"
	"errors"
	"time"

	"${package}/src/routine"
	"${package}/src/util/strings"
)

var globalCache *Cache

func init() {
	globalCache = NewCache(1000, 10*time.Minute)

	routine.Info("[cache] init cache")
	routine.Go("builtin-memcache", func() {
		for {
			time.Sleep(1 * time.Minute)
			globalCache.FlushExpired()
		}
	})
}

func Set(key string, value interface{}) {
	globalCache.Set(key, value)
}

func Get(key string) (interface{}, bool) {
	return globalCache.Get(key)
}

func Delete(key string) {
	globalCache.Del(key)
}

func SetCache[T any](key string, v T, features ...string) error {
	feature := strings.StringJoinWithDelim("_", features...)
	if feature != "" {
		key = feature + ":" + key
	}

	json_text, err := json.Marshal(v)
	if err != nil {
		return err
	}

	globalCache.Set(key, json_text)
	return nil
}

func GetCache[T any](key string, features ...string) (*T, error) {
	feature := strings.StringJoinWithDelim("_", features...)
	if feature != "" {
		key = feature + ":" + key
	}

	var v T
	json_text, ok := globalCache.Get(key)
	if !ok {
		return nil, errors.New("not found")
	}

	err := json.Unmarshal(json_text.([]byte), &v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func DeleteCache(key string) {
	globalCache.Del(key)
}
