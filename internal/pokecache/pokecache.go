package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
    cacheEntryMap map[string]cacheEntry
	mu *sync.Mutex
} 

type cacheEntry struct {
    createdAt time.Time
	val []byte
} 

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cacheEntryMap: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
} 

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.cacheEntryMap[key] = cacheEntry{createdAt: time.Now(), val: value}
	c.mu.Unlock()
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cE, exists := c.cacheEntryMap[key]
	c.mu.Unlock()
	if exists {
	    return cE.val, true
	} 
	return nil, false
} 

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for k, cE := range c.cacheEntryMap {
		    if cE.createdAt.Add(interval).Before(time.Now()) {
				delete(c.cacheEntryMap, k)
		    } 
		} 
		c.mu.Unlock()
	} 
} 
