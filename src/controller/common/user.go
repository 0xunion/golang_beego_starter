package common

import (
	"github.com/0xunion/exercise_back/src/controller"
	"github.com/0xunion/exercise_back/src/service/common"
	"github.com/0xunion/exercise_back/src/types"
	master_types "github.com/0xunion/exercise_back/src/types"
	beego "github.com/beego/beego/v2/server/web"
)

// UserLoginByEmailAndPasswordController
type UserLoginByEmailAndPasswordController struct {
	beego.Controller
}

func (c *UserLoginByEmailAndPasswordController) Post() {
	var request_params struct {
		Email        string `json:"email" valid:"Required;Email" form:"email"`
		Password     string `json:"password" valid:"Required" form:"password"`
		Captcha      string `json:"captcha" valid:"Required" form:"captcha"`
		CaptchaToken string `json:"captcha_token" valid:"Required" form:"captcha_token"`
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
		Phone        string `json:"phone" valid:"Required;Mobile" form:"phone"`
		Password     string `json:"password" valid:"Required" form:"password"`
		Captcha      string `json:"captcha" valid:"Required" form:"captcha"`
		CaptchaToken string `json:"captcha_token" valid:"Required" form:"captcha_token"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.UserLoginByPhoneAndPasswordService(request_params.Phone, request_params.Password, request_params.CaptchaToken, request_params.Captcha), true, false)
}

// InitRootUserController
type InitRootUserController struct {
	beego.Controller
}

func (c *InitRootUserController) Post() {
	var request_params struct {
		Email    string `json:"email" valid:"Required;Email" form:"email"`
		Password string `json:"password" valid:"Required" form:"password"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.InitRootUserService(request_params.Email, request_params.Password), true, false)
}

type InfoSelfController struct {
	beego.Controller
}

func (c *InfoSelfController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	c.Ctx.Output.JSON(master_types.SuccessResponse(user), true, false)
}
