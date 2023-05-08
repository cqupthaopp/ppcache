package lru

import (
	"errors"
	"fmt"
	"log"
)

func (c *Cache) Add(key string, value Value) {

	if v, exist := c.hMap[key]; exist {

		c.usedMemory += value.Len() + len(key) - v.Value.(*cacheElement).Len()
		c.cacheList.MoveToFront(v)
		v.Value.(*cacheElement).value = value

	} else {

		cacheIn := &cacheElement{
			key:   key,
			value: value,
		}
		c.usedMemory += cacheIn.Len()
		element := c.cacheList.PushFront(cacheIn)
		c.hMap[key] = element

	}

	c.len++

	for c.usedMemory > c.maxMemory {
		c.pop()
	}
}

func (c *Cache) pop() {

	if len(c.hMap) == 0 {
		return
	}

	c.len--

	element := c.cacheList.Back().Value.(*cacheElement)
	c.usedMemory -= element.Len()
	delete(c.hMap, element.key)
	c.cacheList.Remove(c.cacheList.Back())

}

func (c *Cache) Get(key string) (Value, error) {

	if v, exist := c.hMap[key]; exist {
		return v.Value.(*cacheElement).value, nil
	} else {
		return nil, errors.New("No cache")
	}

}

func (c *Cache) Del(key string) error {

	if v, exist := c.hMap[key]; exist {
		c.cacheList.Remove(v)
		delete(c.hMap, key)
		c.len--
		return nil
	} else {
		return errors.New("No cache")
	}

}

func (c *Cache) test() {

	log.Println(c)

	if c.usedMemory == 0 {
		return
	}
	e := c.cacheList.Front()

	for {

		fmt.Println(*(e.Value.(*cacheElement)))

		e = e.Next()

		if e == nil {
			break
		}

	}

}
