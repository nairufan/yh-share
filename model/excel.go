package model

type Excel struct {
	MetaFields            `bson:",inline"`
	UserId   string       `bson:"userId"  json:"userId"`
	BatchKey string       `bson:"batchKey"  json:"batchKey"`
	Tel      string       `bson:"tel" json:"tel"`
	Name     string       `bson:"name" json:"name"`
	Data     interface{} `bson:"data" json:"data"`
}















