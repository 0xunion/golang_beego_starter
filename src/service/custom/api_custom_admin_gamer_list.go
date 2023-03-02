package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/gamer/list Service 获取参赛者列表
func ApiCustomAdminGamerListService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    Page int64,
    PageSize int64,
) (*master_types.MasterResponse) {
    var apiCustomAdminGamerListResponse struct {
        Success bool `json:"success"`
        Gamers []master_types.Gamer `json:"gamers"`
    }

    access_controll := false
    if !access_controll && user.IsAdmin() {
        access_controll = true
    }

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }

    
    // list Gamer
    var D_page int64 = 1
    var D_limit int64 = 10
    var D_sort = ""
    var D_value = 1 // 1: asc, -1: desc
    D_page = int64(Page)
    D_limit = int64(PageSize)
    D_sort = "_id"
    D_value = -1


    {
        var skip = int64((D_page - 1) * D_limit)
        var limit = int64(D_limit)
        value, err := model.ModelGetAll[master_types.Gamer](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                // skip
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

        apiCustomAdminGamerListResponse.Gamers = value
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAdminGamerListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
