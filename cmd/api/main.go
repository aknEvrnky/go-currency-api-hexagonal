package main

import (
	"github.com/aknEvrnky/currency-api-hexogonal/config"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/currency"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/web"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/api"
)

func main() {
	currencyAdapter := currency.NewAdapter()

	apiAdapter := api.NewApplication(currencyAdapter)

	server := web.NewAdapter(apiAdapter, config.GetApplicationPort())

	server.Run()
}
