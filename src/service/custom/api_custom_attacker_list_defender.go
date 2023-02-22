package custom

/* @MT-TPL-IMPORT-START */
import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/list_defender Service 获取防守方列表
func ApiCustomAttackerListDefenderService(
    user *master_types.User,
    GameId master_types.PrimaryId,
) (*master_types.MasterResponse) {
    var apiCustomAttackerListDefenderResponse struct {
        Success bool `json:"success"`
        Defenders any `json:"defenders"`
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
                model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_ATTACKER),
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

    return master_types.SuccessResponse(apiCustomAttackerListDefenderResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
