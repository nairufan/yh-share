package controllers

import (
	"io/ioutil"
	"github.com/astaxie/beego/context"
)

func Index(ctx *context.Context) {
	ctx.Redirect(302, "https://www.u365.me/wx/html/qrcode")
}

func WeiXin(ctx *context.Context) {
	content, err := ioutil.ReadFile("static/admin/weixin.html")
	if err != nil {
		panic(err)
	}
	ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Output.Body(content)
}

func Statistic(ctx *context.Context) {
	content, err := ioutil.ReadFile("static/admin/statistic.html")
	if err != nil {
		panic(err)
	}
	ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Output.Body(content)
}

func Login(ctx *context.Context) {
	content, err := ioutil.ReadFile("static/admin/login.html")
	if err != nil {
		panic(err)
	}
	ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Output.Body(content)
}