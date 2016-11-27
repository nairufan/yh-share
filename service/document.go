package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionDocuments = "document"
)

func AddDocument(document *model.Document) *model.Document {
	time := time.Now()
	document.Id = model.NewId()
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