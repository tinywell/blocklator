package store

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"sync"
	"time"
)

// clean params
const (
	Experied     = time.Minute * 10
	ScanInterval = time.Minute
	MaxCacheSize = 150 * 1024 * 1024 // byte
)

// Cache cache
type Cache interface {
	Store(key string, size int64, data interface{}) error
	Load(key string) (interface{}, error)
	KeyGen(raw []byte) (string, error)
}

// MapCache cache base map
type MapCache struct {
	cache     *sync.Map
	create    map[string]time.Time
	size      map[string]int64
	stop      chan struct{}
	cacheLock *sync.RWMutex
}

// NewCache return new mapcache
func NewCache() *MapCache {
	cache := &MapCache{
		cache:     &sync.Map{},
		create:    make(map[string]time.Time),
		stop:      make(chan struct{}),
		size:      make(map[string]int64),
		cacheLock: &sync.RWMutex{},
	}
	go cache.scan()
	return cache
}

// KeyGen genarate a cache key by cache data
func (c *MapCache) KeyGen(raw []byte) (string, error) {
	sha := sha256.New()
	_, err := sha.Write(raw)
	if err != nil {
		return "", err
	}
	res := sha.Sum([]byte{})
	return base64.StdEncoding.EncodeToString(res), nil
}

// PreStore store prepare
func (c *MapCache) preStore(key string, size int64) {
	if _, ok := c.create[key]; ok {
		return
	}
	var total int64 = 0
	for _, v := range c.size {
		total += v
	}
	fmt.Println("total size:", total)
	if total+size > MaxCacheSize {
		c.cacheLock.Lock()
		c.create = make(map[string]time.Time)
		c.cache = &sync.Map{}
		c.cacheLock.Unlock()
		fmt.Println("clear")
	}
	c.size[key] = size
}

// Store cache data in MapCache,return key
func (c *MapCache) Store(key string, size int64, data interface{}) error {
	c.preStore(key, size)
	c.cacheLock.RLock()
	defer c.cacheLock.RUnlock()
	if _, ok := c.create[key]; ok {
		return nil
	}
	c.cache.Store(key, data)
	c.create[key] = time.Now()
	return nil
}

// Load load data by key
func (c *MapCache) Load(key string) (interface{}, error) {
	value, ok := c.cache.Load(key)
	if ok {
		return value, nil
	}
	return nil, errors.New("no data")
}

func (c *MapCache) scan() {
	tick := time.NewTicker(ScanInterval)
	for {
		select {
		case <-tick.C:
			go c.clean()
		case <-c.stop:
			tick.Stop()
			return
		}
	}
}
func (c *MapCache) clean() {
	for k, v := range c.create {
		if time.Now().Sub(v) > Experied {
			c.cache.Delete(k)
			delete(c.create, k)
		}
	}
}
