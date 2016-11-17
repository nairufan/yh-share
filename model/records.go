package model

type Excel struct {
	MetaFields                `bson:",inline"`
	DocumentId string         `bson:"documentId"  json:"documentId"`
	Tel        string         `bson:"tel" json:"tel"`
	Name       string         `bson:"name" json:"name"`
	Data       []string       `bson:"data" json:"data"`
}















