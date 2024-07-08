package web

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockApiAdapter struct {
	mock.Mock
}

func (m *mockApiAdapter) GetAllCurrencies() ([]domain.Currency, error) {
	args := m.Called()
	return args.Get(0).([]domain.Currency), args.Error(1)
}

func (m *mockApiAdapter) GetCurrencyByCode(code string) (domain.Currency, error) {
	args := m.Called(code)
	return args.Get(0).(domain.Currency), args.Error(1)
}

func TestItCanInitializeANewAdapter(t *testing.T) {
	api := &mockApiAdapter{}
	adapter := NewAdapter(api, 3000)

	if adapter.port != 3000 {
		t.Errorf("Expected port to be 3000, got %d", adapter.port)
	}
}

func TestAdapter_Run(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Mock ApiPort implementation
	api := &mockApiAdapter{}

	// Create a new Adapter with the mock ApiPort
	adapter := NewAdapter(api, 3000)
	adapter.FiberApp = app

	var err error
	// Run the server in a separate goroutine
	go func() {
		if err = adapter.Run(); err != nil {
			t.Errorf("Error running the server: %v", err)
		}
	}()
	defer adapter.Shutdown()

	// Allow the server to start
	<-time.After(100 * time.Millisecond)

	assert.Nil(t, err)
}

func TestAdapter_Shutdown(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Mock ApiPort implementation
	api := &mockApiAdapter{}

	// Create a new Adapter with the mock ApiPort
	adapter := NewAdapter(api, 3000)
	adapter.FiberApp = app

	var err error
	// Run the server in a separate goroutine
	go func() {
		if err = adapter.Run(); err != nil {
			t.Errorf("Error running the server: %v", err)
		}
	}()

	// Allow the server to start
	<-time.After(100 * time.Millisecond)

	assert.Nil(t, err)

	// Shutdown the server
	err = adapter.Shutdown()
	assert.NoError(t, err)
}
