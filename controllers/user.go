package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
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
	beego.Info("redirectUrl: ", redirectUrl)
	u.Redirect(redirectUrl, 301)
}

type authResponse struct {
	AccessToken  string       `json:"access_token"`
	ExpiresIn    int64          `json:"expires_in"`
	RefreshToken string       `json:"refresh_token"`
	Openid       string       `json:"openid"`
	Scope        string       `json:"scope"`
}

// @router /wx-login-openid [get]
func (u *UserController) WxLoginResolve() {
	code := u.GetString("code")
	appId := beego.AppConfig.String("wechat.appId")
	secret := beego.AppConfig.String("wechat.appSecret")
	authUrl := beego.AppConfig.String("wechat.openIdUrl")
	authUrl = fmt.Sprintf(authUrl, appId, secret, code)
	resp, err := http.Get(authUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	beego.Info(string(body))
	response := &authResponse{}
	json.Unmarshal(body, response)
	beego.Info("openId:", response.Openid)
	u.SetUserId(response.Openid)
	//u.SetToken(response.AccessToken)
	//u.SetExpire(response.ExpiresIn)
	beego.Info(u.GetToken(), u.GetExpire())
	u.Redirect("/index", 301)
}

// @router /top-wx-login [get]
func (u *UserController) TopWxLogin() {
	appId := beego.AppConfig.String("wechat.appId")
	authUrl := beego.AppConfig.String("wechat.authUrl")
	hostname := beego.AppConfig.String("hostname")
	redirectUrl := fmt.Sprintf(authUrl, appId, hostname + "/api/user/top-wx-login-openid")
	beego.Info("redirectUrl: ", redirectUrl)
	u.Redirect(redirectUrl, 301)
}

// @router /top-wx-login-openid [get]
func (u *UserController) TopWxLoginResolve() {
	code := u.GetString("code")
	appId := beego.AppConfig.String("wechat.appId")
	secret := beego.AppConfig.String("wechat.appSecret")
	authUrl := beego.AppConfig.String("wechat.openIdUrl")
	authUrl = fmt.Sprintf(authUrl, appId, secret, code)
	resp, err := http.Get(authUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	beego.Info(string(body))
	response := &authResponse{}
	json.Unmarshal(body, response)
	beego.Info("openId:", response.Openid)
	u.SetUserId(response.Openid)
	//u.SetToken(response.AccessToken)
	//u.SetExpire(response.ExpiresIn)
	beego.Info(u.GetToken(), u.GetExpire())
	u.Redirect("/top", 301)
}

// @router /mock-login [get]
func (u *UserController) MockLogin() {
	u.SetUserId("1111")
}

type loginRequest struct {
	Tel      string       `json:"tel" validate:"required"`
	Password string       `json:"password" validate:"required"`
}

type loginResponse struct {
	Success bool       `json:"success"`
	Msg     string     `json:"msg"`
}

// @router /admin-login [post]
func (u *UserController) AdminLogin() {
	var request loginRequest
	response := &loginResponse{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &request); err != nil {
		panic(err)
	}
	errs := validate.Struct(request)
	if errs != nil {
		panic(errs)
	}
	if validateAccount(request.Tel, request.Password) {
		u.SetUserId(request.Tel)
		u.SetSession("role", "admin")
		response.Success = true
	} else {
		response.Success = false
		response.Msg = "tel or password is error."
	}
	u.Data["json"] = response
	u.ServeJSON()
}

func validateAccount(tel string, password string) bool {
	if "12345678910" != tel {
		return false
	}
	if password != "jl@2016" {
		return false
	}
	return true
}