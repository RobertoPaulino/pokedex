package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
  cache map[string]cacheEntry
  mu sync.Mutex
}

type cacheEntry struct { 
  createdAt time.Time
  val []byte
}

func NewCache(interval time.Duration) Cache {
  cache := Cache{}
  cache.reapLoop(interval)

  return cache
}

func (c *Cache) add(key string, val []byte) {
  c.mu.Lock()
  defer c.mu.Unlock()
  c.cache[key] = cacheEntry{
    createdAt: time.Now(),
    val: val,
  }

}

func (c *Cache) Get(key string) ([]byte, bool) {
  c.mu.Lock()
  defer c.mu.Unlock()

  entry, ok := c.cache[key]
  if !ok {
    return []byte{}, false
  }
  
  return entry.val, true
  
}

func (c *Cache) reapLoop(interval time.Duration) {
  tick := time.NewTicker(interval)

}
