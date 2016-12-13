package filters

import (
	"github.com/astaxie/beego/context"
)

const (
	UserID = "userId"
	Role = "role"
)

func LoginCheck(ctx *context.Context) {
	_, ok := ctx.Input.Session(UserID).(string)
	if !ok {
		panic("403")
	}
}

func AdminCheck(ctx *context.Context) {
	val, ok := ctx.Input.Session(Role).(string)
	if !ok || val != "admin" {
		ctx.Redirect(302, "/login.html")
	}
}