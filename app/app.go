package app

import (
	"web_test01/handlers"
	"web_test01/routes"
	"web_test01/services"
	"web_test01/storage"
	"web_test01/utility/database"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func InitializeApp() *fiber.App {
	// initialize the logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// initialize the database connection
	db, err := database.NewConnection()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	// initialize the storage
	store := storage.NewStorage(db, logger)

	// initialize the services
	services := services.NewServices(logger, store)

	// initialize handlers
	handlers := handlers.NewHandlers(logger, services)

	routes := routes.Apiroutes{
		Handlers: &handlers,
	}

	routes.Mount(app)

	return app
}
