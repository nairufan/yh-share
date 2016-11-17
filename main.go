package main

import (
	_ "github.com/nairufan/yh-share/docs"
	_ "github.com/nairufan/yh-share/routers"
	"github.com/astaxie/beego"
	"github.com/nairufan/yh-share/controllers"
	"github.com/nairufan/yh-share/filters"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.InsertFilter("/api/excel/*", beego.BeforeRouter, filters.LoginCheck)
	beego.Run()
}
