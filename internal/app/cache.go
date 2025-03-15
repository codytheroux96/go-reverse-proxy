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
