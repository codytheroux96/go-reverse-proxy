package registry

import (
	"fmt"
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

func NewRegistry() *Registry {
	return &Registry{
		servers: make(map[string]Server),
	}
}

func (r *Registry) Register(s Server) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.servers[s.Name]; exists {
		return fmt.Errorf("server '%s' already registered", s.Name)
	}

	r.servers[s.Name] = s
	return nil
}

func (r *Registry) Deregister(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.servers[name]; !exists {
		return fmt.Errorf("server '%s' does not exist... cannot deregister", name)
	}

	delete(r.servers, name)
	return nil
}

func (r *Registry) ListRegistered() []Server {
	r.mu.RLock()
	defer r.mu.RUnlock()

	servers := make([]Server, 0, len(r.servers))
	for _, server := range r.servers {
		servers = append(servers, server)
	}

	return servers
}
