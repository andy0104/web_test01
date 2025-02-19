package handlers_test

import (
	"net/http"
	"testing"
	app_test "web_test01/tests/app"

	"github.com/gofiber/fiber/v2"
)

func TestGetHealth(t *testing.T) {
	fiberApp := app_test.InitializeApp()

	t.Run("GetHealth handler success", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/v1/health", nil)
		if err != nil {
			t.Fatal("Test failed", err)
		}

		res, err := fiberApp.Test(req)
		if err != nil {
			t.Fatal("Test failed", err)
		}
		if res.StatusCode != fiber.StatusOK {
			t.Errorf("Expected status %d but got status %d", fiber.StatusOK, res.StatusCode)
		}
	})
}
