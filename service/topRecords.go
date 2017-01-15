package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionTopRecords = "top_records"
)

func AddTopRecords(document *model.TopRecord) *model.TopRecord {
	time := time.Now()
	document.Id = model.NewId()
	document.CreatedTime = &time
	session := mongo.Get()
	defer session.Close()
	session.MustInsert(collectionTopRecords, document)
	return document
}

func TopRecordsList(userId string, offset int, limit int) []*model.TopRecord {
	session := mongo.Get()
	defer session.Close()
	topRecords := []*model.TopRecord{}

	option := mongo.Option{
		Sort: []string{"-createdTime"},
		Limit: &limit,
		Offset: &offset,
	}
	session.MustFindWithOptions(collectionTopRecords, bson.M{"userId": userId}, option, &topRecords)
	return topRecords
}