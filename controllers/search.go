package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/nairufan/yh-share/service"
	"io/ioutil"
	"strings"
)

func Search(ctx *context.Context) {
	content, err := ioutil.ReadFile("static/uhsearch.html")
	if err != nil {
		panic(err)
	}
	id := ctx.Input.Param(":id")
	document := service.GetDocumentById(id)
	contentString := string(content)
	contentString = strings.Replace(contentString, "{{$title}}", document.Title, 1)
	ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Output.Body([]byte(contentString))
}
