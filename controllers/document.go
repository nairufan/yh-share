package controllers

import (
	"encoding/json"
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/service"
	"github.com/astaxie/beego"
	"errors"
	"time"
	"github.com/nairufan/yh-share/util"
	"strings"
	"os"
)

type ExcelController struct {
	BaseController
}

type Size interface {
	Size() int64
}

// @router /upload [post]
func (u *ExcelController) Upload() {
	f, _, err := u.GetFile("attachment")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		panic(err)
	}
	size := f.(Size).Size()
	bytes := make([]byte, size)
	_, err = f.Read(bytes)
	if err != nil {
		panic(err)
	}
	records, err := util.ParseXlsxFile(bytes)
	if err != nil {
		beego.Error(err)
		filePath := util.SaveFile(bytes)
		records = util.ParseXlsFile(filePath)
		err := os.Remove(filePath)
		beego.Error(err)
	}
	u.SetExcel(records)
	u.Data["json"] = records
	u.ServeJSON();
}

type saveRequest struct {
	SearchColumn  []int     `json:"searchColumn" validate:"required"`
	DisplayColumn []int     `json:"displayColumn" validate:"required"`
	ExpressColumn int       `json:"expressColumn"`
	TitleRow      int       `json:"titleRow"`
	Title         string    `json:"title" validate:"required"`
	DocumentId    string    `json:"documentId"` //for attach
}

type saveResponse struct {
	Id     string    `json:"id"`
	UserId string    `json:"userId"`
}

// @router /save [post]
func (u *ExcelController) Save() {
	beego.Info(string(u.Ctx.Input.RequestBody))
	var request saveRequest
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &request); err != nil {
		panic(err)
	}
	errs := validate.Struct(request)
	if errs != nil {
		panic(errs)
	}
	records := u.ClearExcel()
	if records == nil {
		beego.Error("No excel file found.")
		panic(errors.New("No excel file found."))
	}
	recordModels := []*model.Record{}
	for index, record := range records {
		if len(record) > 2 && index != request.TitleRow {
			recordModel := &model.Record{
				Data: record,
			}
			if request.SearchColumn != nil && len(request.SearchColumn) > 0 {
				recordModel.QueryField1 = strings.TrimSpace(record[request.SearchColumn[0]])
				if len(request.SearchColumn) > 1 {
					recordModel.QueryField2 = strings.TrimSpace(record[request.SearchColumn[1]])
				}
			}
			recordModels = append(recordModels, recordModel)
		}

	}
	if request.DocumentId == "" {
		titleRow := records[request.TitleRow]
		newDisplayColumn := []int{}
		newDisplayColumn = append(newDisplayColumn, request.ExpressColumn)
		for _, col := range request.DisplayColumn {
			if col != request.ExpressColumn {
				newDisplayColumn = append(newDisplayColumn, col)
			}
		}
		document := service.AddDocument(&model.Document{
			UserId: u.GetUserId(),
			Title: request.Title,
			TitleFields: titleRow,
			DisplayColumn: newDisplayColumn,
			ExpressColumn: request.ExpressColumn,
		})
		request.DocumentId = document.Id
	}
	recordModels = service.AddRecords(recordModels, request.DocumentId)
	u.Data["json"] = &saveResponse{
		Id: request.DocumentId,
		UserId: u.GetUserId(),
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
// @router /searchByUser [get]
func (u *ExcelController) SearchByUser() {
	response := []*recordType{}
	query := u.GetString("query")
	userId := u.GetString("userId")
	documents := service.GetDocumentsByEndDay(userId, 30)
	documentIds := []string{}
	for _, document := range documents {
		documentIds = append(documentIds, document.Id)
	}
	records := service.SearchAll(documentIds, query)
	documentMap := convertDocumentMap(documents)
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

const timeFormat = "2006-01-02"

type listTypeResponse struct {
	TimeList []string `json:"timeList"`
	DataMap  map[string][]*model.Document     `json:"dataMap"`
}

// @router /list [get]
func (u *ExcelController) List() {
	response := &listTypeResponse{}
	list := service.DocumentList(u.GetUserId(), 0, 100)
	timeList := []string{}
	dataMap := map[string][]*model.Document{}
	for _, document := range list {
		formatTime := document.CreatedTime.Format(timeFormat)
		tmpDocumentList := dataMap[formatTime]
		if tmpDocumentList == nil {
			tmpDocumentList = []*model.Document{}
			timeList = append(timeList, formatTime)
		}
		tmpDocumentList = append(tmpDocumentList, document)
		dataMap[formatTime] = tmpDocumentList
	}
	response.DataMap = dataMap
	response.TimeList = timeList
	u.Data["json"] = response
	u.ServeJSON()
}

type titleRequest struct {
	Title string     `json:"title" validate:"required"`
	Id    string     `json:"id" validate:"required"`
}

// @router /changeTitle [post]
func (u *ExcelController) ChangeTitle() {
	var request titleRequest
	json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	errs := validate.Struct(request)
	if errs != nil {
		panic(errs)
	}

	document := service.GetDocumentById(request.Id)
	document.Title = request.Title
	service.UpdateDocument(document)
	u.Data["json"] = document
	u.ServeJSON()
}

// @router /statistics [get]
func (u *ExcelController) Statistics() {
	response := model.StatisticResponse{}
	now := time.Now()
	start := now.AddDate(0, 0, -10)
	statistics := service.DocumentStatistics(start, now)
	total := service.DocumentCount()
	response.Statistics = statistics
	response.Total = total
	u.Data["json"] = response
	u.ServeJSON()
}

// @router /order-statistics [get]
func (u *ExcelController) OrderStatistics() {
	response := model.StatisticResponse{}
	now := time.Now()
	start := now.AddDate(0, 0, -10)
	statistics := service.RecordsStatistics(start, now)
	total := service.RecordsCount()
	response.Statistics = statistics
	response.Total = total
	u.Data["json"] = response
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

func convertDocumentMap(documents []*model.Document) map[string]*model.Document {
	documentMap := map[string]*model.Document{}
	for _, document := range documents {
		documentMap[document.Id] = document
	}
	return documentMap
}