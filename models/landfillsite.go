package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LandFillSite struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Location string             `json:"location" bson:"location"`
	Capacity int                `json:"capacity" bson:"capacity"`
	Status   string             `json:"status" bson:"status"`
}
