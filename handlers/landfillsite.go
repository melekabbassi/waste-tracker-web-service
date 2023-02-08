package handlers

import (
	"context"
	"waste-tracker-web-service/database"
	"waste-tracker-web-service/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type LandfillSiteDTO struct {
	Name     string `json:"name" bson:"name"`
	Location string `json:"location" bson:"location"`
	Capacity int    `json:"capacity" bson:"capacity"`
	Status   string `json:"status" bson:"status"`
}

// POST /landfillsites
func CreateLandfillSite(c *fiber.Ctx) error {
	nLandfillSite := new(LandfillSiteDTO)

	if err := c.BodyParser(nLandfillSite); err != nil {
		return err
	}

	landfillSiteCollection := database.GetCollection("landfillsites")
	nDoc, err := landfillSiteCollection.InsertOne(context.TODO(), nLandfillSite)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}

// GET /landfillsites
func GetLandfillSites(c *fiber.Ctx) error {
	landfillSiteCollection := database.GetCollection("landfillsites")
	cursor, err := landfillSiteCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	var landfillSites []models.LandFillSite
	if err = cursor.All(context.TODO(), &landfillSites); err != nil {
		return err
	}

	return c.JSON(landfillSites)
}
