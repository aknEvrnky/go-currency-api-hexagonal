package ports

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
)

type ApiPort interface {
	GetAllCurrencies() ([]domain.Currency, error)
	GetCurrencyByCode(code string) (domain.Currency, error)
}
