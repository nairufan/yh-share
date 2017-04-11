package routers

import (
	"github.com/nairufan/yh-share/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/document",
			beego.NSInclude(
				&controllers.ExcelController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/top",
			beego.NSInclude(
				&controllers.TopRecordController{},
			),
		),
	)
	beego.Get("/search/:id", controllers.Search)
	beego.Get("/s/:id", controllers.Search)
	beego.Get("/u/:id", controllers.SearchByUser)
	beego.Get("/statistic.html", controllers.Statistic)
	beego.Get("", controllers.Index)
	beego.Get("/weixin.html", controllers.WeiXin)
	beego.Get("/login.html", controllers.Login)
	beego.AddNamespace(ns)
}
