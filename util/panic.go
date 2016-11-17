package util

import "github.com/astaxie/beego"

func Panic(err error) {
	beego.Error(err)
	panic("500")
}
