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

func DocumentList(userId string) []*model.Document {
	session := mongo.Get()
	defer session.Close()
	documents := []*model.Document{}
	session.MustFind(collectionDocuments, bson.M{"userId": userId}, &documents)
	return documents
}