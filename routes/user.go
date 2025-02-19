package routes

import (
	"web_test01/middlewares/auth"

	"github.com/gofiber/fiber/v2"
)

func (r *Apiroutes) setUserRoutes(router fiber.Router) {
	router.Post("/register", r.Handlers.User.Register)

	router.Post("/login", r.Handlers.User.Login)

	router.Get("/profile", auth.Authenticate, r.Handlers.User.Profile)

	router.Post("/coverphoto", auth.Authenticate, r.Handlers.User.Coverpic)
}
