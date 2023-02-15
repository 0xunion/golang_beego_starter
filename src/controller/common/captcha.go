package common

import (
	"${package}/src/service/common"
	beego "github.com/beego/beego/v2/server/web"
)

// ImageMathmaticalCaptchaController
type ImageMathmaticalCaptchaController struct {
	beego.Controller
}

// Get
func (c *ImageMathmaticalCaptchaController) Get() {
	c.Ctx.Output.JSON(common.LoginCaptchaService(), true, false)
}
