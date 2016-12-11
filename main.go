package main

import (
	_ "github.com/nairufan/yh-share/docs"
	_ "github.com/nairufan/yh-share/routers"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
	"github.com/nairufan/yh-share/filters"
	"qiniupkg.com/x/rpc.v7/gob"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.WebConfig.StaticDir["/html"] = "static"
	beego.InsertFilter("/api/document/upload", beego.BeforeRouter, filters.LoginCheck)
	beego.InsertFilter("/api/document/save", beego.BeforeRouter, filters.LoginCheck)
	beego.InsertFilter("/api/document/list", beego.BeforeRouter, filters.LoginCheck)
	beego.InsertFilter("/api/document/changeTitle", beego.BeforeRouter, filters.LoginCheck)
	beego.InsertFilter("/index", beego.BeforeRouter, filters.LoginCheck)
	beego.SetStaticPath("/index", "static/uhdingdan.html")
	beego.SetStaticPath("/uhdingdan.js", "static/uhdingdan.js")
	beego.SetStaticPath("/search/uhsearch.js", "static/uhsearch.js")
	beego.SetStaticPath("/s/uhsearch.js", "static/uhsearch.js")
	beego.SetStaticPath("/MP_verify_h0emLIGJ0SsISIPS.txt", "static/MP_verify_h0emLIGJ0SsISIPS.txt")
	beego.Run()
}

func init() {
	gob.Register([][]string{})
}