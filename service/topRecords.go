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

func TopRecordsStatistics(start time.Time, end time.Time) []*model.Statistic {
	results := []*model.Statistic{}
	session := mongo.Get()
	defer session.Close()
	group := bson.M{}
	match := bson.M{}
	//date := bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$createdTime"}}
	year := bson.M{"$substr": []interface{}{"$createdTime", 0, 4}}
	month := bson.M{"$substr": []interface{}{"$createdTime", 5, 2}}
	day := bson.M{"$substr": []interface{}{"$createdTime", 8, 2}}
	date := bson.M{"$concat": []interface{}{year, "-", month, "-", day}}
	group["$group"] = bson.M{"_id": date, "count": bson.M{"$sum": 1}}
	match["$match"] = bson.M{"createdTime": bson.M{"$gte": start, "$lte": end}}
	session.MustPipeAll(collectionTopRecords, []bson.M{match, group}, &results)
	return results
}

func TopRecordsCount() int {
	session := mongo.Get()
	defer session.Close()
	return session.MustCount(collectionTopRecords)
}