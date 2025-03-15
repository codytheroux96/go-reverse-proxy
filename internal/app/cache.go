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

func (rc *ResponseCache) Store() {

}

func (rc *ResponseCache) Get() ([]byte, bool) {
	
}

func (rc *ResponseCache) Cleanup() {
	
}