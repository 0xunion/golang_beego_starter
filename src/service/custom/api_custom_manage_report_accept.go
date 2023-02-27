package custom

/* @MT-TPL-IMPORT-START */
import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/manage/report/accept Service 为红队提交的报告打分
func ApiCustomManageReportAcceptService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    ReportId master_types.PrimaryId,
    Score int,
    DefenderScore int,
) (*master_types.MasterResponse) {
    var apiCustomManageReportAcceptResponse struct {
        Success bool `json:"success"`
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

    
    // get Report

    var D_report *master_types.Report

    {
value, err := model.ModelGet[master_types.Report](
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

    // update Report
    
    {
        err := model.ModelUpdateField[master_types.Report](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", ReportId),
                model.MongoKeyFilter("state", master_types.REPORT_STATE_UNVERIFIED),
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

    // update Defender
    
    {
        err := model.ModelUpdateField[master_types.Defender](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", D_report.DefenderId),
            ),



            model.MongoDecField(
                "score", 
                DefenderScore,
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



            model.MongoIncField(
                "score", 
                Score,
            ),
        )

        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }
    }

        
    // set response directly
    apiCustomManageReportAcceptResponse.Success = true
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomManageReportAcceptResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
