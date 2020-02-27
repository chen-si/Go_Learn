package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type TestContextController struct {
	beego.Controller
}

func (c *TestContextController) Get() {
	c.Ctx.WriteString(c.Ctx.Input.IP() + ":" + strconv.Itoa(c.Ctx.Input.Port()))
	c.Ctx.WriteString(c.Ctx.Input.Query("name"))
}
