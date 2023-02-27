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
type ApiCustomAttackerReportController struct {
    beego.Controller
}

func (c *ApiCustomAttackerReportController) Post() {
    user_interface := c.Ctx.Input.GetData("user")
	if user_interface == nil {
		c.Ctx.Output.JSON(master_types.ErrorResponse(-401, "require login"), true, false)
		return
	}

    user := user_interface.(*master_types.User)

    var request_params struct {
        Content string `json:"content" form:"content" valid:"Required;MinSize(1);MaxSize(16384)"`
        IsolationBreak int `json:"isolation_break" form:"isolation_break" valid:"Min(0);Max(1);Required"`
        VulnType int `json:"vuln_type" form:"vuln_type" valid:"Min(0);Max(9);Required"`
        AchievementType int `json:"achievement_type" form:"achievement_type" valid:"Min(0);Max(1);Required"`
        AttackType int `json:"attack_type" form:"attack_type" valid:"Min(0);Max(5);Required"`
        Uri string `json:"uri" form:"uri" valid:"MinSize(0);MaxSize(256);Required"`
        VulnLevel int `json:"vuln_level" form:"vuln_level" valid:"Min(0);Max(3);Required"`
        Name string `json:"name" form:"name" valid:"MinSize(0);MaxSize(256);Required"`
    }


    request_params_game_id, err := primitive.ObjectIDFromHex(c.GetString("game_id"))
    if err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }
    request_params_defender_id, err := primitive.ObjectIDFromHex(c.GetString("defender_id"))
    if err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    if err := controller.ParseAndValidate(&request_params, c.Controller); err != nil {
        c.Ctx.Output.JSON(master_types.ErrorResponse(-400, err.Error()), true, false)
        return
    }

    response := custom_service.ApiCustomAttackerReportService(
        user,
        request_params_game_id,
        request_params_defender_id,
        request_params.Content,
        request_params.IsolationBreak,
        request_params.VulnType,
        request_params.AchievementType,
        request_params.AttackType,
        request_params.Uri,
        request_params.VulnLevel,
        request_params.Name,
    )
/* @MT-TPL-CONTROLLER-END */

	/* @MT-TPL-CONTROLLER-RESPONSE-START */

    c.Ctx.Output.JSON(response, true, false)
}

    /* @MT-TPL-CONTROLLER-RESPONSE-END */
