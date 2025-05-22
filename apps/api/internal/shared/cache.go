package shared

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	if exists {
		log.Info().Str("key", key).Msg("🟢 Cache hit")
	} else {
		log.Info().Str("key", key).Msg("🔴 Cache miss")
	}
	return value, exists
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	log.Info().Str("key", key).Msg("💾 Cache set")
	c.data[key] = value
}

func (c *Cache) Clear(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
