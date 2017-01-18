package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/nairufan/yh-share/util"
	"strconv"
)

const (
	collectionDocuments = "document"
)

func AddDocument(document *model.Document) *model.Document {
	time := time.Now()
	prefix := util.GetRandomString(2)
	nextId := Increase()
	document.Id = prefix + strconv.FormatInt(nextId, 16)
	document.CreatedTime = &time
	session := mongo.Get()
	defer session.Close()
	session.MustInsert(collectionDocuments, document)
	return document
}

func UpdateDocument(document *model.Document) *model.Document {
	session := mongo.Get()
	defer session.Close()
	session.MustUpdateId(collectionDocuments, document.Id, document)
	return document
}

func GetDocumentById(id string) *model.Document {
	session := mongo.Get()
	defer session.Close()
	document := &model.Document{}
	session.MustFindId(collectionDocuments, id, document)
	return document
}

func GetDocumentByIds(ids []string) []*model.Document {
	session := mongo.Get()
	defer session.Close()
	documents := []*model.Document{}
	query := bson.M{"_id": bson.M{"$in": ids}}
	session.MustFind(collectionDocuments, query, &documents)
	return documents
}

func GetDocumentsByEndDay(userId string, days int) []*model.Document {
	session := mongo.Get()
	defer session.Close()
	now := time.Now()
	past := now.AddDate(0, 0, -1 * days)
	documents := []*model.Document{}
	query := bson.M{"userId": userId, "createdTime": bson.M{"$gte": past}}
	session.MustFind(collectionDocuments, query, &documents)
	return documents
}

func DocumentList(userId string, offset int, limit int) []*model.Document {
	session := mongo.Get()
	defer session.Close()
	documents := []*model.Document{}

	option := mongo.Option{
		Sort: []string{"-createdTime"},
		Limit: &limit,
		Offset: &offset,
	}
	session.MustFindWithOptions(collectionDocuments, bson.M{"userId": userId}, option, &documents)
	return documents
}

func DocumentStatistics(start time.Time, end time.Time) []*model.Statistic {
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
	session.MustPipeAll(collectionDocuments, []bson.M{match, group}, &results)
	return results
}

func DocumentCount() int {
	session := mongo.Get()
	defer session.Close()
	return session.MustCount(collectionDocuments)
}