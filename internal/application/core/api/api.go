package api

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/ports"
)

type Application struct {
	currency ports.CurrencyPort
}

func NewApplication(currency ports.CurrencyPort) *Application {
	return &Application{
		currency: currency,
	}
}

func (a Application) GetAllCurrencies() ([]domain.Currency, error) {
	return a.currency.GetList()
}

func (a Application) GetCurrencyByCode(code string) (domain.Currency, error) {
	return a.currency.GetByCurrencyCode(code)
}
