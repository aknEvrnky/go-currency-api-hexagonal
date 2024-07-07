package ports

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
)

type CurrencyPort interface {
	GetList() ([]domain.Currency, error)
	GetByCurrencyCode(currencyCode string) (domain.Currency, error)
}
