/* @MT-TPL-PACKAGE-START */
package custom
/* @MT-TPL-PACKAGE-END */

/* @MT-TPL-IMPORT-START */
import (
    beego "github.com/beego/beego/v2/server/web"
    controller "github.com/0xunion/exercise_back/src/controller"
    master_types "github.com/0xunion/exercise_back/src/types"
    custom_service "github.com/0xunion/exercise_back/src/service/custom"
    
)
/* @MT-TPL-IMPORT-END */

/* @MT-TPL-CONTROLLER-START */
type ApiCustomAdminGameCreateController struct {
    beego.Controller
}

func (c *ApiCustomAdminGameCreateController) Post() {
    user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

    user := user_interface.(*master_types.User)

    var request_params struct {
        Name string `json:"name" form:"name" validate:"MinSize(3),MaxSize(64)"`
        Description string `json:"description" form:"description" validate:"MinSize(3),MaxSize(64)"`
        HeaderHtml string `json:"header_html" form:"header_html" validate:"MinSize(3),MaxSize(1024)"`
        StartTime int64 `json:"start_time" form:"start_time" validate:""`
        EndTime int64 `json:"end_time" form:"end_time" validate:""`
    }



    if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    response := custom_service.ApiCustomAdminGameCreateService(
        user,
        request_params.Name,
        request_params.Description,
        request_params.HeaderHtml,
        request_params.StartTime,
        request_params.EndTime,
    )
    c.Ctx.Output.JSON(response, true, false)
}
/* @MT-TPL-CONTROLLER-END */
