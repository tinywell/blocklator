package store

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sync"
	"time"
)

// clean params
const (
	Experied     = time.Minute * 60
	ScanInterval = time.Minute
)

// Cache cache
type Cache interface {
	Store(key string, data interface{}) error
	Load(key string) (interface{}, error)
	KeyGen(raw []byte) (string, error)
}

// MapCache cache base map
type MapCache struct {
	cache  *sync.Map
	create map[string]time.Time
	stop   chan struct{}
}

// NewCache return new mapcache
func NewCache() *MapCache {
	cache := &MapCache{
		cache:  &sync.Map{},
		create: make(map[string]time.Time),
		stop:   make(chan struct{}),
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

// Store cache data in MapCache,return key
func (c *MapCache) Store(key string, data interface{}) error {
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
		}
	}
}
