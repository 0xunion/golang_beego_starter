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
type GamerCreateController struct {
    beego.Controller
}
/* @MT-TPL-CREATE-END */

/* @MT-TPL-UPDATE-START */
type GamerUpdateController struct {
    beego.Controller
}
/* @MT-TPL-UPDATE-END */

/* @MT-TPL-DELETE-START */
type GamerDeleteController struct {
    beego.Controller
}
/* @MT-TPL-DELETE-END */

/* @MT-TPL-GET-START */
type GamerGetController struct {
    beego.Controller
}
/* @MT-TPL-GET-END */

/* @MT-TPL-LIST-START */
type GamerListController struct {
    beego.Controller
}
/* @MT-TPL-LIST-END */

/* @MT-TPL-SEARCH-START */
type GamerSearchController struct {
    beego.Controller
}
/* @MT-TPL-SEARCH-END */

/* @MT-TPL-CREATE-FUNC-START */
func (c *GamerCreateController) Post() {
    user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

    user := user_interface.(*master_types.User)

    var request_params struct {
        
        Name string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`
        
        Phone string `json:"phone" form:"phone" validate:"Required;MinSize(3);MaxSize(20)"`
        
        Identity int64 `json:"identity" form:"identity" validate:""`
        
        GameId master_types.PrimaryId `json:"game_id" form:"game_id" validate:""`
        
        Permission int64 `json:"permission" form:"permission" validate:""`
    }

    if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    

    // TODO: service
    c.Ctx.Output.JSON(
        master_service.CreateGamerService(
            user,
            request_params.Name,
            request_params.Phone,
            request_params.Identity,
            request_params.GameId,
            request_params.Permission,
        ),
    true, false)
}
/* @MT-TPL-CREATE-FUNC-END */

/* @MT-TPL-UPDATE-FUNC-START */
func (c *GamerUpdateController) Post() {
    user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

    user := user_interface.(*master_types.User)

    var request_params struct {
    
        Name string `json:"name" form:"name" validate:"Required;MinSize(3);MaxSize(20)"`
    
        Phone string `json:"phone" form:"phone" validate:"Required;MinSize(3);MaxSize(20)"`
    
        Identity int64 `json:"identity" form:"identity" validate:""`
    
        GameId master_types.PrimaryId `json:"game_id" form:"game_id" validate:""`
    
        Permission int64 `json:"permission" form:"permission" validate:""`
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
        master_service.UpdateGamerService(
            user,
            id,
            request_params.Name,
            request_params.Phone,
            request_params.Identity,
            request_params.GameId,
            request_params.Permission,
        ),
    true, false)
}
/* @MT-TPL-UPDATE-FUNC-END */

/* @MT-TPL-DELETE-FUNC-START */
func (c *GamerDeleteController) Post() {
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
        master_service.DeleteGamerService(
            user,
            id,
        ),
    true, false)
}
/* @MT-TPL-DELETE-FUNC-END */

/* @MT-TPL-GET-FUNC-START */
func (c *GamerGetController) Get() {
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
        master_service.GetGamerService(
            user,
            id,
        ),
    true, false)
}
/* @MT-TPL-GET-FUNC-END */

/* @MT-TPL-LIST-FUNC-START */
func (c *GamerListController) Get() {
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
        master_service.ListGamerService(
            user,
            request_params.Index,
            request_params.Limit,
        ),
    true, false)
}
/* @MT-TPL-LIST-FUNC-END */

/* @MT-TPL-SEARCH-FUNC-START */
func (c *GamerSearchController) Get() {
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
        master_service.SearchGamerService(
            user,
            request_params.Name,
            request_params.Index,
            request_params.Limit,
        ),
    true, false)
}
/* @MT-TPL-SEARCH-FUNC-END */