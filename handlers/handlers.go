package handlers

import (
	"web_test01/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type Handlers struct {
	Health interface {
		GetHealth(*fiber.Ctx) error
	}
	User interface {
		Register(*fiber.Ctx) error
		Login(*fiber.Ctx) error
		Profile(*fiber.Ctx) error
		Coverpic(*fiber.Ctx) error
	}
}

func NewHandlers(logger *zap.SugaredLogger, services services.Services) Handlers {
	return Handlers{
		Health: &HealthHandler{logger: logger},
		User:   &UserHandler{logger: logger, services: services},
	}
}
