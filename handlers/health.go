package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type HealthHandler struct {
	logger *zap.SugaredLogger
}

func (h *HealthHandler) GetHealth(c *fiber.Ctx) error {
	h.logger.Infow("Health Handler", "Method", c.Method(), "Request", c.Path())
	return c.JSON(fiber.Map{
		"message": "ok",
	})
}
