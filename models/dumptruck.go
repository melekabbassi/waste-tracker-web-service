package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dumptruck struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Location string             `json:"location" bson:"location"`
	Status   string             `json:"status" bson:"status"`
	Route    string             `json:"route" bson:"route"`
}
