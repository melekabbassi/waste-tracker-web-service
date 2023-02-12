package handlers

import (
	"waste-tracker-web-service/database"
	"waste-tracker-web-service/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type RouteDTO struct {
	StartLocation string   `json:"startLocation" bson:"startLocation"`
	EndLocation   string   `json:"endLocation" bson:"endLocation"`
	WayPoints     []string `json:"wayPoints" bson:"wayPoints"`
	Distance      float64  `json:"distance" bson:"distance"`
	StartTime     string   `json:"startTime" bson:"startTime"`
	EndTime       string   `json:"endTime" bson:"endTime"`
	DumptruckId   string   `json:"dumptruckId" bson:"dumptruckId"`
}

// POST /routes
func CreateRoute(c *fiber.Ctx) error {
	createData := new(RouteDTO)

	if err := c.BodyParser(createData); err != nil {
		return err
	}

	// get the collection reference
	coll := database.GetCollection("dumptrucks")

	// get the filter
	filter := bson.M{"_id": createData.DumptruckId}
	nRouteData := models.Route{
		StartLocation: createData.StartLocation,
		EndLocation:   createData.EndLocation,
		WayPoints:     createData.WayPoints,
		Distance:      createData.Distance,
		StartTime:     createData.StartTime,
		EndTime:       createData.EndTime,
	}
	updatePayload := bson.M{"$push": bson.M{"routes": nRouteData}}

	// update the document
	_, err := coll.UpdateOne(c.Context(), filter, updatePayload)
	if err != nil {
		return err
	}

	return c.SendString("Route created")
}
