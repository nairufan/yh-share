package controllers

import (
	"github.com/astaxie/beego"
)

const (
	UserID = "userId"
	Excel = "excel"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) SetUserId(id string) {
	b.SetSession(UserID, id)
}

func (b *BaseController) GetUserId() string {
	uid := b.GetSession(UserID)
	if uid == nil {
		return ""
	}
	return uid.(string)
}

func (b *BaseController) SetExcel(data [][]string) {
	b.SetSession(Excel, data)
}

func (b *BaseController) getExcel() [][]string {
	excel := b.GetSession(Excel)
	if excel == nil {
		return nil
	}
	return excel.([][]string)
}

func (b *BaseController) ClearExcel() [][]string {
	excel := b.getExcel()
	b.SetSession(Excel, nil)
	return excel
}