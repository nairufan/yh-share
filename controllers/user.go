package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UserController struct {
	BaseController
}

// @router /wx-login [get]
func (u *UserController) WxLogin() {
	appId := beego.AppConfig.String("wechat.appId")
	authUrl := beego.AppConfig.String("wechat.authUrl")
	hostname := beego.AppConfig.String("hostname")
	redirectUrl := fmt.Sprintf(authUrl, appId, hostname + "/api/user/wx-login-openid")
	beego.Info("redirectUrl: ")
	u.Redirect(redirectUrl, 301)
}

// @router /wx-login-openid [get]
func (u *UserController) WxLoginResolve() {
	openId := u.GetString("code")
	beego.Info("openId:", openId)
	//Todo mock login
	//u.SetUserId("1111")
}

// @router /mock-login [get]
func (u *UserController) MockLogin() {
	u.SetUserId("1111")
}