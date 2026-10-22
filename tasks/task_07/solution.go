package main

import (
	"container/list"
	"sync"
)

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity <= 0 {
		return value, false
	}

	if elem, exists := c.items[key]; exists {
		c.ll.MoveToFront(elem)
		return elem.Value.(*entry[K, V]).value, true
	}

	return value, false
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity <= 0 {
		return
	}

	if elem, exists := c.items[key]; exists {
		elem.Value.(*entry[K, V]).value = value
		return
	}

	if c.ll.Len() >= c.capacity {
		victim := c.ll.Back()
		if victim != nil {
			c.ll.Remove(victim)
			delete(c.items, victim.Value.(*entry[K, V]).key)
		}
	}

	newEntry := &entry[K, V]{key: key, value: value}
	elem := c.ll.PushFront(newEntry)
	c.items[key] = elem
}
