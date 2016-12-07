package model

type Sequence struct {
	Id    string  `bson:"_id"  json:"id"`
	Value int64   `bson:"value"  json:"-"`
}
