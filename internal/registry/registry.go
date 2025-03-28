package registry

import (
	"sync"
	"time"
)

type Registry struct {
	servers map[string]Server
	mu      sync.RWMutex
}

type Server struct {
	Name         string
	BaseURL      string
	Prefixes     []string
	RegisteredAt time.Time
}

func (r *Registry) NewRegistry() *Registry {
	return &Registry{
		servers: make(map[string]Server),
	}
}
