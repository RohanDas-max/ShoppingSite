package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rohandas-max/ambassador/src/database"
	"github.com/rohandas-max/ambassador/src/router"
)

func main() {

	database.Connection()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.Setup(app)

	app.Listen(":3000")
}
