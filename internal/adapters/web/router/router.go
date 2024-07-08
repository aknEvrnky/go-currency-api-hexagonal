package router

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/web/handler"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, apiAdapter ports.ApiPort) {
	// Middleware
	api := app.Group("/api", logger.New())

	webHandler := handler.NewHandler(apiAdapter)

	api.Get("/currencies", webHandler.GetAllCurrencies)
	api.Get("/currencies/:code", webHandler.GetCurrency)
}
