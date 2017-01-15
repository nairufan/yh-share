package controllers

import (
	"github.com/nairufan/yh-share/service"
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