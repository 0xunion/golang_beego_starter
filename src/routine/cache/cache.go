package cache

/*
	cache module likes memcache, but it's not a real memcache.
*/

import (
	"container/list"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	// cache data
	data map[string]*list.Element
	// cache list
	list *list.List
	// cache size
	size int
	// cache expire time, when item is expired, it will be removed
	expire time.Duration
}

type Item struct {
	Key        string
	Value      interface{}
	ExpireTime time.Time // expire time, when item is expired, it will be removed
}

func NewCache(size int, expire time.Duration) *Cache {
	return &Cache{
		data:   make(map[string]*list.Element, size),
		list:   list.New(),
		size:   size,
		expire: expire,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		c.list.MoveToFront(e)
		e.Value.(*Item).Value = value
		return
	}

	e := c.list.PushFront(&Item{
		Key:        key,
		Value:      value,
		ExpireTime: time.Now().Add(c.expire),
	})
	c.data[key] = e

	if c.list.Len() > c.size {
		c.removeOldest()
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	if e, ok := c.data[key]; ok {
		item := e.Value.(*Item)
		if item.ExpireTime.After(time.Now()) {
			// refresh expire time
			item.ExpireTime = time.Now().Add(c.expire)
			c.list.MoveToFront(e)
			return item.Value, true
		}
	}
	return nil, false
}

func (c *Cache) Del(key string) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		c.list.Remove(e)
		delete(c.data, key)
	}
}

func (c *Cache) removeOldest() {
	e := c.list.Back()
	if e != nil {
		c.list.Remove(e)
		kv := e.Value.(*Item)
		delete(c.data, kv.Key)
	}
}

func (c *Cache) Len() int {
	c.RLock()
	defer c.RUnlock()

	return c.list.Len()
}

func (c *Cache) Clear() {
	c.Lock()
	defer c.Unlock()

	c.list.Init()
	c.data = make(map[string]*list.Element, c.size)
}

func (c *Cache) Keys() []string {
	c.RLock()
	defer c.RUnlock()

	keys := make([]string, 0, c.list.Len())
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) Items() map[string]interface{} {
	c.RLock()
	defer c.RUnlock()

	items := make(map[string]interface{}, c.list.Len())
	for k, v := range c.data {
		items[k] = v.Value.(*Item).Value
	}
	return items
}

func (c *Cache) Expire(key string, expire time.Duration) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		e.Value.(*Item).ExpireTime = time.Now().Add(expire)
	}
}

func (c *Cache) ExpireAt(key string, t time.Time) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		e.Value.(*Item).ExpireTime = t
	}
}

func (c *Cache) ExpireDuration(key string, d time.Duration) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		e.Value.(*Item).ExpireTime = e.Value.(*Item).ExpireTime.Add(d)
	}
}

func (c *Cache) TTL(key string) time.Duration {
	c.RLock()
	defer c.RUnlock()

	if e, ok := c.data[key]; ok {
		return e.Value.(*Item).ExpireTime.Sub(time.Now())
	}
	return -1
}

func (c *Cache) Has(key string) bool {
	c.RLock()
	defer c.RUnlock()

	_, ok := c.data[key]
	return ok
}

func (c *Cache) Incr(key string) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		e.Value.(*Item).Value = e.Value.(*Item).Value.(int) + 1
	}
}

func (c *Cache) Decr(key string) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		e.Value.(*Item).Value = e.Value.(*Item).Value.(int) - 1
	}
}

func (c *Cache) IncrBy(key string, delta int) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		e.Value.(*Item).Value = e.Value.(*Item).Value.(int) + delta
	}
}

func (c *Cache) DecrBy(key string, delta int) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.data[key]; ok {
		e.Value.(*Item).Value = e.Value.(*Item).Value.(int) - delta
	}
}

func (c *Cache) FlushExpired() {
	c.Lock()
	defer c.Unlock()

	for k, v := range c.data {
		if v.Value.(*Item).ExpireTime.Before(time.Now()) {
			c.list.Remove(v)
			delete(c.data, k)
		}
	}
}

func (c *Cache) Flush() {
	c.Lock()
	defer c.Unlock()

	c.list.Init()
	c.data = make(map[string]*list.Element, c.size)
}

func (c *Cache) Range(f func(key string, value interface{})) {
	c.RLock()
	defer c.RUnlock()

	for k, v := range c.data {
		f(k, v.Value.(*Item).Value)
	}
}
