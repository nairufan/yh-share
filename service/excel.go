package service

import (
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/db/mongo"
	"time"
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