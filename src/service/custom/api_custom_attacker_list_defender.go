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
        Defenders []master_types.Defender `json:"defenders"`
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

    
    // list Defender
    var D_page int64 = 1
    var D_limit int64 = 10
    var D_sort = ""
    var D_value = 1 // 1: asc, -1: desc


    {
        var skip = int64((D_page - 1) * D_limit)
        var limit = int64(D_limit)
        value, err := model.ModelGetAll[master_types.Defender](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
            ),
            &model.MongoOptions{
                Skip:  &skip,
                Limit: &limit,
                Sort:  model.MongoSort(D_sort, D_value),
            },
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }

        apiCustomAttackerListDefenderResponse.Defenders = value
    }

        
    // set response directly
    apiCustomAttackerListDefenderResponse.Success = true
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAttackerListDefenderResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
