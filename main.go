package main

import (
	"fmt"
	"go-boilerplate-clean-arch/config"
	"go-boilerplate-clean-arch/infrastructure/database"

	appRoute "go-boilerplate-clean-arch/app"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/helmet/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	database.MigrateDB()
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Get("/monitor", monitor.New())
	appRoute.Bootstrap(app)
	appPort := fmt.Sprintf(":%s", config.Config("APP_PORT"))

	app.Listen(appPort)
}
