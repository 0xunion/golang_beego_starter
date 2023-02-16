package curd

/* @MT-TPL-IMPORT-START */
import (
	controller "github.com/0xunion/exercise_back/src/controller"
	master_service "github.com/0xunion/exercise_back/src/service/curd"
	master_types "github.com/0xunion/exercise_back/src/types"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-CREATE-START */
type GameCreateController struct {
	beego.Controller
}

/* @MT-TPL-CREATE-END */

/* @MT-TPL-UPDATE-START */
type GameUpdateController struct {
	beego.Controller
}

/* @MT-TPL-UPDATE-END */

/* @MT-TPL-DELETE-START */
type GameDeleteController struct {
	beego.Controller
}

/* @MT-TPL-DELETE-END */

/* @MT-TPL-GET-START */
type GameGetController struct {
	beego.Controller
}

/* @MT-TPL-GET-END */

/* @MT-TPL-LIST-START */
type GameListController struct {
	beego.Controller
}

/* @MT-TPL-LIST-END */

/* @MT-TPL-SEARCH-START */
type GameSearchController struct {
	beego.Controller
}

/* @MT-TPL-SEARCH-END */

/* @MT-TPL-CREATE-FUNC-START */
func (c *GameCreateController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	var request_params struct {
		Name string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`

		Description string `json:"description" form:"description" validate:"Required;MinSize(3);MaxSize(20)"`

		HeaderHtml string `json:"header_html" form:"header_html" validate:"Required;MinSize(3);MaxSize(1024)"`

		StartTime int64 `json:"start_time" form:"start_time" validate:""`

		EndTime int64 `json:"end_time" form:"end_time" validate:""`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	// TODO: service
	c.Ctx.Output.JSON(
		master_service.CreateGameService(
			user,
			request_params.Name,
			request_params.Description,
			request_params.HeaderHtml,
			request_params.StartTime,
			request_params.EndTime,
		),
		true, false)
}

/* @MT-TPL-CREATE-FUNC-END */

/* @MT-TPL-UPDATE-FUNC-START */
func (c *GameUpdateController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	var request_params struct {
		Name string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`

		Description string `json:"description" form:"description" validate:"Required;MinSize(3);MaxSize(20)"`

		HeaderHtml string `json:"header_html" form:"header_html" validate:"Required;MinSize(3);MaxSize(1024)"`

		StartTime int64 `json:"start_time" form:"start_time" validate:""`

		EndTime int64 `json:"end_time" form:"end_time" validate:""`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	id, err := primitive.ObjectIDFromHex(c.GetString("id", ""))
	if err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	// TODO: service
	c.Ctx.Output.JSON(
		master_service.UpdateGameService(
			user,
			id,
			request_params.Name,
			request_params.Description,
			request_params.HeaderHtml,
			request_params.StartTime,
			request_params.EndTime,
		),
		true, false)
}

/* @MT-TPL-UPDATE-FUNC-END */

/* @MT-TPL-DELETE-FUNC-START */
func (c *GameDeleteController) Post() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	id, err := primitive.ObjectIDFromHex(c.GetString("id", ""))
	if err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	// TODO
	c.Ctx.Output.JSON(
		master_service.DeleteGameService(
			user,
			id,
		),
		true, false)
}

/* @MT-TPL-DELETE-FUNC-END */

/* @MT-TPL-GET-FUNC-START */
func (c *GameGetController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	id, err := primitive.ObjectIDFromHex(c.GetString("id", ""))
	if err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	// TODO
	c.Ctx.Output.JSON(
		master_service.GetGameService(
			user,
			id,
		),
		true, false)
}

/* @MT-TPL-GET-FUNC-END */

/* @MT-TPL-LIST-FUNC-START */
func (c *GameListController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	var request_params struct {
		Index int64 `json:"index" form:"index"`
		Limit int64 `json:"limit" form:"limit"`
	}

	// TODO
	c.Ctx.Output.JSON(
		master_service.ListGameService(
			user,
			request_params.Index,
			request_params.Limit,
		),
		true, false)
}

/* @MT-TPL-LIST-FUNC-END */

/* @MT-TPL-SEARCH-FUNC-START */
func (c *GameSearchController) Get() {
	user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

	user := user_interface.(*master_types.User)

	var request_params struct {
		Name  string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`
		Index int64  `json:"index" form:"index"`
		Limit int64  `json:"limit" form:"limit"`
	}

	if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
		return
	}

	// TODO
	c.Ctx.Output.JSON(
		master_service.SearchGameService(
			user,
			request_params.Name,
			request_params.Index,
			request_params.Limit,
		),
		true, false)
}

/* @MT-TPL-SEARCH-FUNC-END */
