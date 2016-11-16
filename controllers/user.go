package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UserController struct {
	beego.Controller
}

// @router /wx-login [get]
func (u *UserController) WxLogin() {
	openId := u.GetString("code")
	fmt.Println("openId:", openId)
}