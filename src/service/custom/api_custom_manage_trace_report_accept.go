package custom

import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/manage/trace_report/accept Service 裁判接受溯源报告
func ApiCustomManageTraceReportAcceptService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    ReportId master_types.PrimaryId,
    Score int,
    AttackScore int,
) (*master_types.MasterResponse) {
    var apiCustomManageTraceReportAcceptResponse struct {
        Success bool `json:"success"`
    }

    access_controll := false
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

    
    // get TraceReport

    var D_report *master_types.TraceReport

    {
value, err := model.ModelGet[master_types.TraceReport](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", ReportId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }

        D_report = value

    }

    // update TraceReport
    
    {
        err := model.ModelUpdateField[master_types.TraceReport](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", ReportId),
            ),



            model.MongoIncField(
                "score", 
                Score,
            ),
            model.MongoSetField(
                "state", 
                master_types.REPORT_STATE_ACCEPTED,
            ),
        )

        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }
    }

    // update RedTeam
    
    {
        err := model.ModelUpdateField[master_types.RedTeam](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", D_report.AttackTeamId),
            ),



            model.MongoDecField(
                "score", 
                AttackScore,
            ),
        )

        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }
    }

    // update Defender
    
    {
        err := model.ModelUpdateField[master_types.Defender](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", D_report.DefenderId),
            ),



            model.MongoIncField(
                "score", 
                Score,
            ),
        )

        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageTraceReportAcceptResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
