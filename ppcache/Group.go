package ppcache

import (
	"errors"
	"ppCache/ppcache/lru"
	"sync"
)

type CacheGroup struct {
	main   *mainCache
	getter func(key string) (value []byte, err error)
}

const DefaultMaxMemory = 2 << 20

func NewGroup(cacheSize int, getter func(key string) (value []byte, err error)) *CacheGroup {

	//最大字节数 and getterFunc

	return &CacheGroup{
		main: &mainCache{
			lock:  sync.Mutex{},
			cache: lru.NewCache(cacheSize),
		},
		getter: getter,
	}

}

func (g *CacheGroup) SetGetter(getter func(key string) (value []byte, err error)) {
	g.getter = getter
}

func (g *CacheGroup) Get(key string) (ByteView, error) {

	if v, err := g.main.get(key); err == nil {
		return v, err
	}

	if g.getter == nil {
		return ByteView{}, errors.New("No cache Hit")
	}

	return g.load(key)

}

func (g *CacheGroup) load(key string) (ByteView, error) {
	v, err := g.getter(key)
	if err != nil {
		return ByteView{}, err
	}
	g.main.add(key, v)
	return v, err
}

func (g *CacheGroup) Del(key string) error {
	err := g.main.del(key)
	if err != nil {
		return err
	}

	return nil
}

func (g *CacheGroup) Add(key string, data []byte) {
	g.main.add(key, data)
}

func (g *CacheGroup) Len() int {
	return g.main.cache.Len()
}

func (g *CacheGroup) UsedMemory() int {
	return g.main.cache.UsedMemory()
}

func (g *CacheGroup) MaxMemory() int {
	return g.main.cache.MaxMemory()
}
