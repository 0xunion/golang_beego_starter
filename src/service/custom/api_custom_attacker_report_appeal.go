package custom

/* @MT-TPL-IMPORT-START */
import (
	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/attacker/report/appeal Service 申诉攻击报告
func ApiCustomAttackerReportAppealService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    ReportId master_types.PrimaryId,
    Reason string,
) (*master_types.MasterResponse) {
    var apiCustomAttackerReportAppealResponse struct {
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
                model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_ATTACKER),
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
                model.MongoKeyFilter("identity", master_types.GAMER_IDENTITY_ATTACKER),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }

        D_gamer = value

    }

    // update Report
    
    {
        err := model.ModelUpdateField[master_types.Report](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("attack_team_id", D_gamer.GroupId),
                model.MongoKeyFilter("_id", ReportId),
            ),



            model.MongoSetField(
                "state", 
                master_types.REPORT_STATE_UNVERIFIED,
            ),
        )

        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }
    }

    // create ReportComment
    var report_comment *master_types.ReportComment = &master_types.ReportComment{
        GameId: GameId,
        ReportId: ReportId,
        Owner: user.Id,
        Content: Reason,
        CreateAt: time.Now().Unix(),
    }

    err := model.ModelInsert(report_comment, nil)
    if err != nil {
        return master_types.ErrorResponse(-500, err.Error())
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAttackerReportAppealResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
