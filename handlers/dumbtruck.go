package handlers

import (
	"context"

	"waste-tracker-web-service/database"
	"waste-tracker-web-service/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type DumptruckDTO struct {
	Location string `json:"location" bson:"location"`
	Status   string `json:"status" bson:"status"`
	Route    string `json:"route" bson:"route"`
}

// POST /dumptrucks
func CreateDumptruck(c *fiber.Ctx) error {
	nDumptruck := new(DumptruckDTO)

	if err := c.BodyParser(nDumptruck); err != nil {
		return err
	}

	dumptruckCollection := database.GetCollection("dumptrucks")
	nDoc, err := dumptruckCollection.InsertOne(context.TODO(), nDumptruck)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}

// GET /dumptrucks
func GetDumptrucks(c *fiber.Ctx) error {
	dumptruckCollection := database.GetCollection("dumptrucks")
	cursor, err := dumptruckCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	var dumptrucks []models.Dumptruck
	if err = cursor.All(context.TODO(), &dumptrucks); err != nil {
		return err
	}

	return c.JSON(dumptrucks)
}

// DELETE /dumptrucks/:id
func DeleteDumptruck(c *fiber.Ctx) error {
	dumptruckCollection := database.GetCollection("dumptrucks")
	id := c.Params("id")
	_, err := dumptruckCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
