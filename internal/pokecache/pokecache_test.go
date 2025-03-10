package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cache := NewCache(time.Minute)
	cache.Add("test", []byte("test"))
	if len(cache.cache) != 1 {
		t.Errorf("Expected 1 item in cache, got %d", len(cache.cache))
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(time.Minute)
	cache.Add("test", []byte("test"))
	val, ok := cache.Get("test")
	if !ok {
		t.Errorf("Expected item in cache, got none")
	}
	if string(val) != "test" {
		t.Errorf("Expected test, got %s", string(val))
	}
}
