package model

type TopDocument struct {
	MetaFields                   `bson:",inline"`
	UserId string                `bson:"userId"  json:"userId"`
	Title  string                `bson:"title"  json:"title"`
	Fields [][]string            `bson:"fileds"  json:"fileds"`
}















