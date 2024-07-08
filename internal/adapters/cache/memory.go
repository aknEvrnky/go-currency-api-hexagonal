package cache

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/cache"
	"sync"
	"time"
)

type InMemoryCacheAdapter struct {
	cache map[string]cache.Cache
	mutex sync.RWMutex
}

func NewInMemoryCacheAdapter() *InMemoryCacheAdapter {
	return &InMemoryCacheAdapter{
		cache: make(map[string]cache.Cache),
	}
}

func (i *InMemoryCacheAdapter) Get(key string) (cache.Value, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	cch, ok := i.cache[key]
	if !ok {
		return cache.Cache{}, nil
	}

	if time.Now().After(cch.Expiry) {
		delete(i.cache, key)
		return cache.Cache{}, nil
	}

	return cch.Value, nil
}

func (i *InMemoryCacheAdapter) Set(key string, ttl time.Duration, value cache.Value) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.cache[key] = cache.Cache{
		Expiry: time.Now().Add(ttl),
		Value:  value,
	}

	return nil
}

func (i *InMemoryCacheAdapter) Remove(key string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	delete(i.cache, key)
	return nil
}

func (i *InMemoryCacheAdapter) Exists(key string) (bool, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	_, ok := i.cache[key]
	return ok, nil
}

func (i *InMemoryCacheAdapter) Remember(key string, ttl time.Duration, f func() (cache.Value, error)) (cache.Value, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if cch, ok := i.cache[key]; ok {
		if time.Now().Before(cch.Expiry) {
			return cch.Value, nil
		} else {
			delete(i.cache, key)
		}
	}

	val, err := f()
	if err != nil {
		return cache.Cache{}, err
	}

	i.cache[key] = cache.Cache{
		Expiry: time.Now().Add(ttl),
		Value:  val,
	}
	return val, nil
}
