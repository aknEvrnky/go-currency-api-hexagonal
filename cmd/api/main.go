package main

import (
	"context"
	"errors"
	"github.com/aknEvrnky/currency-api-hexogonal/config"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/cache"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/currency"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/web"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/api"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	currencyAdapter := currency.NewAdapter(cache.NewInMemoryCacheAdapter())
	apiAdapter := api.NewApplication(currencyAdapter)
	server := web.NewAdapter(apiAdapter, config.GetApplicationPort())

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal.")
	if err := server.Shutdown(); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
	}

}
