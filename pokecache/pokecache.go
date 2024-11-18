package pokecache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	mu       *sync.Mutex
	entries  map[string]CacheEntry
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(s string) Cache {
	var newCache Cache
	newCache.mu = &sync.Mutex{}
	newCache.entries = map[string]CacheEntry{}
	var err error
	newCache.interval, err = time.ParseDuration(s)
	if err != nil {
		log.Fatal(err)
	}
	newCache.reapLoop()
	return newCache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, found := c.entries[key]
	return data.val, found
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for range ticker.C {
			for key, entry := range c.entries {
				t := time.Now()
				if t.Sub(entry.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
		}
	}()
}
