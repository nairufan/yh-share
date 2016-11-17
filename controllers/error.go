package controllers

import (
	"github.com/astaxie/beego"
)

type Error struct {
	Error     string        `json:"Error"`
	ErrorCode int           `json:"Error Code"`
	Reason    interface{}  `json:"reason"`
}
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = &Error{
		Error: "Not found",
		ErrorCode: 404,
		Reason: c.Ctx.Input.Data(),
	}
	c.ServeJSON()
}

func (c *ErrorController) Error403() {
	c.Data["json"] = &Error{
		Error: "Forbidden",
		ErrorCode: 403,
	}
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = &Error{
		Error: "Internal server error.",
		ErrorCode: 500,
		Reason: c.Ctx.Input.Data(),
	}
	c.ServeJSON()
}