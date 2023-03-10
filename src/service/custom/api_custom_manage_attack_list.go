package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

// /api/custom/manage/attack/list Service 列出所有攻击
func ApiCustomManageAttackListService(
	user *master_types.User,
	GameId master_types.PrimaryId,
	Page int,
	PageSize int,
	State int64,
	Content string,
) *master_types.MasterResponse {
	var apiCustomManageAttackListResponse struct {
		Attacks any `json:"attacks"`
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

	// list Attack
	var D_page int64 = 1
	var D_limit int64 = 10
	var D_sort = ""
	var D_value = 1 // 1: asc, -1: desc
	D_page = int64(Page)
	D_sort = "_id"
	D_value = -1

	{
		filters := make([]model.MongoFilterItem, 0)
		filters = append(filters, model.MongoKeyFilter("game_id", GameId))
		if State != -1 {
			filters = append(filters, model.MongoKeyFilter("state", State))
		}
		if Content != "" {
			filters = append(filters, model.MongoSearchFilter("reason", Content))
		}

		var skip = int64((D_page - 1) * D_limit)
		var limit = int64(D_limit)
		value, err := model.ModelGetAll[master_types.Attack](
			model.NewMongoFilter(
				filters...,
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

		apiCustomManageAttackListResponse.Attacks = value
	}

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageAttackListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
