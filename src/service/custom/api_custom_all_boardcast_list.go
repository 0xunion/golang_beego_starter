package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

// /api/custom/all/boardcast/list Service 获取所有公告列表
func ApiCustomAllBoardcastListService(
	user *master_types.User,
	GameId master_types.PrimaryId,
	Page int64,
	PageSize int64,
) *master_types.MasterResponse {
	var apiCustomAllBoardcastListResponse struct {
		Success    bool `json:"success"`
		Boardcasts any  `json:"boardcasts"`
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

	// list Boardcast
	var D_page int64 = 1
	var D_limit int64 = 10
	var D_sort = ""
	var D_value = 1 // 1: asc, -1: desc
	D_page = int64(Page)
	D_sort = "create_at"
	D_value = -1

	{
		var skip = int64((D_page - 1) * D_limit)
		var limit = int64(D_limit)
		value, err := model.ModelGetAll[master_types.Boardcast](
			model.NewMongoFilter(
				model.MongoKeyFilter("game_id", GameId),
				// skip
				// skip
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

		apiCustomAllBoardcastListResponse.Boardcasts = value
	}

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAllBoardcastListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
