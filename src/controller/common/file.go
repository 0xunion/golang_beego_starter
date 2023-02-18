package common

import (
	master_servce "github.com/0xunion/exercise_back/src/service/common"
	master_types "github.com/0xunion/exercise_back/src/types"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateFileController struct {
	beego.Controller
}

func (c *CreateFileController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	file, header, err := c.GetFile("file")
	if err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	response := master_servce.CreateFileService(
		user,
		file,
		header,
	)
	c.Ctx.Output.JSON(response, true, false)
}

type GetFileController struct {
	beego.Controller
}

func (c *GetFileController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	file_id, err := primitive.ObjectIDFromHex(c.GetString("file_id"))
	if err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	path, err := master_servce.GetFileService(
		user,
		file_id,
	)

	if err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	c.Ctx.Output.Download(path)
}