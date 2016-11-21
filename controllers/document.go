package controllers

import (
	"github.com/nairufan/yh-share/util"
	"encoding/json"
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/service"
	"github.com/astaxie/beego"
	"errors"
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
		URL: "/query/" + recordModels[0].DocumentId,
	}
	u.ServeJSON()
}

type searchResponse struct {
	
}
// @router /search [get]
func (u *ExcelController) Search() {
	query := u.GetString("query")
	key := u.GetString("documentId")
	beego.Info(query)
	records := service.Search(key, query)
	u.Data["json"] = records
	u.ServeJSON()
}
// @router /list [get]
func (u *ExcelController) List() {
	list := service.DocumentList(u.GetUserId(), 0, 10)
	u.Data["json"] = list
	u.ServeJSON()
}