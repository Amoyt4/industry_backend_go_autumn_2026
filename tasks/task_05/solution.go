package main

type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	link := &Cache[K, V]{
		capacity: capacity,
		items:    make(map[K]V),
	}
	return link
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	if c.capacity <= 0 {
		var zero V
		return zero, false
	}
	val, ok := c.items[key]
	return val, ok
}

func (c *Cache[K, V]) Set(key K, value V) bool {
	if c.capacity <= 0 {
		return false
	}
	_, exists := c.items[key]
	if exists {
		c.items[key] = value
		return true
	}
	if len(c.items) >= c.capacity {
		return false
	}
	c.items[key] = value
	return true
}
