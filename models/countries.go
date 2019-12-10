package models

import "gopkg.in/mgo.v2/bson"

type Country struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	AllowedVisa []string      `json:"allowedvisa" bson:"allowedvisa"`
}
