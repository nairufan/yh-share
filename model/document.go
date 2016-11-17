package model

type Document struct {
	MetaFields              `bson:",inline"`
	UserId string         `bson:"userId"  json:"userId"`
	Title  string         `bson:"title"  json:"title"`
}















