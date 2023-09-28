package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "https://example.com"
	val := []byte("somedata")

	cache.Add(key, val)

	expected := cache.cache[key].val

	if string(cache.cache[key].val) != string(val) {
		t.Errorf("Found %s, expected %s", string(expected), string(val))
	}
}

func TestGetFound(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "https://example.com"
	val := []byte("somedata")

	cache.Add(key, val)

	result, ok := cache.Get(key)

	if !ok {
		t.Errorf("Expected to find key %s", key)
	}

	if string(result) != string(val) {
		t.Errorf("Found %s, expected %s", string(result), string(val))
	}
}

func TestGetMissing(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "https://example.com"

	result, ok := cache.Get(key)

	if ok {
		t.Errorf("Expected to not find key %s", key)
	}

	if result != nil {
		t.Errorf("Found %s, expected %v", string(result), nil)
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
	}
}
