package app

import (
	"sync"
	"time"
)

type ResponseCache struct {
	mu    sync.RWMutex
	cache map[string]CacheEntry
	ttl   time.Duration
}

type CacheEntry struct {
	response []byte
	expires  time.Time
}

func NewResponseCache(ttl time.Duration) *ResponseCache {
	return &ResponseCache{
		cache: make(map[string]CacheEntry),
		ttl:   ttl,
	}
}

func (rc *ResponseCache) Store(key string, data []byte) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	rc.cache[key] = CacheEntry{
		response: data,
		expires:  time.Now().Add(rc.ttl),
	}
}

func (rc *ResponseCache) Get(key string) ([]byte, bool) {

}

func (rc *ResponseCache) Cleanup() {

}
