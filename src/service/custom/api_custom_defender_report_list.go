package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

// /api/custom/defender/report/list Service 蓝队获取自己的溯源报告列表
func ApiCustomDefenderReportListService(
	user *master_types.User,
	GameId master_types.PrimaryId,
	Page int64,
	PageSize int64,
) *master_types.MasterResponse {
	var apiCustomDefenderReportListResponse struct {
		Success bool `json:"success"`
		Reports any  `json:"reports"`
	}

	access_controll := false
	access_controll = true

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

	// list TraceReport
	var D_page int64 = 1
	var D_limit int64 = 10
	var D_sort = ""
	var D_value = 1 // 1: asc, -1: desc
	D_page = int64(Page)
	D_sort = "_id"
	D_value = -1

	{
		var skip = int64((D_page - 1) * D_limit)
		var limit = int64(D_limit)
		value, err := model.ModelGetAll[master_types.TraceReport](
			model.NewMongoFilter(
				model.MongoKeyFilter("game_id", GameId),
				model.MongoKeyFilter("defender_id", D_gamer.GroupId),
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

		apiCustomDefenderReportListResponse.Reports = value
	}

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomDefenderReportListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
