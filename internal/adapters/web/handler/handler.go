package handler

import (
	"github.com/aknEvrnky/currency-api-hexogonal/internal/ports"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	api ports.ApiPort
}

func NewHandler(api ports.ApiPort) *Handler {
	return &Handler{
		api: api,
	}
}

func (h *Handler) GetAllCurrencies(c *fiber.Ctx) error {
	currencies, err := h.api.GetAllCurrencies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(currencies)
}

func (h *Handler) GetCurrency(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("code is required")
	}

	currency, err := h.api.GetCurrencyByCode(code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(currency)
}
