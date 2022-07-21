package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, wasInCache := c.items[key]
	if wasInCache {
		c.queue.MoveToFront(item)
		item.Value.(*cacheItem).value = value
	} else {
		if len(c.items) == c.capacity {
			lastItem := c.queue.Back()
			c.queue.Remove(lastItem)
			delete(c.items, lastItem.Value.(*cacheItem).key)
		}
		newItem := c.queue.PushFront(&cacheItem{
			key:   key,
			value: value,
		})
		c.items[key] = newItem
	}
	return wasInCache
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, wasInCache := c.items[key]
	if !wasInCache {
		return nil, wasInCache
	}
	c.queue.MoveToFront(item)
	return item.Value.(*cacheItem).value, wasInCache
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = new(list)
}
