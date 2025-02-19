package handlers_test

import (
	"net/http"
	"testing"
	app_test "web_test01/tests/app"

	"github.com/gofiber/fiber/v2"
)

func TestUserProfileHandler(t *testing.T) {
	fiberApp := app_test.InitializeApp()

	t.Run("Test /api/v1/user/profile api call without jwt token", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/v1/user/profile", nil)
		if err != nil {
			t.Fatal("Test fail:", err)
		}

		res, err := fiberApp.Test(req)
		if err != nil {
			t.Fatal("Test fail:", err)
		}

		if res.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("Expected status code %d, but received %d", fiber.StatusUnauthorized, res.StatusCode)
		}
	})

	t.Run("Test /api/v1/user/profile api call with invalid jwt token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/user/profile", nil)
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIiwiZXhwIjoxNzM5NTE3MDAzLCJpYXQiOjE3Mzk1MTY0MDMsImlzcyI6IndlYl90ZXN0MDEiLCJzdWIiOiI2In0.R_tHUShnOsKEiKXhGEvoC16dyV1VUBNGuA9WWaz0dzo")

		res, err := fiberApp.Test(req)
		if err != nil {
			t.Fatal("Test fail:", err)
		}

		if res.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("Expected status code %d, but received %d", fiber.StatusUnauthorized, res.StatusCode)
		}
	})
}
