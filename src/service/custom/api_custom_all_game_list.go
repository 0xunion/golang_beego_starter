package custom

import (
	"github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/all/game/list Service 获取所有比赛列表
func ApiCustomAllGameListService(
    user *master_types.User,
) (*master_types.MasterResponse) {
    var apiCustomAllGameListResponse struct {
        Success bool `json:"success"`
        Games any `json:"games"`
    }

    access_controll := false
        access_controll = true

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do
	type game struct {
		Id       master_types.PrimaryId `json:"id" bson:"_id"`
		Name     string                 `json:"name" bson:"name"`
		Identity int64                  `json:"identity" bson:"identity"`
		CreateAt int64                  `json:"create_at" bson:"create_at"`
		Owner    master_types.PrimaryId `json:"owner" bson:"owner"`
		GameId   master_types.PrimaryId `json:"game_id" bson:"game_id"`
		Score    int64                  `json:"score" bson:"score"`
		Game     master_types.Game      `json:"game" bson:"game" join:"game_id=_id"`
	}

	games, err := model.ModelGetAllJoin[master_types.Gamer, game](
		model.NewMongoFilter(
			model.MongoKeyFilter("owner", user.Id),
		),
	)

	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	apiCustomAllGameListResponse.Success = true
	apiCustomAllGameListResponse.Games = games

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAllGameListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
