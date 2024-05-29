package pokecache

/* objectives of this package
1. cacheEntry is a struct that contains the result of a API request with record of creation time
   entries is the map that maps the URL string to the cache
2. if the query was already made, result is retrived from the cache
   otherwise make a new query and add result to the cache
3. use mutex to protect cacheEntry map because map is not thread-safe in Go
4. remove older entries periodically by using a time ticker

byte slice
- array of character bytes - can represent a string */

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	fmt.Println("NewCache creating")
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	// once a cache is created we spawn a goroutine that uses a ticker to delete older entries
	go c.reapLoop(interval)

	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	// C is a channel of type Time that delivers "ticks" of "a clock" at intervals
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	// delete older cache entries
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			//fmt.Println("Cache entry deleted")
			delete(c.cache, k)
		}
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
	fmt.Println(string("\033[31m"), "Cache entry added", string("\033[0m"))
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	if ok {
		fmt.Println(string("\033[32m"), "Cache entry retrieved", string("\033[0m"))
	}
	return val.val, ok
}
