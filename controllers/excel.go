package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nairufan/yh-share/util"
	"encoding/json"
	"github.com/nairufan/yh-share/model"
	"github.com/nairufan/yh-share/service"
	"fmt"
)

type ExcelController struct {
	beego.Controller
}

const (
	Excel = "excel"
	URL = "url"
)

// @router /upload [post]
func (u *ExcelController) Upload() {
	f, _, err := u.GetFile("records")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		panic(err)
	}
	records := util.ParseFile(f)
	u.SetSession(Excel, records)
	u.Data["json"] = records
	u.ServeJSON();
}

type saveRequest struct {
	TelCol  int       `json:"telCol,required"`
	NameCol int       `json:"nameCol,required"`
}
// @router /save [post]
func (u *ExcelController) Save() {
	var request saveRequest
	json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	fmt.Println(u.GetSession(Excel))
	records := u.GetSession(Excel).([][]string)
	excelRecords := []*model.Excel{}
	for _, record := range records {
		if len(record) > 2 {
			excelRecords = append(excelRecords, &model.Excel{
				UserId: "1212",
				Tel: record[request.TelCol],
				Name: record[request.NameCol],
				Data: records,
			})
		}
	}
	excelRecords = service.AddRecords(excelRecords)
	u.Data["json"] = map[string]string{
		URL: "/query/" + excelRecords[0].BatchKey,
	}
	u.ServeJSON()
}
// @router /search [get]
func (u *ExcelController) Search() {
	query := u.GetString("query")
	key := u.GetString("key")
	records := service.Search(key, query)
	u.Data["json"] = records
	u.ServeJSON()
}
