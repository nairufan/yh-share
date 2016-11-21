package model

type Document struct {
	MetaFields                   `bson:",inline"`
	UserId        string         `bson:"userId"  json:"userId"`
	Title         string         `bson:"title"  json:"title"`
	TitleFields   []string       `bson:"titleFields"  json:"titleFields"`
	DisplayColumn []int          `bson:"displayColumn"  json:"displayColumn"`
}















