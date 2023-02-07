package main

import (
	"os"
	"waste-tracker-web-service/database"

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
