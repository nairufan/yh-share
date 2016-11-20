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
	TelCol     int       `json:"telCol,required"`
	NameCol    int       `json:"nameCol,required"`
	Title      string    `json:"title,required"`
	DocumentId string    `json:"documentId"` //for attach
}
// @router /save [post]
func (u *ExcelController) Save() {
	var request saveRequest
	json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	records := u.ClearExcel()
	if records == nil {
		util.Panic(errors.New("No excel file found."))
	}
	excelRecords := []*model.Excel{}
	for _, record := range records {
		if len(record) > 2 {
			excelRecords = append(excelRecords, &model.Excel{
				Tel: record[request.TelCol],
				Name: record[request.NameCol],
				Data: record,
			})
		}
	}
	if request.DocumentId == "" {
		document := service.AddDocument(&model.Document{
			UserId: u.GetUserId(),
			Title: request.Title,
		})
		request.DocumentId = document.Id
	}
	excelRecords = service.AddRecords(excelRecords, request.DocumentId)
	u.Data["json"] = map[string]string{
		URL: "/query/" + excelRecords[0].DocumentId,
	}
	u.ServeJSON()
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