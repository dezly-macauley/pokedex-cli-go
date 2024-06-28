package pokecache

import (
	"sync"
	"time"
)

// NOTE: The purpose of this file is to cache the locations of the map call
// That way a network call does not need to be made to the PokeAPI each time
// the user needs to navigate to a page

type Cache struct {
    cache map[string] cacheEntry
    mux *sync.Mutex
}

type cacheEntry struct {
    // caching HTTP response bodies
    val []byte
    createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
    c := Cache {
        cache: make(map[string]cacheEntry),
        mux: &sync.Mutex{},
    }
    
    // NOTE: The caching needs to be done in a new goroutine or the program 
    // will stall forever

    go c.reapLoop(interval)
    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.mux.Lock()
    defer c.mux.Unlock()
    c.cache[key] = cacheEntry {
        val: val,
        createdAt: time.Now().UTC(),
    }

}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mux.Lock()
    defer c.mux.Unlock()
    cacheE, ok := c.cache[key]
    return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {

    ticker := time.NewTicker(interval)

    for range ticker.C {
        c.reap(interval)
    }

}

func (c *Cache) reap(interval time.Duration) {
    c.mux.Lock()
    defer c.mux.Unlock()

    timeAgo := time.Now().UTC().Add(-interval)
    for k, v := range c.cache {
        
        if v.createdAt.Before(timeAgo) {
            delete(c.cache, k)
        }

    }

}
