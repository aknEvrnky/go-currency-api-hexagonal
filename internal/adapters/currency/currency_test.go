package currency

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/cache"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const xmlResponse = `<Tarih_Date>
	<Currency CrossOrder="0" Kod="USD" CurrencyCode="USD">
		<Unit>1</Unit>
		<Isim>US DOLLAR</Isim>
		<CurrencyName>US DOLLAR</CurrencyName>
		<ForexBuying>1.0</ForexBuying>
		<ForexSelling>1.2</ForexSelling>
		<BanknoteBuying>1.1</BanknoteBuying>
		<BanknoteSelling>1.3</BanknoteSelling>
	</Currency>
</Tarih_Date>`

func TestItCanGetAllCurrencies(t *testing.T) {
	// mock the http server
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(xmlResponse))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Create a new in-memory cache adapter for testing
	cacheAdapter := cache.NewInMemoryCacheAdapter()
	adapter := NewAdapter(cacheAdapter)

	// Override ENDPOINT with the mock server URL
	adapter.endpoint = server.URL

	currencies, err := adapter.GetList()

	assert.NoError(t, err)
	assert.Len(t, currencies, 1)
	assert.Equal(t, "USD", currencies[0].Code)
	assert.Equal(t, "US DOLLAR", currencies[0].Title)
	assert.Equal(t, uint(1), currencies[0].Unit)
	assert.Equal(t, 1.1, currencies[0].BuyingRate)
	assert.Equal(t, 1.3, currencies[0].SellingRate)
}

func TestItCanGetAllCurrenciesFromCache(t *testing.T) {
	// Create a new in-memory cache adapter for testing
	cacheAdapter := cache.NewInMemoryCacheAdapter()
	adapter := NewAdapter(cacheAdapter)

	// Set the cache value
	cacheAdapter.Set("currencies", 10, []domain.Currency{
		{
			Code:        "USD",
			Title:       "US DOLLAR",
			Unit:        1,
			BuyingRate:  1.1,
			SellingRate: 1.3,
		},
	})

	currencies, err := adapter.GetList()

	assert.NoError(t, err)
	assert.Len(t, currencies, 1)
	assert.Equal(t, "USD", currencies[0].Code)
	assert.Equal(t, "US DOLLAR", currencies[0].Title)
	assert.Equal(t, uint(1), currencies[0].Unit)
	assert.Equal(t, 1.1, currencies[0].BuyingRate)
	assert.Equal(t, 1.3, currencies[0].SellingRate)
}

func TestItCanGetAllCurrenciesFromCacheWhenCacheExpired(t *testing.T) {
	// mock the http server
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(xmlResponse))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Create a new in-memory cache adapter for testing
	cacheAdapter := cache.NewInMemoryCacheAdapter()
	adapter := NewAdapter(cacheAdapter)

	// Override ENDPOINT with the mock server URL
	adapter.endpoint = server.URL

	// Set the cache value
	cacheAdapter.Set("currencies", 1, []domain.Currency{
		{
			Code:        "USD",
			Title:       "US DOLLAR",
			Unit:        1,
			BuyingRate:  1.1,
			SellingRate: 1.3,
		},
	})

	// Wait for the cache to expire
	<-time.After(2 * time.Second)

	currencies, err := adapter.GetList()

	assert.NoError(t, err)
	assert.Len(t, currencies, 1)
	assert.Equal(t, "USD", currencies[0].Code)
	assert.Equal(t, "US DOLLAR", currencies[0].Title)
	assert.Equal(t, uint(1), currencies[0].Unit)
	assert.Equal(t, 1.1, currencies[0].BuyingRate)
	assert.Equal(t, 1.3, currencies[0].SellingRate)
}

func TestItCanHandleApiError(t *testing.T) {
	// mock the http server
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Create a new in-memory cache adapter for testing
	cacheAdapter := cache.NewInMemoryCacheAdapter()
	adapter := NewAdapter(cacheAdapter)

	// Override ENDPOINT with the mock server URL
	adapter.endpoint = server.URL

	currencies, err := adapter.GetList()

	assert.ErrorIs(t, err, ErrApiError)
	assert.Nil(t, currencies)
}

func TestAdapter_GetByCurrencyCode(t *testing.T) {
	// Mock HTTP server
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(xmlResponse))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Create a new in-memory cache adapter for testing
	cacheAdapter := cache.NewInMemoryCacheAdapter()
	adapter := NewAdapter(cacheAdapter)

	// Override ENDPOINT with the mock server URL
	adapter.endpoint = server.URL

	currency, err := adapter.GetByCurrencyCode("USD")

	assert.NoError(t, err)
	assert.Equal(t, "USD", currency.Code)
	assert.Equal(t, "US DOLLAR", currency.Title)
	assert.Equal(t, uint(1), currency.Unit)
	assert.Equal(t, 1.1, currency.BuyingRate)
	assert.Equal(t, 1.3, currency.SellingRate)
}

func TestAdapter_GetByCurrencyCodeFromCache(t *testing.T) {
	// Create a new in-memory cache adapter for testing
	cacheAdapter := cache.NewInMemoryCacheAdapter()
	adapter := NewAdapter(cacheAdapter)

	// Set the cache value
	cacheAdapter.Set("currencies", 10, []domain.Currency{
		{
			Code:        "USD",
			Title:       "US DOLLAR",
			Unit:        1,
			BuyingRate:  1.1,
			SellingRate: 1.3,
		},
	})

	currency, err := adapter.GetByCurrencyCode("USD")

	assert.NoError(t, err)
	assert.Equal(t, "USD", currency.Code)
	assert.Equal(t, "US DOLLAR", currency.Title)
	assert.Equal(t, uint(1), currency.Unit)
	assert.Equal(t, 1.1, currency.BuyingRate)
	assert.Equal(t, 1.3, currency.SellingRate)
}
