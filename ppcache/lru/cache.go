package lru

import (
	"container/list"
)

type Cache struct {
	maxMemory  int //最大内存
	usedMemory int //已经使用的内存

	hMap      map[string]*list.Element //哈希表
	cacheList *list.List               //内存链表

	len int //已存缓存个数

}

type cacheElement struct {
	key   string
	value Value
}

func (c *Cache) UsedMemory() int {
	return c.usedMemory
}

func (c *Cache) MaxMemory() int {
	return c.maxMemory
}

func (c *Cache) Len() int {
	return c.len
}

func (c *cacheElement) Len() int {
	return len(c.key) + c.value.Len()
}

const DefaultMaxMemory = 2 << 20 // 2MB

func NewCache(maxMemory int) *Cache {
	return &Cache{
		maxMemory:  maxMemory,
		usedMemory: 0,
		hMap:       make(map[string]*list.Element),
		cacheList:  list.New(),
	}
}
