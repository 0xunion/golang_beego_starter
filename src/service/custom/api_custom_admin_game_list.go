package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/list Service 列出所有比赛
func ApiCustomAdminGameListService(
    user *master_types.User,
    Page int64,
    PageSize int64,
) (*master_types.MasterResponse) {
    var apiCustomAdminGameListResponse struct {
        Success bool `json:"success"`
        Games []master_types.Game `json:"games"`
    }

    access_controll := false
    if !access_controll && user.IsAdmin() {
        access_controll = true
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do
	// list all games
	skip := (Page - 1) * PageSize
	games, err := model.ModelGetAll[master_types.Game](
		model.NewMongoFilter(
			model.MongoNoFlagFilter(master_types.BASIC_TYPE_FLAG_DELETED),
		),
		&model.MongoOptions{
			Skip:  &skip,
			Limit: &PageSize,
		},
	)

	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	apiCustomAdminGameListResponse.Success = true
	apiCustomAdminGameListResponse.Games = games

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAdminGameListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
