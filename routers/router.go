// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/nairufan/yh-share/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"io/ioutil"
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
	)
	beego.Get("/search/:id", func(ctx *context.Context){
		content, err := ioutil.ReadFile("static/uhsearch.html")
		if err != nil {
			panic(err)
		}
		ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		ctx.Output.Body(content)
	})
	beego.AddNamespace(ns)
}
