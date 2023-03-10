package custom

import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	/* @MT-TPL-IMPORT-TIME-START */
    /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/defender/report/submit Service 蓝队提交溯源报告
func ApiCustomDefenderReportSubmitService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    Content string,
    AttackTeamId master_types.PrimaryId,
    Title string,
) (*master_types.MasterResponse) {
    var apiCustomDefenderReportSubmitResponse struct {
        Success bool `json:"success"`
    }

    access_controll := false
        access_controll = true

    if !access_controll {
        return master_types.ErrorResponse(-403, "Permission denied")
    }

    
    // get RedTeam


    {
_, err := model.ModelGet[master_types.RedTeam](
            model.NewMongoFilter(
                model.MongoKeyFilter("game_id", GameId),
                model.MongoKeyFilter("_id", AttackTeamId),
            ),
        )
        if err != nil {
            return master_types.ErrorResponse(-500, err.Error())
        }


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

    // create TraceReport
    var trace_report *master_types.TraceReport = &master_types.TraceReport{
        GameId: GameId,
        DefenderId: D_gamer.GroupId,
        AttackTeamId: AttackTeamId,
        Content: Content,
        Owner: user.Id,
        Title: Title,
        CreateAt: time.Now().Unix(),
    }

    err := model.ModelInsert(trace_report, nil)
    if err != nil {
        return master_types.ErrorResponse(-500, err.Error())
    }
/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomDefenderReportSubmitResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
