package main

import (
	"context"
	"waste-tracker-web-service/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		// write a todo to database
		sampleDoc := bson.M{"name": "sample todo"}

		collection := database.GetCollection("todos")

		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
		}

		// send down info about the todo
		return c.JSON(nDoc)
	})

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen(":3000")
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
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}
