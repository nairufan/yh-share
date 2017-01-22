package controllers

import (
	"github.com/nairufan/yh-share/service"
	"github.com/nairufan/yh-share/model"
	"time"
)

type TopRecordController struct {
	BaseController
}

// @router /list [get]
func (u *TopRecordController) TopList() {
	list := service.TopRecordsList(u.GetUserId(), 0, 10)
	u.Data["json"] = list
	u.ServeJSON();
}

// @router /document-statistics [get]
func (u *TopRecordController) TopDocumentStatistics() {
	response := model.StatisticResponse{}
	now := time.Now()
	start := now.AddDate(0, 0, -10)
	statistics := service.TopDocumentStatistics(start, now)
	total := service.TopDocumentCount()
	response.Statistics = statistics
	response.Total = total
	u.Data["json"] = response
	u.ServeJSON()
}

// @router /statistics [get]
func (u *TopRecordController) TopStatistics() {
	response := model.StatisticResponse{}
	now := time.Now()
	start := now.AddDate(0, 0, -10)
	statistics := service.TopRecordsStatistics(start, now)
	total := service.TopRecordsCount()
	response.Statistics = statistics
	response.Total = total
	u.Data["json"] = response
	u.ServeJSON()
}