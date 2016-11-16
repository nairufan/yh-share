package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:ExcelController"],
		beego.ControllerComments{
			"Upload",
			`/upload`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:ExcelController"],
		beego.ControllerComments{
			"Save",
			`/save`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:ExcelController"],
		beego.ControllerComments{
			"Search",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/nairufan/yh-share/controllers:UserController"],
		beego.ControllerComments{
			"WxLogin",
			`/wx-login`,
			[]string{"get"},
			nil})

}
