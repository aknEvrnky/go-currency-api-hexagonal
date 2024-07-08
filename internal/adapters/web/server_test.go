package web

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
	"testing"
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
