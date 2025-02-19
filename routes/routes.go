package routes

import (
	"web_test01/handlers"

	"github.com/gofiber/fiber/v2"
)

type Apiroutes struct {
	Handlers *handlers.Handlers
}

func (r *Apiroutes) Mount(router *fiber.App) {
	api := router.Group("/api")

	// define v1
	v1 := api.Group("/v1")

	// define health route
	v1.Route("/health", r.setHealthRoute)

	// define the user routes
	v1.Route("/user", r.setUserRoutes)
}
