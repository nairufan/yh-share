package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionTopDocuments = "top_document"
)

func AddTopDocuments(documents []*model.TopDocument) []*model.TopDocument {

	time := time.Now()
	documentInterfaces := make([]interface{}, len(documents))
	for index, document := range documents {
		document.Id = model.NewId()
		document.CreatedTime = &time
		documentInterfaces[index] = document
	}
	session := mongo.Get()
	defer session.Close()
	session.MustInsert(collectionTopDocuments, documentInterfaces...)
	return documents
}

func TopDocumentStatistics(start time.Time, end time.Time) []*model.Statistic {
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
	session.MustPipeAll(collectionTopDocuments, []bson.M{match, group}, &results)
	return results
}

func TopDocumentCount() int {
	session := mongo.Get()
	defer session.Close()
	return session.MustCount(collectionTopDocuments)
}