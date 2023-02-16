package common

import (
	"github.com/0xunion/exercise_back/src/controller"
	"github.com/0xunion/exercise_back/src/service/common"
	"github.com/0xunion/exercise_back/src/types"
	beego "github.com/beego/beego/v2/server/web"
)

// UserLoginByEmailAndPasswordController
type UserLoginByEmailAndPasswordController struct {
	beego.Controller
}

func (c *UserLoginByEmailAndPasswordController) Post() {
	var request_params struct {
		Email        string `json:"email" valid:"Required;Email"`
		Password     string `json:"password" valid:"Required"`
		Captcha      string `json:"captcha" valid:"Required"`
		CaptchaToken string `json:"captcha_token" valid:"Required"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.UserLoginByEmailAndPasswordService(request_params.Email, request_params.Password, request_params.CaptchaToken, request_params.Captcha), true, false)
}

// UserLoginByPhoneAndPasswordController
type UserLoginByPhoneAndPasswordController struct {
	beego.Controller
}

func (c *UserLoginByPhoneAndPasswordController) Post() {
	var request_params struct {
		Phone        string `json:"phone" valid:"Required;Mobile"`
		Password     string `json:"password" valid:"Required"`
		Captcha      string `json:"captcha" valid:"Required"`
		CaptchaToken string `json:"captcha_token" valid:"Required"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.UserLoginByPhoneAndPasswordService(request_params.Phone, request_params.Password, request_params.CaptchaToken, request_params.Captcha), true, false)
}
