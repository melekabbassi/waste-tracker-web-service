package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Route struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	StartLocation string             `json:"startLocation" bson:"startLocation"`
	EndLocation   string             `json:"endLocation" bson:"endLocation"`
	WayPoints     []string           `json:"wayPoints" bson:"wayPoints"`
	Distance      float64            `json:"distance" bson:"distance"`
	StartTime     string             `json:"startTime" bson:"startTime"`
	EndTime       string             `json:"endTime" bson:"endTime"`
	Dumptrucks    []string           `json:"dumptrucks" bson:"dumptrucks"`
}
