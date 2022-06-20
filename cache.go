package cache

import "time"

type Cache struct {
	store map[string]Entry
}

type Entry struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{store: make(map[string]Entry)}
}

func NewEntry(newValue string) Entry {
	return Entry{value: newValue, deadline: time.Now().AddDate(72, 0, 0)}
}

func (entry *Entry) isExpired() bool {
	return entry.deadline.Before(time.Now())
}

func (cache Cache) Get(key string) (string, bool) {
	entry := cache.store[key]
	if entry.value == "" {
		return "", false
	}

	if entry.isExpired() {
		delete(cache.store, key)
		return "", false
	}

	return entry.value, true
}

func (cache *Cache) Put(key, value string) {
	cache.store[key] = NewEntry(value)
}

func (cache Cache) Keys() []string {
	keys := make([]string, 0, len(cache.store))
	for k, v := range cache.store {
		if !v.isExpired() {
			keys = append(keys, k)
		} else {
			delete(cache.store, k)
		}

	}
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.store[key] = Entry{value: value, deadline: deadline}
}
