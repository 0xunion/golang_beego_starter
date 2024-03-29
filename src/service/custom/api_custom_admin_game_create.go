package custom

/* @MT-TPL-IMPORT-START */
import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/create Service 创建一个比赛
func ApiCustomAdminGameCreateService(
    user *master_types.User,
    Name string,
    Description string,
    HeaderHtml string,
    StartTime int64,
    EndTime int64,
    PositionCode string,
) (*master_types.MasterResponse) {
    var apiCustomAdminGameCreateResponse struct {
        Success bool `json:"success"`
    }

    access_controll := false
    if !access_controll && user.IsAdmin() {
        access_controll = true
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }

    
    // create Game
    var game *master_types.Game = &master_types.Game{
        Name: Name,
        Description: Description,
        HeaderHtml: HeaderHtml,
        StartTime: StartTime,
        EndTime: EndTime,
        CreateAt: time.Now().Unix(),
        Owner: user.Id,
        PositionCode: PositionCode,
    }

    err := model.ModelInsert(game, nil)
    if err != nil {
        return master_types.ErrorResponse(-500, err.Error())
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do
	apiCustomAdminGameCreateResponse.Success = true

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAdminGameCreateResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
