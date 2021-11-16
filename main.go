package main

import (
	"fmt"
	"go-boilerplate-clean-arch/config"
	"go-boilerplate-clean-arch/infrastructure/database"

	appRoute "go-boilerplate-clean-arch/app"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	database.MigrateDB()
	appRoute.Bootstrap(app)
	appPort := fmt.Sprintf(":%s", config.Config("APP_PORT"))

	app.Listen(appPort)
}
