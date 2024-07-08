package cache

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/cache"
	"testing"
	"time"
)

func TestItCanGetAValueFromCache(t *testing.T) {
	var memoryCache = NewInMemoryCacheAdapter()

	memoryCache.Set("key", 10, "value")
	value, err := memoryCache.Get("key")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if value != "value" {
		t.Errorf("Expected value, got %v", value)
	}
}

func TestItCanSetAValueToCache(t *testing.T) {
	var memoryCache = NewInMemoryCacheAdapter()
	err := memoryCache.Set("key", 10, "foo")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	value, _ := memoryCache.Get("key")
	if value != "foo" {
		t.Errorf("Expected foo, got %v", value)
	}
}

func TestItCanForgetAValueFromCache(t *testing.T) {
	var memoryCache = NewInMemoryCacheAdapter()
	memoryCache.Set("key", 10, "value")

	memoryCache.Remove("key")
	val, _ := memoryCache.Get("key")
	if val == nil {
		t.Errorf("Expected nil, got %s", val)
	}
}

func TestItCanCheckWhetherValueExistsOrNot(t *testing.T) {
	var memoryCache = NewInMemoryCacheAdapter()
	memoryCache.Set("key", 10*time.Millisecond, "value")

	exists, _ := memoryCache.Exists("key")
	if !exists {
		t.Errorf("Expected true, got %v", exists)
	}

	exists, _ = memoryCache.Exists("foo")
	if exists {
		t.Errorf("Expected false, got %v", exists)
	}
}

func TestItCanRememberValues(t *testing.T) {
	var memoryCache = NewInMemoryCacheAdapter()
	numberOfCalls := 0
	var val cache.Value
	var err error

	for i := 0; i < 3; i++ {
		val, err = memoryCache.Remember("key", 10*time.Second, func() (cache.Value, error) {
			numberOfCalls++
			return "remembered_value", nil
		})
	}

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if val != "remembered_value" {
		t.Errorf("Expected remembered_value, got %v", val)
	}

	if numberOfCalls != 1 {
		t.Errorf("Expected 1 calss, got %v", numberOfCalls)
	}
}
