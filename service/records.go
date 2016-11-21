package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
)

const (
	collectionRecords = "records"
)

func AddRecords(records []*model.Record, documentId string) []*model.Record {
	if documentId == "" {
		beego.Error("Document id must not empty")
		panic("500")
	}

	time := time.Now()
	recordInterfaces := make([]interface{}, len(records))
	for index, record := range records {
		record.Id = model.NewId()
		record.DocumentId = documentId
		record.CreatedTime = &time
		recordInterfaces[index] = record
	}
	session := mongo.Get()
	defer session.Close()
	session.MustInsert(collectionRecords, recordInterfaces...)
	return records
}

func Search(documentId string, query string) []*model.Record {
	session := mongo.Get()
	defer session.Close()
	records := []*model.Record{}
	session.MustFind(collectionRecords, bson.M{"$or": []bson.M{bson.M{"queryField1": query}, bson.M{"queryField2": query} }, "documentId": documentId}, &records)
	return records
}
