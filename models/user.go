package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name"  bson:"name"`
	// Email  string        `json:"email" bson:"email"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}
