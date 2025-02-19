package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Apiroutes) setHealthRoute(router fiber.Router) {
	router.Get("/", r.Handlers.Health.GetHealth)
}
