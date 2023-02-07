package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type dumptruck struct {
  ID            primitive.ObjectID  `json:"id" bson:"_id"`
  Completed     bool                `json:"completed" bson:"completed`
  Date          string              `json:"date" bson:"date"`
  Driver        string              `json:"driver" bson:"driver"`
  Location      string              `json:"location" bson:"location"`
  Path          string              `json:"path" bson:"path"`
}
