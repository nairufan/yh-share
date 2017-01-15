package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
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