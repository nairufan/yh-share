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

func SearchAll(documentIds []string, query string) []*model.Record {
	session := mongo.Get()
	defer session.Close()
	records := []*model.Record{}
	session.MustFind(collectionRecords, bson.M{"$or": []bson.M{bson.M{"queryField1": query}, bson.M{"queryField2": query} }, "documentId": bson.M{"$in": documentIds}}, &records)
	return records
}

func RecordsStatistics(start time.Time, end time.Time) []*model.Statistic {
	results := []*model.Statistic{}
	session := mongo.Get()
	defer session.Close()
	group := bson.M{}
	match := bson.M{}
	//date := bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$createdTime"}}
	date := bson.M{"$concat": []interface{}{bson.M{"$year": "$createdTime"}, "-", bson.M{"month": "$createdTime"}, "-", bson.M{"dayOfMonth": "$createdTime"}}}
	group["$group"] = bson.M{"_id": date, "count": bson.M{"$sum": 1}}
	match["$match"] = bson.M{"createdTime": bson.M{"$gte": start, "$lte": end}}
	session.MustPipeAll(collectionRecords, []bson.M{match, group}, &results)
	return results
}

func RecordsCount() int {
	session := mongo.Get()
	defer session.Close()
	return session.MustCount(collectionRecords)
}