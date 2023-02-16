package curd

/* @MT-TPL-IMPORT-START */
import (
	controller "github.com/0xunion/exercise_back/src/controller"
	master_types "github.com/0xunion/exercise_back/src/types"
    master_service "github.com/0xunion/exercise_back/src/service/curd"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
/* @MT-TPL-IMPORT-END */

/* @MT-TPL-CREATE-START */
type DefenderCreateController struct {
    beego.Controller
}
/* @MT-TPL-CREATE-END */

/* @MT-TPL-UPDATE-START */
type DefenderUpdateController struct {
    beego.Controller
}
/* @MT-TPL-UPDATE-END */

/* @MT-TPL-DELETE-START */
type DefenderDeleteController struct {
    beego.Controller
}
/* @MT-TPL-DELETE-END */

/* @MT-TPL-GET-START */
type DefenderGetController struct {
    beego.Controller
}
/* @MT-TPL-GET-END */

/* @MT-TPL-LIST-START */
type DefenderListController struct {
    beego.Controller
}
/* @MT-TPL-LIST-END */

/* @MT-TPL-SEARCH-START */
type DefenderSearchController struct {
    beego.Controller
}
/* @MT-TPL-SEARCH-END */

/* @MT-TPL-CREATE-FUNC-START */
func (c *DefenderCreateController) Post() {
    user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

    user := user_interface.(*master_types.User)

    var request_params struct {
        
        Name string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`
        
        GameId master_types.PrimaryId `json:"game_id" form:"game_id" validate:""`
    }

    if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    

    // TODO: service
    c.Ctx.Output.JSON(
        master_service.CreateDefenderService(
            user,
            request_params.Name,
            request_params.GameId,
        ),
    true, false)
}
/* @MT-TPL-CREATE-FUNC-END */

/* @MT-TPL-UPDATE-FUNC-START */
func (c *DefenderUpdateController) Post() {
    user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

    user := user_interface.(*master_types.User)

    var request_params struct {
    
        Name string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`
    
        GameId master_types.PrimaryId `json:"game_id" form:"game_id" validate:""`
    
        Score int64 `json:"score" form:"score" validate:""`
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
        master_service.UpdateDefenderService(
            user,
            id,
            request_params.Name,
            request_params.GameId,
            request_params.Score,
        ),
    true, false)
}
/* @MT-TPL-UPDATE-FUNC-END */

/* @MT-TPL-DELETE-FUNC-START */
func (c *DefenderDeleteController) Post() {
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
        master_service.DeleteDefenderService(
            user,
            id,
        ),
    true, false)
}
/* @MT-TPL-DELETE-FUNC-END */

/* @MT-TPL-GET-FUNC-START */
func (c *DefenderGetController) Get() {
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
        master_service.GetDefenderService(
            user,
            id,
        ),
    true, false)
}
/* @MT-TPL-GET-FUNC-END */

/* @MT-TPL-LIST-FUNC-START */
func (c *DefenderListController) Get() {
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
        master_service.ListDefenderService(
            user,
            request_params.Index,
            request_params.Limit,
        ),
    true, false)
}
/* @MT-TPL-LIST-FUNC-END */

/* @MT-TPL-SEARCH-FUNC-START */
func (c *DefenderSearchController) Get() {
    user_interface := c.Ctx.Input.GetData("user")
    if user_interface == nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
        return
    }

    user := user_interface.(*master_types.User)

    var request_params struct {
        Name string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`
        Index int64 `json:"index" form:"index"`
        Limit int64 `json:"limit" form:"limit"`
    }

    if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    // TODO
    c.Ctx.Output.JSON(
        master_service.SearchDefenderService(
            user,
            request_params.Name,
            request_params.Index,
            request_params.Limit,
        ),
    true, false)
}
/* @MT-TPL-SEARCH-FUNC-END */