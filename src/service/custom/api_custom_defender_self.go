package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */ /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/defender/self Service 获取防守方自身信息
func ApiCustomDefenderSelfService(
	user *master_types.User,
	GameId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomDefenderSelfResponse struct {
		Defender *master_types.Defender `json:"defender"`
	}

	access_controll := false
	if !access_controll {
		model_instance, err := model.ModelGet[master_types.Gamer](
			model.NewMongoFilter(
				model.MongoKeyFilter("game_id", GameId),
				model.MongoKeyFilter("owner", user.Id),
				model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_DEFENDER),
			),
		)
		if err == nil && model_instance != nil {
			access_controll = true
		}
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Permission denied")
	}

	// get Gamer

	var D_gamer *master_types.Gamer

	{
		value, err := model.ModelGet[master_types.Gamer](
			model.NewMongoFilter(
				model.MongoKeyFilter("game_id", GameId),
				model.MongoKeyFilter("owner", user.Id),
				model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_DEFENDER),
			),
		)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}

		D_gamer = value

	}

	// get Defender

	{
		value, err := model.ModelGet[master_types.Defender](
			model.NewMongoFilter(
				model.MongoKeyFilter("game_id", GameId),
				model.MongoKeyFilter("_id", D_gamer.GroupId),
			),
		)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}

		apiCustomDefenderSelfResponse.Defender = value
	}
	/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomDefenderSelfResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
