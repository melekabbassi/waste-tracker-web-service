package main

import (
	"waste-tracker-web-service/handlers"

	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	app := fiber.New()

	// create healthcheck route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the dumptruck group and routes
	dumptruckGroup := app.Group("/dumptrucks")
	dumptruckGroup.Get("/", handlers.TestHandler)

	return app
}
