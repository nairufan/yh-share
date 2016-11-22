package controllers

import (
	"github.com/nairufan/yh-share/util"
	"encoding/json"
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/service"
	"github.com/astaxie/beego"
	"errors"
	"time"
)

type ExcelController struct {
	BaseController
}

const (
	URL = "url"
)

// @router /upload [post]
func (u *ExcelController) Upload() {
	f, _, err := u.GetFile("attachment")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		util.Panic(err)
	}
	records := util.ParseFile(f)
	u.SetExcel(records)
	u.Data["json"] = records
	u.ServeJSON();
}

type saveRequest struct {
	SearchColumn  []int     `json:"searchColumn,required"`
	DisplayColumn []int    `json:"displayColumn,required"`
	TitleRow      int       `json:"titleRow,required"`
	Title         string    `json:"title,required"`
	DocumentId    string    `json:"documentId"` //for attach
}
// @router /save [post]
func (u *ExcelController) Save() {
	var request saveRequest
	json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	records := u.ClearExcel()
	if records == nil {
		util.Panic(errors.New("No excel file found."))
	}
	recordModels := []*model.Record{}
	beego.Info(records)
	for _, record := range records {
		if len(record) > 2 {
			recordModel := &model.Record{
				Data: record,
			}
			if request.SearchColumn != nil && len(request.SearchColumn) > 0 {
				recordModel.QueryField1 = record[request.SearchColumn[0]]
				if len(request.SearchColumn) > 1 {
					recordModel.QueryField2 = record[request.SearchColumn[1]]
				}
			}
			recordModels = append(recordModels, recordModel)
		}

	}
	if request.DocumentId == "" {
		titleRow := records[request.TitleRow]
		document := service.AddDocument(&model.Document{
			UserId: u.GetUserId(),
			Title: request.Title,
			TitleFields: titleRow,
			DisplayColumn: request.DisplayColumn,
		})
		request.DocumentId = document.Id
	}
	recordModels = service.AddRecords(recordModels, request.DocumentId)
	u.Data["json"] = map[string]string{
		URL: "/api/document/list",
	}
	u.ServeJSON()
}

type recordType struct {
	Title      []string     `json:"title"`
	Data       []string     `json:"data"`
	CreateTime *time.Time   `bson:"createdTime"`
}
// @router /search [get]
func (u *ExcelController) Search() {
	response := []*recordType{}
	query := u.GetString("query")
	key := u.GetString("documentId")
	beego.Info(query)
	records := service.Search(key, query)
	documentIds := getDistinctIds(records)
	documentMap := getDocumentMap(documentIds)
	for _, record := range records {
		resultRecord := []string{}
		resultTitle := []string{}
		document := documentMap[record.DocumentId]
		data := record.Data
		title := document.TitleFields
		for _, col := range document.DisplayColumn {
			resultRecord = append(resultRecord, data[col])
			resultTitle = append(resultTitle, title[col])
		}
		response = append(response, &recordType{
			Title: resultTitle,
			Data: resultRecord,
			CreateTime: document.CreatedTime,
		})
	}
	u.Data["json"] = response
	u.ServeJSON()
}
// @router /list [get]
func (u *ExcelController) List() {
	list := service.DocumentList(u.GetUserId(), 0, 10)
	u.Data["json"] = list
	u.ServeJSON()
}

func getDistinctIds(records []*model.Record) []string {
	documentIds := []string{}
	documentIdMap := map[string]bool{}
	for _, record := range records {
		if !documentIdMap[record.DocumentId] {
			documentIds = append(documentIds, record.DocumentId)
		}
		documentIdMap[record.DocumentId] = true
	}
	return documentIds
}

func getDocumentMap(ids []string) map[string]*model.Document {
	documentMap := map[string]*model.Document{}
	documents := service.GetDocumentByIds(ids)
	for _, document := range documents {
		documentMap[document.Id] = document
	}
	return documentMap
}