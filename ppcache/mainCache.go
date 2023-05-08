package ppcache

import (
	"errors"
	"ppCache/ppcache/lru"
	"sync"
)

type mainCache struct {
	lock  sync.Mutex
	cache *lru.Cache
}

func (mc *mainCache) add(key string, value ByteView) {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.cache.Add(key, value)
}

func (mc *mainCache) del(key string) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	return mc.cache.Del(key)
}

func (mc *mainCache) get(key string) (ByteView, error) {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	if v, err := mc.cache.Get(key); err == nil {
		return v.(ByteView), nil
	} else {
		return ByteView{}, errors.New("No cache hit")
	}

}

func (mc *mainCache) len() int {
	return mc.cache.Len()
}
