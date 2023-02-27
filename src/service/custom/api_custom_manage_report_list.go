package custom

/* @MT-TPL-IMPORT-START */
import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/manage/report/list Service 获取红队提交的报告列表
func ApiCustomManageReportListService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    Page int,
    PageSize int,
    Order int,
) (*master_types.MasterResponse) {
    var apiCustomManageReportListResponse struct {
        Success bool `json:"success"`
        Reports any `json:"reports"`
        Total int `json:"total"`
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

    
    // list Report
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
        value, err := model.ModelGetAll[master_types.Report](
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

        apiCustomManageReportListResponse.Reports = value
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageReportListResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
