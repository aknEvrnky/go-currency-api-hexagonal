package ports

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/cache"
	"time"
)

type CachePort interface {
	Get(key string) (cache.Value, error)
	Set(key string, ttl time.Duration, value cache.Value) error
	Remove(key string) error
	Exists(key string) (bool, error)

	Remember(key string, ttl time.Duration, f func() (cache.Value, error)) (cache.Value, error)
}
