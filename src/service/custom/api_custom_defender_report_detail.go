package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/defender/report/detail Service 蓝队获取自己的溯源报告详情
func ApiCustomDefenderReportDetailService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    ReportId master_types.PrimaryId,
) (*master_types.MasterResponse) {
    var apiCustomDefenderReportDetailResponse struct {
        Success bool `json:"success"`
        Report any `json:"report"`
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

    // get TraceReport


    {
value, err := model.ModelGet[master_types.TraceReport](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("defender_id", D_gamer.GroupId),
                model.MongoKeyFilter("_id", ReportId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }


        apiCustomDefenderReportDetailResponse.Report = value
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomDefenderReportDetailResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
