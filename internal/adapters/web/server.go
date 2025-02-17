package web

import (
	"fmt"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/adapters/web/router"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/ports"
	"github.com/gofiber/fiber/v2"
)

type Adapter struct {
	Api      ports.ApiPort
	port     int
	FiberApp *fiber.App
}

func NewAdapter(api ports.ApiPort, port int) *Adapter {
	return &Adapter{
		Api:      api,
		port:     port,
		FiberApp: fiber.New(),
	}
}

func (a *Adapter) Run() error {
	router.SetupRoutes(a.FiberApp, a.Api)

	return a.FiberApp.Listen(fmt.Sprintf(":%d", a.port))
}

func (a *Adapter) Shutdown() error {
	return a.FiberApp.Shutdown()
}
