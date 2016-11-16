package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionExcel = "excel"
)

func AddRecords(records []*model.Excel) []*model.Excel {
	key := model.NewId()
	time := time.Now()
	recordInterfaces := make([]interface{}, len(records))
	for index, record := range records {
		record.Id = model.NewId()
		record.BatchKey = key
		record.CreatedTime = &time
		recordInterfaces[index] = record
	}
	session := mongo.Get()
	defer session.Close()
	session.MustInsert(collectionExcel, recordInterfaces...)
	return records
}

func Search(key string, query string) []*model.Excel {
	session := mongo.Get()
	defer session.Close()
	records := []*model.Excel{}
	session.MustFind(collectionExcel, bson.M{"$or": []bson.M{bson.M{"tel": query}, bson.M{"name": query} }, "batchKey": key}, &records)
	return records
}