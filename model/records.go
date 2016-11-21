package model

type Record struct {
	MetaFields                 `bson:",inline"`
	DocumentId  string         `bson:"documentId"  json:"documentId"`
	QueryField1 string         `bson:"queryField1" json:"queryField1"`
	QueryField2 string         `bson:"queryField2" json:"queryField2"`
	Data        []string       `bson:"data" json:"data"`
}















