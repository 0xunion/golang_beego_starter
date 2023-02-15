package test

import (
	beego "github.com/beego/beego/v2/server/web"
)


type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Ctx.WriteString("Hello World!")
}
