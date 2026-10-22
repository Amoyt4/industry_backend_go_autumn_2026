package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
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
	if c.capacity <= 0 {
		return
	}

	if elem, exists := c.items[key]; exists {
		elem.Value.(*entry[K, V]).value = value
		return
	}

	if c.ll.Len() >= c.capacity {
		backElem := c.ll.Back()
		if backElem != nil {
			c.ll.Remove(backElem)
			oldKey := backElem.Value.(*entry[K, V]).key
			delete(c.items, oldKey)
		}
	}
	newEntry := &entry[K, V]{key: key, value: value}
	newElem := c.ll.PushFront(newEntry)
	c.items[key] = newElem
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
