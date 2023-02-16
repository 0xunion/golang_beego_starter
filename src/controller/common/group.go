package common

import (
	"github.com/0xunion/exercise_back/src/controller"
	"github.com/0xunion/exercise_back/src/service/common"
	"github.com/0xunion/exercise_back/src/types"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateGroupController
type CreateGroupController struct {
	beego.Controller
}

func (c *CreateGroupController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	var request_params struct {
		Name        string `json:"name" valid:"Required" form:"name"`
		Description string `json:"description" valid:"Required" form:"description"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.CreateGroupService(user, request_params.Name, request_params.Description), true, false)
}

// ListMyGroupsController
type ListMyGroupsController struct {
	beego.Controller
}

func (c *ListMyGroupsController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	var request_params struct {
		Index int64 `json:"index" form:"index"`
		Limit int64 `json:"limit" form:"limit"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.ListMyGroupsService(user, request_params.Index, request_params.Limit), true, false)
}

// ListGroupsController
type ListGroupsController struct {
	beego.Controller
}

func (c *ListGroupsController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	var request_params struct {
		Index int64 `json:"index" form:"index"`
		Limit int64 `json:"limit" form:"limit"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.ListGroupsService(user, request_params.Index, request_params.Limit), true, false)
}

// InfoGroupController
type InfoGroupController struct {
	beego.Controller
}

func (c *InfoGroupController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.InfoGroupService(user, gid), true, false)
}

// UpdateGroupController
type UpdateGroupController struct {
	beego.Controller
}

func (c *UpdateGroupController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	var request_params struct {
		Field string      `json:"field" valid:"Required" form:"field"`
		Value interface{} `json:"value" valid:"Required" form:"value"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.UpdateGroupService(user, gid, request_params.Field, request_params.Value), true, false)
}

// DeleteGroupController
type DeleteGroupController struct {
	beego.Controller
}

func (c *DeleteGroupController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.DeleteGroupService(user, gid), true, false)
}

// ListGroupMembersController
type ListGroupMembersController struct {
	beego.Controller
}

func (c *ListGroupMembersController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	var request_params struct {
		Index int64 `json:"index" form:"index"`
		Limit int64 `json:"limit" form:"limit"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.ListGroupMembersService(user, gid, request_params.Index, request_params.Limit), true, false)
}

// CreateUserInGroupByEmailAndPasswordController
type CreateUserInGroupByEmailAndPasswordController struct {
	beego.Controller
}

func (c *CreateUserInGroupByEmailAndPasswordController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	var request_params struct {
		Email    string `json:"email" valid:"Required" form:"email"`
		Password string `json:"password" valid:"Required" form:"password"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.CreateUserInGroupByEmailAndPasswordService(user, gid, request_params.Email, request_params.Password), true, false)
}

// CreateUserInGroupByPhoneAndPasswordController
type CreateUserInGroupByPhoneAndPasswordController struct {
	beego.Controller
}

func (c *CreateUserInGroupByPhoneAndPasswordController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	var request_params struct {
		Phone    string `json:"phone" valid:"Required" form:"phone"`
		Password string `json:"password" valid:"Required" form:"password"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.CreateUserInGroupByPhoneAndPasswordService(user, gid, request_params.Phone, request_params.Password), true, false)
}

// UpdateGroupMemberRoleController
type UpdateGroupMemberRoleController struct {
	beego.Controller
}

func (c *UpdateGroupMemberRoleController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	uid, err := primitive.ObjectIDFromHex(c.GetString("uid"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	var request_params struct {
		Permission string `json:"permission" valid:"Required" form:"permission"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.UpdateGroupMemberRoleService(user, gid, uid, request_params.Permission), true, false)
}

// DeleteGroupMemberController
type DeleteGroupMemberController struct {
	beego.Controller
}

func (c *DeleteGroupMemberController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	uid, err := primitive.ObjectIDFromHex(c.GetString("uid"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.DeleteGroupMemberService(user, gid, uid), true, false)
}

// CreateUserInGroupByExcelController
type CreateUserInGroupByExcelController struct {
	beego.Controller
}

func (c *CreateUserInGroupByExcelController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*types.User)

	gid, err := primitive.ObjectIDFromHex(c.GetString("id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	fid, err := primitive.ObjectIDFromHex(c.GetString("file_id"))
	if err != nil {
		c.Ctx.Output.JSON(types.ErrorResponse(-400, err.Error()), true, false)
	}

	c.Ctx.Output.JSON(common.CreateUserInGroupByExcelService(user, gid, fid), true, false)
}
