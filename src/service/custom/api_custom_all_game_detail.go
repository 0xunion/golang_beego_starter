package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/all/game/detail Service 获取比赛详情
func ApiCustomAllGameDetailService(
    user *master_types.User,
    GameId master_types.PrimaryId,
) (*master_types.MasterResponse) {
    var apiCustomAllGameDetailResponse struct {
        Success bool `json:"success"`
        Game *master_types.Game `json:"game"`
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
            ),
        )
        if err == nil && model_instance != nil {
            access_controll = true
        }
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }

    
    // get Game


    {
value, err := model.ModelGet[master_types.Game](
            model.NewMongoFilter(
                model.MongoKeyFilter("_id", GameId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }


        apiCustomAllGameDetailResponse.Game = value
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAllGameDetailResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
