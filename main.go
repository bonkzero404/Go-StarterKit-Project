package main

import (
	"fmt"
	"go-starterkit-project/config"
	"go-starterkit-project/database"
	"go-starterkit-project/database/driver"

	appRoute "go-starterkit-project/app"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
)

func main() {
	// Fiber app
	app := fiber.New()

	// Call database connection
	driver.ConnectDB()

	// connect casbin
	driver.ConnectCasbin()

	// Casbin load policy
	// driver.CASBIN.LoadPolicy()

	// driver.CASBIN.Enforce("alice", "data1", "read")

	// // Modify the policy.
	// // e.AddPolicy(...)
	// // e.RemovePolicy(...)

	// // Save the policy back to DB.
	// driver.CASBIN.SavePolicy()

	// Auto migration table
	database.MigrateDB()

	// Handling global cors
	app.Use(cors.New())

	// Securing with helmet
	app.Use(helmet.New())

	// Call bootstrap all module
	appRoute.Bootstrap(app)

	// Set port
	appPort := fmt.Sprintf(":%s", config.Config("APP_PORT"))

	// Listen app
	app.Listen(appPort)
}
