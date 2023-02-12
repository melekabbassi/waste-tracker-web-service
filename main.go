package main

import (
	"os"
	"waste-tracker-web-service/database"
	"waste-tracker-web-service/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer close database
	defer database.CloseMongoDB()

	// setup fiber
	app := generateApp()

	// get the port from the env
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	app := fiber.New()

	// create healthcheck route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the dumptruck group and routes
	dumptruckGroup := app.Group("/dumptrucks")
	dumptruckGroup.Get("/", handlers.GetDumptrucks)
	dumptruckGroup.Post("/", handlers.CreateDumptruck)
	dumptruckGroup.Delete("/:id", handlers.DeleteDumptruck)
	dumptruckGroup.Post("/route", handlers.CreateRoute)

	// create the landfillsite group and routes
	landfillsiteGroup := app.Group("/landfillsites")
	landfillsiteGroup.Get("/", handlers.GetLandfillSites)
	landfillsiteGroup.Post("/", handlers.CreateLandfillSite)

	return app
}

func initApp() error {
	// setup evn
	err := loadENV()
	if err != nil {
		return err
	}

	// setup database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadENV() error {
	goENV := os.Getenv("GO_ENV")
	if goENV == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
