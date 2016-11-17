package filters

import (
	"github.com/astaxie/beego/context"
)

const (
	UserID = "userId"
)
func LoginCheck(ctx *context.Context) {
	_, ok := ctx.Input.Session(UserID).(string)
	if !ok {
		panic("403")
	}
}