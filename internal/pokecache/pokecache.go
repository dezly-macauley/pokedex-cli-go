package pokecache

import "time"

// NOTE: The purpose of this file is to cache the locations of the map call
// That way a network call does not need to be made to the PokeAPI each time
// the user needs to navigate to a page

type Cache struct {
    cache map[string] cacheEntry
}

type cacheEntry struct {
    // caching HTTP response bodies
    val []byte
    createdAt time.Time
}

func NewCache() Cache {
    return Cache {
        cache: make(map[string]cacheEntry),
    }
}

func (c *Cache) Add(key string, val []byte) {
    c.cache[key] = cacheEntry {
        val: val,
        createdAt: time.Now().UTC(),
    }

}

func (c *Cache) Get(key string) ([]byte, bool) {
    cacheE, ok := c.cache[key]
    return cacheE.val, ok
}
