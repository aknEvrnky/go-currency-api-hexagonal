package api

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/currency"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
	"testing"
)

var currencies = []domain.Currency{
	{
		Code:        "USD",
		Title:       "US Dollar",
		Unit:        1,
		BuyingRate:  33,
		SellingRate: 35,
	},
	{
		Code:        "EUR",
		Title:       "EURO",
		Unit:        1,
		BuyingRate:  35,
		SellingRate: 38,
	},
}

type currencyMock struct {
	mock.Mock
}

func (c *currencyMock) GetList() ([]domain.Currency, error) {
	return currencies, nil
}

func (c *currencyMock) GetByCurrencyCode(currencyCode string) (domain.Currency, error) {
	for _, mockCurrency := range currencies {
		if mockCurrency.Code == currencyCode {
			return mockCurrency, nil
		}
	}
	return domain.Currency{}, currency.ErrCurrencyNotFound
}

var application = NewApplication(&currencyMock{})

func TestItGetsAllCurrencies(t *testing.T) {
	c := currencyMock{}
	c.On("GetList").Return(currencies, nil)

	_, err := application.GetAllCurrencies()
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}
}

func TestItGetsCurrencyByCode(t *testing.T) {
	c := currencyMock{}
	c.On("GetByCurrencyCode", "USD").Return(currencies[0], nil)

	_, err := application.GetCurrencyByCode("USD")
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}
}
