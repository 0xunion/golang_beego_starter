package custom

/* @MT-TPL-IMPORT-START */
import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/manage/report/list Service 获取红队提交的报告列表
func ApiCustomManageReportListService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    Page int,
    PageSize int,
    Order int,
) (*master_types.MasterResponse) {
    var apiCustomManageReportListResponse struct {
        Success bool `json:"success"`
        Reports any `json:"reports"`
        Total int `json:"total"`
    }

    access_controll := false
    if !access_controll && user.IsAdmin() {
        access_controll = true
    }
    if !access_controll {
        model_instance, err := model.ModelGet[master_types.Gamer](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("owner", user.Id),
                model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_JUDGEMENT),
            ),
        )
        if err == nil && model_instance != nil {
            access_controll = true
        }
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageReportListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
