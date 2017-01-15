package model

import "github.com/nairufan/yh-share/util"

type TopRecord struct {
	MetaFields                         `bson:",inline"`
	UserId        string               `bson:"userId"  json:"-"`
	DocumentNames []string             `bson:"documentNames" json:"documentNames"`
	Records       []*util.TopRecord     `bson:"records" json:"records"`
}















