package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"qiniupkg.com/x/jsonutil.v7"
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
	ExpiresIn    int          `json:"expires_in"`
	RefreshToken string       `json:"refresh_token"`
	Openid       string       `json:"openid,required"`
	Scope        string       `json:"openid"`
}

// @router /wx-login-openid [get]
func (u *UserController) WxLoginResolve() {
	code := u.GetString("code")
	appId := beego.AppConfig.String("wechat.appId")
	secret := beego.AppConfig.String("wechat.appSecret")
	authUrl := beego.AppConfig.String("wechat.accessTokenUrl")
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
	beego.Info(body)
	response := &authResponse{}
	jsonutil.Unmarshal(string(body), response)
	beego.Info("openId:", response.Openid)
	u.SetUserId(response.Openid)
	u.Redirect("/index", 301)
}

// @router /mock-login [get]
func (u *UserController) MockLogin() {
	u.SetUserId("1111")
}