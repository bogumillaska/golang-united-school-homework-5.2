package cache

import (
	"testing"
	"time"
)

func TestExpiredValues(t *testing.T) {
	cache := NewCache()
	cache.PutTill("key1", "value1", time.Now().AddDate(0, -1, 0))
	value1, ok := cache.Get("key1")
	if ok {
		t.Error("OK should be false")
	}
	if value1 != "" {
		t.Errorf("Value should be empty, got : %v", value1)
	}
}
