package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type MetaFields struct {
	Id          string     `bson:"_id"  json:"id"`
	CreatedTime *time.Time `bson:"createdTime,omitempty"`
}

func NewId() string {
	return bson.NewObjectId().Hex()
}