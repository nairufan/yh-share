package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

// @router /wx-login [get]
func (u *UserController) WxLogin() {
	openId := u.GetString("code")
	beego.Info("openId:", openId)
	//Todo mock login
	u.SetUserId("1111")
}