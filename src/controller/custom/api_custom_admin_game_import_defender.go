/* @MT-TPL-PACKAGE-START */
package custom
/* @MT-TPL-PACKAGE-END */

/* @MT-TPL-IMPORT-START */
import (
    beego "github.com/beego/beego/v2/server/web"
    controller "github.com/0xunion/exercise_back/src/controller"
    master_types "github.com/0xunion/exercise_back/src/types"
    custom_service "github.com/0xunion/exercise_back/src/service/custom"
    


	"go.mongodb.org/mongo-driver/bson/primitive"
)
/* @MT-TPL-IMPORT-END */

/* @MT-TPL-CONTROLLER-START */
type ApiCustomAdminGameImportDefenderController struct {
    beego.Controller
}

func (c *ApiCustomAdminGameImportDefenderController) Post() {
    user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

    user := user_interface.(*master_types.User)

    var request_params struct {
    }


    request_params_game_id, err := primitive.ObjectIDFromHex(c.GetString("game_id"))
    if err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }
    request_params_defender_file_id, err := primitive.ObjectIDFromHex(c.GetString("defender_file_id"))
    if err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    response := custom_service.ApiCustomAdminGameImportDefenderService(
        user,
        request_params_game_id,
        request_params_defender_file_id,
    )
    c.Ctx.Output.JSON(response, true, false)
}
/* @MT-TPL-CONTROLLER-END */